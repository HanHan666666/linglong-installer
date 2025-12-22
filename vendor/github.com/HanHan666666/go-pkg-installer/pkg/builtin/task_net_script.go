// Package builtin provides the net_script task implementation.
package builtin

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// NetScriptTask downloads and executes a script from the network.
type NetScriptTask struct {
	core.BaseTask
	URL              string
	SHA256           string
	Timeout          time.Duration
	Env              map[string]string
	WorkDir          string
	RequirePrivilege bool
}

// RegisterNetScriptTask registers the net_script task factory.
func RegisterNetScriptTask() {
	core.Tasks.Register("net_script", func(config map[string]any, ctx *core.InstallContext) (core.Task, error) {
		env := make(map[string]string)
		if envMap, ok := config["env"].(map[string]any); ok {
			for k, v := range envMap {
				if s, ok := v.(string); ok {
					env[k] = ctx.Render(s)
				}
			}
		}

		task := &NetScriptTask{
			BaseTask: core.BaseTask{
				TaskID:   getConfigString(config, "id"),
				TaskType: "net_script",
				Config:   config,
			},
			URL:              ctx.Render(getConfigString(config, "url")),
			SHA256:           getConfigString(config, "sha256"),
			Timeout:          time.Duration(getConfigIntAny(config, 300, "timeoutSec", "timeout")) * time.Second,
			Env:              env,
			WorkDir:          ctx.Render(getConfigStringAny(config, "workDir", "workdir")),
			RequirePrivilege: getConfigBool(config, "requirePrivilege"),
		}

		if task.TaskID == "" {
			task.TaskID = "net-script"
		}

		return task, nil
	})
}

// Validate validates the net_script task configuration.
func (t *NetScriptTask) Validate() error {
	if t.URL == "" {
		return errors.New("net_script: url is required")
	}
	return nil
}

// Execute downloads and runs the script.
func (t *NetScriptTask) Execute(ctx *core.InstallContext, bus *core.EventBus) error {
	if err := ensurePrivilege(ctx, t.RequirePrivilege); err != nil {
		return err
	}

	ctx.AddLog(core.LogInfo, fmt.Sprintf("Downloading script: %s", t.URL))

	client := &http.Client{Timeout: t.Timeout}
	resp, err := client.Get(t.URL)
	if err != nil {
		return fmt.Errorf("failed to download script: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("script download failed with status: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read script: %w", err)
	}

	if t.SHA256 != "" {
		hash := sha256.Sum256(data)
		actual := hex.EncodeToString(hash[:])
		if actual != t.SHA256 {
			return fmt.Errorf("checksum mismatch: expected %s, got %s", t.SHA256, actual)
		}
	}

	scriptPath := filepath.Join(os.TempDir(), fmt.Sprintf("gpki-script-%d.sh", time.Now().UnixNano()))
	if err := os.WriteFile(scriptPath, data, 0700); err != nil {
		return fmt.Errorf("failed to write script: %w", err)
	}
	defer os.Remove(scriptPath)

	ctx.AddLog(core.LogInfo, fmt.Sprintf("Executing script: %s", scriptPath))

	timeout := t.Timeout
	if timeout == 0 {
		timeout = 300 * time.Second
	}

	cmd := execCommand("sh", scriptPath)
	if t.WorkDir != "" {
		cmd.Dir = t.WorkDir
	}
	if len(t.Env) > 0 {
		cmd.Env = os.Environ()
		for k, v := range t.Env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
		}
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start script: %w", err)
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case err := <-done:
		if err != nil {
			return fmt.Errorf("script failed: %w: %s", err, stderr.String())
		}
	case <-ctxWithTimeout.Done():
		_ = cmd.Process.Kill()
		return fmt.Errorf("script timed out after %v", timeout)
	}

	if stdout.Len() > 0 {
		ctx.AddLog(core.LogInfo, fmt.Sprintf("stdout: %s", stdout.String()))
	}
	return nil
}

// CanRollback returns false by default.
func (t *NetScriptTask) CanRollback() bool {
	return false
}

// Rollback does nothing for net_script operations.
func (t *NetScriptTask) Rollback(ctx *core.InstallContext, bus *core.EventBus) error {
	return nil
}
