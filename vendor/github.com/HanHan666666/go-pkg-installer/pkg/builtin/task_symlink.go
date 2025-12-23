// Package builtin provides the symlink task implementation.
package builtin

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// SymlinkTask creates a symbolic link.
type SymlinkTask struct {
	core.BaseTask
	Target           string
	LinkPath         string
	Overwrite        bool
	RequirePrivilege bool

	// For rollback
	createdLink string
	oldTarget   string
}

// RegisterSymlinkTask registers the symlink task factory.
func RegisterSymlinkTask() {
	core.Tasks.Register("symlink", func(config map[string]any, ctx *core.InstallContext) (core.Task, error) {
		task := &SymlinkTask{
			BaseTask: core.BaseTask{
				TaskID:   getConfigString(config, "id"),
				TaskType: "symlink",
				Config:   config,
			},
			Target:           ctx.Render(getConfigString(config, "target")),
			LinkPath:         ctx.Render(getConfigString(config, "link")),
			Overwrite:        getConfigBool(config, "overwrite"),
			RequirePrivilege: getConfigBool(config, "requirePrivilege"),
		}

		if task.TaskID == "" {
			task.TaskID = fmt.Sprintf("symlink-%s", filepath.Base(task.LinkPath))
		}

		return task, nil
	})
}

// Validate validates the symlink task configuration.
func (t *SymlinkTask) Validate() error {
	if t.Target == "" {
		return errors.New("symlink: target is required")
	}
	if t.LinkPath == "" {
		return errors.New("symlink: link path is required")
	}
	return nil
}

// Execute creates the symbolic link.
func (t *SymlinkTask) Execute(ctx *core.InstallContext, bus *core.EventBus) error {
	if err := ensurePrivilege(ctx, t.RequirePrivilege); err != nil {
		return err
	}

	ctx.AddLog(core.LogInfo, fmt.Sprintf("Creating symlink %s -> %s", t.LinkPath, t.Target))

	// Ensure parent directory exists
	linkDir := filepath.Dir(t.LinkPath)
	if err := os.MkdirAll(linkDir, 0755); err != nil {
		return fmt.Errorf("failed to create parent directory: %w", err)
	}

	// Check if link already exists
	if info, err := os.Lstat(t.LinkPath); err == nil {
		if info.Mode()&os.ModeSymlink != 0 {
			// It's a symlink, save old target for rollback
			oldTarget, err := os.Readlink(t.LinkPath)
			if err == nil {
				t.oldTarget = oldTarget
			}
		}

		if !t.Overwrite {
			return fmt.Errorf("link already exists: %s", t.LinkPath)
		}

		// Remove existing file/link
		if err := os.Remove(t.LinkPath); err != nil {
			return fmt.Errorf("failed to remove existing link: %w", err)
		}
	}

	// Create the symlink
	if err := os.Symlink(t.Target, t.LinkPath); err != nil {
		return fmt.Errorf("failed to create symlink: %w", err)
	}

	t.createdLink = t.LinkPath
	return nil
}

// CanRollback returns true if the task can be rolled back.
func (t *SymlinkTask) CanRollback() bool {
	return t.createdLink != ""
}

// Rollback removes the symlink or restores the old target.
func (t *SymlinkTask) Rollback(ctx *core.InstallContext, bus *core.EventBus) error {
	if t.createdLink == "" {
		return nil
	}

	ctx.AddLog(core.LogInfo, fmt.Sprintf("Rolling back symlink: %s", t.createdLink))

	if err := os.Remove(t.createdLink); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove symlink: %w", err)
	}

	// Restore old symlink if there was one
	if t.oldTarget != "" {
		if err := os.Symlink(t.oldTarget, t.createdLink); err != nil {
			ctx.AddLog(core.LogWarn, fmt.Sprintf("Failed to restore old symlink: %v", err))
		}
	}

	return nil
}
