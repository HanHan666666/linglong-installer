// Package builtin provides the desktop_entry task implementation.
package builtin

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// DesktopEntryTask creates a .desktop file for Linux desktop integration.
type DesktopEntryTask struct {
	core.BaseTask
	Name             string
	Exec             string
	Icon             string
	Comment          string
	Categories       []string
	Terminal         bool
	EntryType        string // Application, Link, Directory
	StartupWMClass   string
	MimeTypes        []string
	Keywords         []string
	Destination      string
	RequirePrivilege bool

	// For rollback
	createdFile string
}

// RegisterDesktopEntryTask registers the desktop_entry task factory.
func RegisterDesktopEntryTask() {
	factory := func(config map[string]any, ctx *core.InstallContext) (core.Task, error) {
		taskType := getConfigString(config, "type")
		if taskType == "" {
			taskType = "desktopEntry"
		}

		task := &DesktopEntryTask{
			BaseTask: core.BaseTask{
				TaskID:   getConfigString(config, "id"),
				TaskType: taskType,
				Config:   config,
			},
			Name:             ctx.Render(getConfigString(config, "name")),
			Exec:             ctx.Render(getConfigString(config, "exec")),
			Icon:             ctx.Render(getConfigString(config, "icon")),
			Comment:          ctx.Render(getConfigString(config, "comment")),
			Categories:       getConfigStringSlice(config, "categories"),
			Terminal:         getConfigBool(config, "terminal"),
			EntryType:        getConfigStringAny(config, "entryType", "entry_type"),
			StartupWMClass:   ctx.Render(getConfigString(config, "startup_wm_class")),
			MimeTypes:        getConfigStringSlice(config, "mime_types"),
			Keywords:         getConfigStringSlice(config, "keywords"),
			Destination:      ctx.Render(getConfigString(config, "destination")),
			RequirePrivilege: getConfigBool(config, "requirePrivilege"),
		}

		if task.TaskID == "" {
			task.TaskID = fmt.Sprintf("desktopEntry-%s", strings.ToLower(strings.ReplaceAll(task.Name, " ", "-")))
		}

		// Default type is Application
		if task.EntryType == "" {
			task.EntryType = "Application"
		}

		// Default destination to user's applications directory
		if task.Destination == "" {
			home, _ := os.UserHomeDir()
			filename := strings.ToLower(strings.ReplaceAll(task.Name, " ", "-")) + ".desktop"
			task.Destination = filepath.Join(home, ".local", "share", "applications", filename)
		}

		return task, nil
	}

	core.Tasks.Register("desktopEntry", factory)
	core.Tasks.Register("createDesktopEntry", factory)
}

// Validate validates the desktop_entry task configuration.
func (t *DesktopEntryTask) Validate() error {
	if t.Name == "" {
		return errors.New("desktopEntry: name is required")
	}
	if t.Exec == "" {
		return errors.New("desktopEntry: exec is required")
	}
	return nil
}

// Execute creates the .desktop file.
func (t *DesktopEntryTask) Execute(ctx *core.InstallContext, bus *core.EventBus) error {
	if err := ensurePrivilege(ctx, t.RequirePrivilege); err != nil {
		return err
	}

	ctx.AddLog(core.LogInfo, fmt.Sprintf("Creating desktop entry: %s", t.Destination))

	// Ensure parent directory exists
	destDir := filepath.Dir(t.Destination)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Build desktop entry content
	var sb strings.Builder
	sb.WriteString("[Desktop Entry]\n")
	sb.WriteString(fmt.Sprintf("Type=%s\n", t.EntryType))
	sb.WriteString(fmt.Sprintf("Name=%s\n", t.Name))
	sb.WriteString(fmt.Sprintf("Exec=%s\n", t.Exec))

	if t.Icon != "" {
		sb.WriteString(fmt.Sprintf("Icon=%s\n", t.Icon))
	}

	if t.Comment != "" {
		sb.WriteString(fmt.Sprintf("Comment=%s\n", t.Comment))
	}

	if len(t.Categories) > 0 {
		sb.WriteString(fmt.Sprintf("Categories=%s;\n", strings.Join(t.Categories, ";")))
	}

	if t.Terminal {
		sb.WriteString("Terminal=true\n")
	} else {
		sb.WriteString("Terminal=false\n")
	}

	if t.StartupWMClass != "" {
		sb.WriteString(fmt.Sprintf("StartupWMClass=%s\n", t.StartupWMClass))
	}

	if len(t.MimeTypes) > 0 {
		sb.WriteString(fmt.Sprintf("MimeType=%s;\n", strings.Join(t.MimeTypes, ";")))
	}

	if len(t.Keywords) > 0 {
		sb.WriteString(fmt.Sprintf("Keywords=%s;\n", strings.Join(t.Keywords, ";")))
	}

	// Write file
	if err := os.WriteFile(t.Destination, []byte(sb.String()), 0755); err != nil {
		return fmt.Errorf("failed to write desktop entry: %w", err)
	}

	t.createdFile = t.Destination
	ctx.AddLog(core.LogInfo, "Desktop entry created successfully")

	return nil
}

// CanRollback returns true if the task can be rolled back.
func (t *DesktopEntryTask) CanRollback() bool {
	return t.createdFile != ""
}

// Rollback removes the created desktop entry.
func (t *DesktopEntryTask) Rollback(ctx *core.InstallContext, bus *core.EventBus) error {
	if t.createdFile == "" {
		return nil
	}

	ctx.AddLog(core.LogInfo, fmt.Sprintf("Removing desktop entry: %s", t.createdFile))

	if err := os.Remove(t.createdFile); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove desktop entry: %w", err)
	}

	return nil
}
