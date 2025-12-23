// Package builtin provides the remove_desktop_entry task implementation.
package builtin

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// RemoveDesktopEntryTask removes a .desktop file.
type RemoveDesktopEntryTask struct {
	core.BaseTask
	Name             string
	Path             string
	RequirePrivilege bool
}

// RegisterRemoveDesktopEntryTask registers the remove_desktop_entry task factory.
func RegisterRemoveDesktopEntryTask() {
	core.Tasks.Register("removeDesktopEntry", func(config map[string]any, ctx *core.InstallContext) (core.Task, error) {
		task := &RemoveDesktopEntryTask{
			BaseTask: core.BaseTask{
				TaskID:   getConfigString(config, "id"),
				TaskType: "removeDesktopEntry",
				Config:   config,
			},
			Name:             ctx.Render(getConfigString(config, "name")),
			Path:             ctx.Render(getConfigString(config, "path")),
			RequirePrivilege: getConfigBool(config, "requirePrivilege"),
		}

		if task.TaskID == "" {
			task.TaskID = fmt.Sprintf("remove-desktop-%s", strings.ToLower(strings.ReplaceAll(task.Name, " ", "-")))
		}

		return task, nil
	})
}

// Validate validates the remove_desktop_entry task configuration.
func (t *RemoveDesktopEntryTask) Validate() error {
	if t.Name == "" && t.Path == "" {
		return errors.New("removeDesktopEntry: name or path is required")
	}
	return nil
}

// Execute removes the desktop entry file.
func (t *RemoveDesktopEntryTask) Execute(ctx *core.InstallContext, bus *core.EventBus) error {
	if err := ensurePrivilege(ctx, t.RequirePrivilege); err != nil {
		return err
	}

	paths := t.resolvePaths()
	for _, path := range paths {
		if err := removeFile(path); err != nil {
			return err
		}
		ctx.AddLog(core.LogInfo, fmt.Sprintf("Removed desktop entry: %s", path))
		return nil
	}

	return fmt.Errorf("desktop entry not found for %s", t.Name)
}

// CanRollback returns false by default.
func (t *RemoveDesktopEntryTask) CanRollback() bool {
	return false
}

// Rollback does nothing for remove operations.
func (t *RemoveDesktopEntryTask) Rollback(ctx *core.InstallContext, bus *core.EventBus) error {
	return nil
}

func (t *RemoveDesktopEntryTask) resolvePaths() []string {
	if t.Path != "" {
		return []string{t.Path}
	}

	filename := strings.ToLower(strings.ReplaceAll(t.Name, " ", "-")) + ".desktop"
	paths := []string{}

	if home, err := os.UserHomeDir(); err == nil {
		paths = append(paths, filepath.Join(home, ".local", "share", "applications", filename))
	}
	paths = append(paths, filepath.Join("/usr/share/applications", filename))
	return paths
}

func removeFile(path string) error {
	if path == "" {
		return nil
	}
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return os.Remove(path)
}
