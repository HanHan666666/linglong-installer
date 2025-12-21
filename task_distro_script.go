package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

const (
	defaultScriptsDir = "scripts/distros"
	metaPrefix        = "meta:"
)

type distroScriptTask struct {
	core.BaseTask
	ScriptsDir string
}

type scriptMeta struct {
	RepoName   string
	RepoURL    string
	NextSteps  string
	Commands   []string
	DistroName string
}

func RegisterDistroScriptTask() {
	core.Tasks.Register("distroScript", func(config map[string]any, ctx *core.InstallContext) (core.Task, error) {
		scriptsDir := defaultScriptsDir
		if val, ok := config["scriptsDir"].(string); ok && val != "" {
			scriptsDir = val
		} else if val, ok := config["scripts_dir"].(string); ok && val != "" {
			scriptsDir = val
		}

		task := &distroScriptTask{
			BaseTask: core.BaseTask{
				TaskID:   getConfigString(config, "id", "resolve-distro-script"),
				TaskType: "distroScript",
				Config:   config,
			},
			ScriptsDir: scriptsDir,
		}

		return task, nil
	})
}

func (t *distroScriptTask) Validate() error {
	if t.ScriptsDir == "" {
		return errors.New("distroScript: scriptsDir is required")
	}
	return nil
}

func (t *distroScriptTask) Execute(ctx *core.InstallContext, bus *core.EventBus) error {
	id := strings.TrimSpace(ctx.Env.Distro)
	version := strings.TrimSpace(ctx.Env.DistroVersion)

	ctx.Set("distro.id", id)
	ctx.Set("distro.version", version)

	prettyName := readPrettyName()
	if prettyName == "" {
		prettyName = strings.TrimSpace(fmt.Sprintf("%s %s", id, version))
	}
	if prettyName != "" {
		ctx.Set("distro.name", prettyName)
	}

	if id == "" || version == "" {
		ctx.Set("distro.supported", false)
		ctx.Set("distro.error", "Missing distro ID or VERSION_ID")
		ctx.Set("distro.next_steps", "No install script can be selected without distro ID/version.")
		ctx.Set("distro.commands", "(none)")
		ctx.Set("distro.repo_name", "(unknown)")
		ctx.Set("distro.repo_url", "")
		return nil
	}

	candidates := []string{fmt.Sprintf("%s_%s.sh", id, version)}
	if lower := fmt.Sprintf("%s_%s.sh", strings.ToLower(id), version); lower != candidates[0] {
		candidates = append(candidates, lower)
	}

	var meta scriptMeta
	var scriptPath string
	var err error
	for _, name := range candidates {
		scriptPath, meta, err = resolveScript(name, t.ScriptsDir)
		if err == nil {
			ctx.AddLog(core.LogInfo, fmt.Sprintf("Resolved distro script: %s", scriptPath))
			break
		}
	}
	if err != nil {
		ctx.Set("distro.supported", false)
		ctx.Set("distro.error", fmt.Sprintf("Unsupported distro: %s %s", id, version))
		ctx.Set("distro.next_steps", "No matching install script was found for this distro.")
		ctx.Set("distro.commands", "(none)")
		ctx.Set("distro.repo_name", "(unknown)")
		ctx.Set("distro.repo_url", "")
		ctx.AddLog(core.LogWarn, fmt.Sprintf("No distro script found for %s %s", id, version))
		return nil
	}

	ctx.Set("distro.supported", true)
	ctx.Set("distro.script", scriptPath)
	ctx.Set("distro.script_name", filepath.Base(scriptPath))

	if meta.DistroName != "" {
		ctx.Set("distro.name", meta.DistroName)
	}

	if meta.RepoName == "" {
		meta.RepoName = "(unknown)"
	}
	ctx.Set("distro.repo_name", meta.RepoName)
	ctx.Set("distro.repo_url", meta.RepoURL)

	if meta.NextSteps == "" {
		meta.NextSteps = "Check Linglong runtime and install if needed."
	}
	ctx.Set("distro.next_steps", meta.NextSteps)

	commands := "(none)"
	if len(meta.Commands) > 0 {
		lines := make([]string, 0, len(meta.Commands))
		for _, cmd := range meta.Commands {
			lines = append(lines, "- "+cmd)
		}
		commands = strings.Join(lines, "\n")
	}
	ctx.Set("distro.commands", commands)

	return nil
}

func resolveScript(scriptName, scriptsDir string) (string, scriptMeta, error) {
	if filepath.IsAbs(scriptsDir) {
		path := filepath.Join(scriptsDir, scriptName)
		if fileExists(path) {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return "", scriptMeta{}, err
			}
			meta, err := readScriptMetaFile(absPath)
			return absPath, meta, err
		}
	} else {
		roots := []string{}
		if cwd, err := os.Getwd(); err == nil {
			roots = append(roots, filepath.Join(cwd, scriptsDir))
		}
		if exe, err := os.Executable(); err == nil {
			roots = append(roots, filepath.Join(filepath.Dir(exe), scriptsDir))
		}
		for _, root := range roots {
			path := filepath.Join(root, scriptName)
			if fileExists(path) {
				absPath, err := filepath.Abs(path)
				if err != nil {
					return "", scriptMeta{}, err
				}
				meta, err := readScriptMetaFile(absPath)
				return absPath, meta, err
			}
		}
	}

	path, meta, err := extractEmbeddedScript(scriptName)
	if err == nil {
		return path, meta, nil
	}

	return "", scriptMeta{}, fmt.Errorf("script not found: %s", scriptName)
}

func extractEmbeddedScript(scriptName string) (string, scriptMeta, error) {
	scriptKey := filepath.ToSlash(filepath.Join("scripts", "distros", scriptName))
	scriptData, err := embeddedScripts.ReadFile(scriptKey)
	if err != nil {
		return "", scriptMeta{}, err
	}
	meta := parseScriptMeta(bytes.NewReader(scriptData))

	commonData, err := embeddedScripts.ReadFile("scripts/common.sh")
	if err != nil {
		return "", scriptMeta{}, err
	}

	tmpDir, err := os.MkdirTemp("", "linglong-installer-")
	if err != nil {
		return "", scriptMeta{}, err
	}

	distrosDir := filepath.Join(tmpDir, "distros")
	if err := os.MkdirAll(distrosDir, 0o755); err != nil {
		return "", scriptMeta{}, err
	}

	commonPath := filepath.Join(tmpDir, "common.sh")
	if err := os.WriteFile(commonPath, commonData, 0o644); err != nil {
		return "", scriptMeta{}, err
	}

	scriptPath := filepath.Join(distrosDir, scriptName)
	if err := os.WriteFile(scriptPath, scriptData, 0o755); err != nil {
		return "", scriptMeta{}, err
	}

	return scriptPath, meta, nil
}

func readScriptMetaFile(path string) (scriptMeta, error) {
	file, err := os.Open(path)
	if err != nil {
		return scriptMeta{}, err
	}
	defer file.Close()

	return parseScriptMeta(file), nil
}

func parseScriptMeta(r io.Reader) scriptMeta {
	meta := scriptMeta{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#!") {
			continue
		}
		if !strings.HasPrefix(line, "#") {
			break
		}

		content := strings.TrimSpace(strings.TrimPrefix(line, "#"))
		if content == "" {
			continue
		}
		if !strings.HasPrefix(strings.ToLower(content), metaPrefix) {
			continue
		}

		payload := strings.TrimSpace(content[len(metaPrefix):])
		key, value := splitMeta(payload)
		if key == "" {
			continue
		}

		switch key {
		case "repo_name":
			meta.RepoName = value
		case "repo_url":
			meta.RepoURL = value
		case "next_steps":
			meta.NextSteps = value
		case "command":
			if value != "" {
				meta.Commands = append(meta.Commands, value)
			}
		case "distro_name":
			meta.DistroName = value
		}
	}

	return meta
}

func splitMeta(payload string) (string, string) {
	if payload == "" {
		return "", ""
	}

	separator := "="
	if !strings.Contains(payload, separator) && strings.Contains(payload, ":") {
		separator = ":"
	}

	parts := strings.SplitN(payload, separator, 2)
	if len(parts) != 2 {
		return "", ""
	}

	key := strings.ToLower(strings.TrimSpace(parts[0]))
	key = strings.ReplaceAll(key, "-", "_")
	value := strings.TrimSpace(parts[1])

	return key, value
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func readPrettyName() string {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			value := strings.TrimPrefix(line, "PRETTY_NAME=")
			return strings.Trim(value, "\"")
		}
	}

	return ""
}

func getConfigString(config map[string]any, key string, fallback string) string {
	if val, ok := config[key].(string); ok && val != "" {
		return val
	}
	return fallback
}
