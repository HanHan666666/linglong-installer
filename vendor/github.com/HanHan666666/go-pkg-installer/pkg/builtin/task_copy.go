// Package builtin provides the copy task implementation.
package builtin

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// CopyTask copies files or directories.
type CopyTask struct {
	core.BaseTask
	Source           string
	Destination      string
	Mode             os.FileMode
	Overwrite        bool
	RequirePrivilege bool

	// For rollback
	copiedFiles []string
	createdDirs []string
}

// RegisterCopyTask registers the copy task factory.
func RegisterCopyTask() {
	core.Tasks.Register("copy", func(config map[string]any, ctx *core.InstallContext) (core.Task, error) {
		mode := os.FileMode(0644)
		if modeVal := getConfigInt(config, "mode", 0); modeVal != 0 {
			mode = os.FileMode(modeVal)
		}

		task := &CopyTask{
			BaseTask: core.BaseTask{
				TaskID:   getConfigString(config, "id"),
				TaskType: "copy",
				Config:   config,
			},
			Source:           ctx.Render(getConfigStringAny(config, "from", "source")),
			Destination:      ctx.Render(getConfigStringAny(config, "to", "destination")),
			Mode:             mode,
			Overwrite:        getConfigBool(config, "overwrite"),
			RequirePrivilege: getConfigBool(config, "requirePrivilege"),
		}

		if task.TaskID == "" {
			task.TaskID = fmt.Sprintf("copy-%s", filepath.Base(task.Source))
		}

		return task, nil
	})
}

// Validate validates the copy task configuration.
func (t *CopyTask) Validate() error {
	if t.Source == "" {
		return errors.New("copy: source is required")
	}
	if t.Destination == "" {
		return errors.New("copy: destination is required")
	}
	return nil
}

// Execute copies the file or directory.
func (t *CopyTask) Execute(ctx *core.InstallContext, bus *core.EventBus) error {
	if err := ensurePrivilege(ctx, t.RequirePrivilege); err != nil {
		return err
	}

	ctx.AddLog(core.LogInfo, fmt.Sprintf("Copying %s to %s", t.Source, t.Destination))

	info, err := os.Stat(t.Source)
	if err != nil {
		return fmt.Errorf("source not found: %w", err)
	}

	if info.IsDir() {
		return t.copyDir(ctx)
	}
	return t.copyFile(ctx, t.Source, t.Destination)
}

func (t *CopyTask) copyFile(ctx *core.InstallContext, src, dst string) error {
	// Check if destination exists
	if _, err := os.Stat(dst); err == nil && !t.Overwrite {
		return fmt.Errorf("destination already exists: %s", dst)
	}

	// Ensure parent directory exists
	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source: %w", err)
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, t.Mode)
	if err != nil {
		return fmt.Errorf("failed to create destination: %w", err)
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return fmt.Errorf("failed to copy: %w", err)
	}

	t.copiedFiles = append(t.copiedFiles, dst)
	return nil
}

func (t *CopyTask) copyDir(ctx *core.InstallContext) error {
	return filepath.Walk(t.Source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calculate relative path
		relPath, err := filepath.Rel(t.Source, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(t.Destination, relPath)

		if info.IsDir() {
			if err := os.MkdirAll(dstPath, info.Mode()); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
			t.createdDirs = append(t.createdDirs, dstPath)
			return nil
		}

		return t.copyFile(ctx, path, dstPath)
	})
}

// CanRollback returns true if the task can be rolled back.
func (t *CopyTask) CanRollback() bool {
	return len(t.copiedFiles) > 0 || len(t.createdDirs) > 0
}

// Rollback removes copied files and directories.
func (t *CopyTask) Rollback(ctx *core.InstallContext, bus *core.EventBus) error {
	ctx.AddLog(core.LogInfo, "Rolling back copied files")

	// Remove files first
	for _, file := range t.copiedFiles {
		if err := os.Remove(file); err != nil && !os.IsNotExist(err) {
			ctx.AddLog(core.LogWarn, fmt.Sprintf("Failed to remove file: %s", file))
		}
	}

	// Remove directories in reverse order
	for i := len(t.createdDirs) - 1; i >= 0; i-- {
		dir := t.createdDirs[i]
		if err := os.Remove(dir); err != nil && !os.IsNotExist(err) {
			ctx.AddLog(core.LogWarn, fmt.Sprintf("Failed to remove directory: %s", dir))
		}
	}

	return nil
}
