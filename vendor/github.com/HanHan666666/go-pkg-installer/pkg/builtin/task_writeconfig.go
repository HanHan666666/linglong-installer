// Package builtin provides the write_config task implementation.
package builtin

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
	"gopkg.in/yaml.v3"
)

// WriteConfigTask writes a configuration file.
type WriteConfigTask struct {
	core.BaseTask
	Destination      string
	Format           string // json, yaml, or text
	Content          any
	Mode             os.FileMode
	RequirePrivilege bool

	// For rollback
	wroteFile    string
	previousData []byte
	fileExisted  bool
}

// RegisterWriteConfigTask registers the write_config task factory.
func RegisterWriteConfigTask() {
	core.Tasks.Register("writeConfig", func(config map[string]any, ctx *core.InstallContext) (core.Task, error) {
		mode := os.FileMode(0644)
		if modeVal := parseFileMode(config["mode"]); modeVal != 0 {
			mode = os.FileMode(modeVal)
		}

		content := config["content"]
		if content == nil {
			if tmpl, ok := config["template"].(string); ok && tmpl != "" {
				content = tmpl
			}
		}

		task := &WriteConfigTask{
			BaseTask: core.BaseTask{
				TaskID:   getConfigString(config, "id"),
				TaskType: "writeConfig",
				Config:   config,
			},
			Destination:      ctx.Render(getConfigStringAny(config, "path", "destination")),
			Format:           getConfigString(config, "format"),
			Content:          content,
			Mode:             mode,
			RequirePrivilege: getConfigBool(config, "requirePrivilege"),
		}

		if task.TaskID == "" {
			task.TaskID = fmt.Sprintf("writeConfig-%s", filepath.Base(task.Destination))
		}

		// Auto-detect format from extension if not specified
		if task.Format == "" {
			ext := strings.ToLower(filepath.Ext(task.Destination))
			switch ext {
			case ".json":
				task.Format = "json"
			case ".yaml", ".yml":
				task.Format = "yaml"
			default:
				task.Format = "text"
			}
		}

		return task, nil
	})
}

// Validate validates the write_config task configuration.
func (t *WriteConfigTask) Validate() error {
	if t.Destination == "" {
		return errors.New("writeConfig: destination is required")
	}
	if t.Content == nil {
		return errors.New("writeConfig: content is required")
	}
	return nil
}

// Execute writes the configuration file.
func (t *WriteConfigTask) Execute(ctx *core.InstallContext, bus *core.EventBus) error {
	if err := ensurePrivilege(ctx, t.RequirePrivilege); err != nil {
		return err
	}

	ctx.AddLog(core.LogInfo, fmt.Sprintf("Writing config to %s (format: %s)", t.Destination, t.Format))

	// Ensure parent directory exists
	destDir := filepath.Dir(t.Destination)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Backup existing file for rollback
	if data, err := os.ReadFile(t.Destination); err == nil {
		t.previousData = data
		t.fileExisted = true
	}

	// Render content (if it's a string, apply template)
	content := t.renderContent(ctx)

	// Convert content to bytes based on format
	var data []byte
	var err error

	switch t.Format {
	case "json":
		data, err = json.MarshalIndent(content, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}
		data = append(data, '\n')

	case "yaml":
		data, err = yaml.Marshal(content)
		if err != nil {
			return fmt.Errorf("failed to marshal YAML: %w", err)
		}

	case "text":
		switch v := content.(type) {
		case string:
			data = []byte(v)
		case []byte:
			data = v
		default:
			return fmt.Errorf("text format requires string content")
		}

	default:
		return fmt.Errorf("unsupported format: %s", t.Format)
	}

	// Write file
	if err := os.WriteFile(t.Destination, data, t.Mode); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	t.wroteFile = t.Destination
	return nil
}

func (t *WriteConfigTask) renderContent(ctx *core.InstallContext) any {
	switch content := t.Content.(type) {
	case string:
		return ctx.Render(content)
	case map[string]any:
		return t.renderMap(ctx, content)
	case []any:
		return t.renderSlice(ctx, content)
	default:
		return content
	}
}

func (t *WriteConfigTask) renderMap(ctx *core.InstallContext, m map[string]any) map[string]any {
	result := make(map[string]any)
	for k, v := range m {
		switch val := v.(type) {
		case string:
			result[k] = ctx.Render(val)
		case map[string]any:
			result[k] = t.renderMap(ctx, val)
		case []any:
			result[k] = t.renderSlice(ctx, val)
		default:
			result[k] = val
		}
	}
	return result
}

func (t *WriteConfigTask) renderSlice(ctx *core.InstallContext, s []any) []any {
	result := make([]any, len(s))
	for i, v := range s {
		switch val := v.(type) {
		case string:
			result[i] = ctx.Render(val)
		case map[string]any:
			result[i] = t.renderMap(ctx, val)
		case []any:
			result[i] = t.renderSlice(ctx, val)
		default:
			result[i] = val
		}
	}
	return result
}

// CanRollback returns true if the task can be rolled back.
func (t *WriteConfigTask) CanRollback() bool {
	return t.wroteFile != ""
}

// Rollback restores the previous file content or removes the file.
func (t *WriteConfigTask) Rollback(ctx *core.InstallContext, bus *core.EventBus) error {
	if t.wroteFile == "" {
		return nil
	}

	ctx.AddLog(core.LogInfo, fmt.Sprintf("Rolling back config file: %s", t.wroteFile))

	if t.fileExisted {
		// Restore previous content
		if err := os.WriteFile(t.wroteFile, t.previousData, t.Mode); err != nil {
			return fmt.Errorf("failed to restore config: %w", err)
		}
	} else {
		// Remove created file
		if err := os.Remove(t.wroteFile); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to remove config: %w", err)
		}
	}

	return nil
}
