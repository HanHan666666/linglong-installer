// Package builtin provides the unpack task implementation.
package builtin

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// UnpackTask extracts an archive file.
type UnpackTask struct {
	core.BaseTask
	Source           string
	Destination      string
	StripPrefix      int
	RequirePrivilege bool

	// For rollback
	createdFiles []string
	createdDirs  []string
}

// RegisterUnpackTask registers the unpack task factory.
func RegisterUnpackTask() {
	core.Tasks.Register("unpack", func(config map[string]any, ctx *core.InstallContext) (core.Task, error) {
		task := &UnpackTask{
			BaseTask: core.BaseTask{
				TaskID:   getConfigString(config, "id"),
				TaskType: "unpack",
				Config:   config,
			},
			Source:           ctx.Render(getConfigStringAny(config, "from", "source")),
			Destination:      ctx.Render(getConfigStringAny(config, "to", "destination")),
			StripPrefix:      getConfigIntAny(config, 0, "stripPrefix", "strip_prefix"),
			RequirePrivilege: getConfigBool(config, "requirePrivilege"),
		}

		if task.TaskID == "" {
			task.TaskID = fmt.Sprintf("unpack-%s", filepath.Base(task.Source))
		}

		return task, nil
	})
}

// Validate validates the unpack task configuration.
func (t *UnpackTask) Validate() error {
	if t.Source == "" {
		return errors.New("unpack: source is required")
	}
	if t.Destination == "" {
		return errors.New("unpack: destination is required")
	}
	return nil
}

// Execute extracts the archive.
func (t *UnpackTask) Execute(ctx *core.InstallContext, bus *core.EventBus) error {
	if err := ensurePrivilege(ctx, t.RequirePrivilege); err != nil {
		return err
	}

	ctx.AddLog(core.LogInfo, fmt.Sprintf("Unpacking %s to %s", t.Source, t.Destination))

	// Ensure destination directory exists
	if err := os.MkdirAll(t.Destination, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Detect archive type
	ext := strings.ToLower(filepath.Ext(t.Source))

	// Handle .tar.gz, .tgz
	if strings.HasSuffix(strings.ToLower(t.Source), ".tar.gz") || ext == ".tgz" {
		return t.extractTarGz(ctx, bus)
	}

	// Handle .tar
	if ext == ".tar" {
		return t.extractTar(ctx, bus)
	}

	// Handle .zip
	if ext == ".zip" {
		return t.extractZip(ctx, bus)
	}

	return fmt.Errorf("unsupported archive format: %s", ext)
}

func (t *UnpackTask) extractTarGz(ctx *core.InstallContext, bus *core.EventBus) error {
	file, err := os.Open(t.Source)
	if err != nil {
		return fmt.Errorf("failed to open archive: %w", err)
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzr.Close()

	return t.extractTarReader(tar.NewReader(gzr), ctx, bus)
}

func (t *UnpackTask) extractTar(ctx *core.InstallContext, bus *core.EventBus) error {
	file, err := os.Open(t.Source)
	if err != nil {
		return fmt.Errorf("failed to open archive: %w", err)
	}
	defer file.Close()

	return t.extractTarReader(tar.NewReader(file), ctx, bus)
}

func (t *UnpackTask) extractTarReader(tr *tar.Reader, ctx *core.InstallContext, bus *core.EventBus) error {
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read tar: %w", err)
		}

		// Apply strip prefix
		name := t.stripPath(header.Name)
		if name == "" {
			continue
		}

		// Security: prevent path traversal
		target := filepath.Join(t.Destination, name)
		if !strings.HasPrefix(target, filepath.Clean(t.Destination)+string(os.PathSeparator)) {
			return fmt.Errorf("invalid path in archive: %s", header.Name)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, os.FileMode(header.Mode)); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
			t.createdDirs = append(t.createdDirs, target)

		case tar.TypeReg:
			// Ensure parent directory exists
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return fmt.Errorf("failed to create parent directory: %w", err)
			}

			outFile, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(header.Mode))
			if err != nil {
				return fmt.Errorf("failed to create file: %w", err)
			}

			if _, err := io.Copy(outFile, tr); err != nil {
				outFile.Close()
				return fmt.Errorf("failed to write file: %w", err)
			}
			outFile.Close()
			t.createdFiles = append(t.createdFiles, target)

		case tar.TypeSymlink:
			// Ensure parent directory exists
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return fmt.Errorf("failed to create parent directory: %w", err)
			}

			if err := os.Symlink(header.Linkname, target); err != nil {
				return fmt.Errorf("failed to create symlink: %w", err)
			}
			t.createdFiles = append(t.createdFiles, target)
		}
	}

	ctx.AddLog(core.LogInfo, fmt.Sprintf("Extracted %d files", len(t.createdFiles)))
	return nil
}

func (t *UnpackTask) extractZip(ctx *core.InstallContext, bus *core.EventBus) error {
	r, err := zip.OpenReader(t.Source)
	if err != nil {
		return fmt.Errorf("failed to open zip: %w", err)
	}
	defer r.Close()

	for _, f := range r.File {
		// Apply strip prefix
		name := t.stripPath(f.Name)
		if name == "" {
			continue
		}

		// Security: prevent path traversal
		target := filepath.Join(t.Destination, name)
		if !strings.HasPrefix(target, filepath.Clean(t.Destination)+string(os.PathSeparator)) {
			return fmt.Errorf("invalid path in archive: %s", f.Name)
		}

		if f.FileInfo().IsDir() {
			// Use at least 0755 to ensure directories are accessible
			mode := f.Mode() | 0755
			if err := os.MkdirAll(target, mode); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
			t.createdDirs = append(t.createdDirs, target)
			continue
		}

		// Ensure parent directory exists
		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			return fmt.Errorf("failed to create parent directory: %w", err)
		}

		outFile, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, f.Mode())
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return fmt.Errorf("failed to open zip entry: %w", err)
		}

		_, err = io.Copy(outFile, rc)
		rc.Close()
		outFile.Close()
		if err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}

		t.createdFiles = append(t.createdFiles, target)
	}

	ctx.AddLog(core.LogInfo, fmt.Sprintf("Extracted %d files", len(t.createdFiles)))
	return nil
}

func (t *UnpackTask) stripPath(path string) string {
	if t.StripPrefix <= 0 {
		return path
	}

	parts := strings.Split(path, "/")
	if len(parts) <= t.StripPrefix {
		return ""
	}
	return strings.Join(parts[t.StripPrefix:], "/")
}

// CanRollback returns true if the task can be rolled back.
func (t *UnpackTask) CanRollback() bool {
	return len(t.createdFiles) > 0 || len(t.createdDirs) > 0
}

// Rollback removes extracted files and directories.
func (t *UnpackTask) Rollback(ctx *core.InstallContext, bus *core.EventBus) error {
	ctx.AddLog(core.LogInfo, "Rolling back unpacked files")

	// Remove files first
	for _, file := range t.createdFiles {
		if err := os.Remove(file); err != nil && !os.IsNotExist(err) {
			ctx.AddLog(core.LogWarn, fmt.Sprintf("Failed to remove file: %s", file))
		}
	}

	// Remove directories in reverse order (deepest first)
	for i := len(t.createdDirs) - 1; i >= 0; i-- {
		dir := t.createdDirs[i]
		if err := os.RemoveAll(dir); err != nil && !os.IsNotExist(err) {
			ctx.AddLog(core.LogWarn, fmt.Sprintf("Failed to remove directory: %s", dir))
		}
	}

	return nil
}
