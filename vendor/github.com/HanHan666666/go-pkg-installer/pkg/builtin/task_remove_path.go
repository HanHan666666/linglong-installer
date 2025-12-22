// Package builtin provides the remove_path task implementation.
package builtin

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// RemovePathTask removes a file or directory.
type RemovePathTask struct {
	core.BaseTask
	Path             string
	Recursive        bool
	Force            bool // Ignore errors if path doesn't exist
	UserData         bool
	RequirePrivilege bool

	// For rollback - we don't support full rollback for remove
	// but we track what was removed for logging
	removedPath string
}

// RegisterRemovePathTask registers the remove_path task factory.
func RegisterRemovePathTask() {
	core.Tasks.Register("removePath", func(config map[string]any, ctx *core.InstallContext) (core.Task, error) {
		task := &RemovePathTask{
			BaseTask: core.BaseTask{
				TaskID:   getConfigString(config, "id"),
				TaskType: "removePath",
				Config:   config,
			},
			Path:             ctx.Render(getConfigString(config, "path")),
			Recursive:        getConfigBool(config, "recursive"),
			Force:            getConfigBool(config, "force"),
			UserData:         getConfigBool(config, "userData"),
			RequirePrivilege: getConfigBoolDefault(config, "requirePrivilege", false),
		}

		if task.TaskID == "" {
			task.TaskID = fmt.Sprintf("removePath-%s", filepath.Base(task.Path))
		}

		return task, nil
	})
}

// Validate validates the remove_path task configuration.
func (t *RemovePathTask) Validate() error {
	if t.Path == "" {
		return errors.New("removePath: path is required")
	}
	return nil
}

// Execute removes the file or directory.
func (t *RemovePathTask) Execute(ctx *core.InstallContext, bus *core.EventBus) error {
	if t.UserData {
		if keep, ok := ctx.Get("uninstall.keepUserData"); ok {
			if keepBool, ok := keep.(bool); ok && keepBool {
				ctx.AddLog(core.LogInfo, fmt.Sprintf("Skipping user data removal: %s", t.Path))
				return nil
			}
		}
	}

	if err := ensurePrivilege(ctx, t.RequirePrivilege); err != nil {
		return err
	}

	ctx.AddLog(core.LogInfo, fmt.Sprintf("Removing: %s", t.Path))

	// Check if path exists
	info, err := os.Lstat(t.Path)
	if err != nil {
		if os.IsNotExist(err) {
			if t.Force {
				ctx.AddLog(core.LogInfo, "Path does not exist, skipping")
				return nil
			}
			return fmt.Errorf("path does not exist: %s", t.Path)
		}
		return fmt.Errorf("failed to stat path: %w", err)
	}

	// If it's a directory and not recursive, check if empty
	if info.IsDir() && !t.Recursive {
		entries, err := os.ReadDir(t.Path)
		if err != nil {
			return fmt.Errorf("failed to read directory: %w", err)
		}
		if len(entries) > 0 {
			return fmt.Errorf("directory not empty: %s (use recursive: true)", t.Path)
		}
	}

	// Remove
	var removeErr error
	if t.Recursive {
		removeErr = os.RemoveAll(t.Path)
	} else {
		removeErr = os.Remove(t.Path)
	}

	if removeErr != nil {
		if t.Force && os.IsNotExist(removeErr) {
			return nil
		}
		return fmt.Errorf("failed to remove: %w", removeErr)
	}

	t.removedPath = t.Path
	ctx.AddLog(core.LogInfo, fmt.Sprintf("Removed: %s", t.Path))

	return nil
}

// CanRollback returns false as we can't restore deleted files.
func (t *RemovePathTask) CanRollback() bool {
	return false
}

// Rollback does nothing for remove operations.
func (t *RemovePathTask) Rollback(ctx *core.InstallContext, bus *core.EventBus) error {
	if t.removedPath != "" {
		ctx.AddLog(core.LogWarn, fmt.Sprintf("Cannot restore removed path: %s", t.removedPath))
	}
	return nil
}
