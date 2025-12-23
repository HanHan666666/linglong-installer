// Package builtin provides the download task implementation.
package builtin

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// DownloadTask downloads a file from a URL.
type DownloadTask struct {
	core.BaseTask
	URL              string
	Destination      string
	SHA256           string
	Timeout          time.Duration
	Headers          map[string]string
	RequirePrivilege bool

	// For rollback
	downloadedFile string
}

// RegisterDownloadTask registers the download task factory.
func RegisterDownloadTask() {
	core.Tasks.Register("download", func(config map[string]any, ctx *core.InstallContext) (core.Task, error) {
		headers := make(map[string]string)
		if headerMap, ok := config["headers"].(map[string]any); ok {
			for k, v := range headerMap {
				if s, ok := v.(string); ok {
					headers[k] = ctx.Render(s)
				}
			}
		}

		task := &DownloadTask{
			BaseTask: core.BaseTask{
				TaskID:   getConfigString(config, "id"),
				TaskType: "download",
				Config:   config,
			},
			URL:              ctx.Render(getConfigString(config, "url")),
			Destination:      ctx.Render(getConfigStringAny(config, "to", "destination")),
			SHA256:           getConfigString(config, "sha256"),
			Timeout:          time.Duration(getConfigIntAny(config, 300, "timeoutSec", "timeout")) * time.Second,
			Headers:          headers,
			RequirePrivilege: getConfigBool(config, "requirePrivilege"),
		}

		if task.TaskID == "" {
			task.TaskID = fmt.Sprintf("download-%s", filepath.Base(task.URL))
		}

		if task.Destination == "" && task.URL != "" {
			task.Destination = filepath.Join(os.TempDir(), filepath.Base(task.URL))
		}

		return task, nil
	})
}

// Validate validates the download task configuration.
func (t *DownloadTask) Validate() error {
	if t.URL == "" {
		return errors.New("download: url is required")
	}
	if t.Destination == "" {
		return errors.New("download: destination is required")
	}
	return nil
}

// Execute downloads the file.
func (t *DownloadTask) Execute(ctx *core.InstallContext, bus *core.EventBus) error {
	if err := ensurePrivilege(ctx, t.RequirePrivilege); err != nil {
		return err
	}

	ctx.AddLog(core.LogInfo, fmt.Sprintf("Downloading %s to %s", t.URL, t.Destination))

	// Ensure destination directory exists
	destDir := filepath.Dir(t.Destination)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: t.Timeout,
	}

	// Start download
	req, err := http.NewRequest(http.MethodGet, t.URL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	for k, v := range t.Headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status: %s", resp.Status)
	}

	// Create destination file
	out, err := os.Create(t.Destination)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer out.Close()

	// Track progress if content length is known
	var reader io.Reader = resp.Body
	contentLength := resp.ContentLength
	if contentLength > 0 {
		reader = &progressReader{
			reader:       resp.Body,
			total:        contentLength,
			bus:          bus,
			taskID:       t.TaskID,
			lastProgress: -1,
		}
	}

	// Download with hash calculation
	hasher := sha256.New()
	writer := io.MultiWriter(out, hasher)

	_, err = io.Copy(writer, reader)
	if err != nil {
		os.Remove(t.Destination)
		return fmt.Errorf("download failed: %w", err)
	}

	// Verify checksum if provided
	if t.SHA256 != "" {
		actualHash := hex.EncodeToString(hasher.Sum(nil))
		if actualHash != t.SHA256 {
			os.Remove(t.Destination)
			return fmt.Errorf("checksum mismatch: expected %s, got %s", t.SHA256, actualHash)
		}
		ctx.AddLog(core.LogInfo, "Checksum verified")
	}

	t.downloadedFile = t.Destination
	ctx.AddLog(core.LogInfo, fmt.Sprintf("Downloaded %s successfully", t.URL))

	return nil
}

// CanRollback returns true if the task can be rolled back.
func (t *DownloadTask) CanRollback() bool {
	return t.downloadedFile != ""
}

// Rollback removes the downloaded file.
func (t *DownloadTask) Rollback(ctx *core.InstallContext, bus *core.EventBus) error {
	if t.downloadedFile != "" {
		ctx.AddLog(core.LogInfo, fmt.Sprintf("Removing downloaded file: %s", t.downloadedFile))
		if err := os.Remove(t.downloadedFile); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to remove downloaded file: %w", err)
		}
	}
	return nil
}

// progressReader wraps an io.Reader and reports progress.
type progressReader struct {
	reader       io.Reader
	total        int64
	current      int64
	bus          *core.EventBus
	taskID       string
	lastProgress int // Last reported progress percentage
}

func (r *progressReader) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	r.current += int64(n)

	// Calculate progress percentage
	progress := int(float64(r.current) / float64(r.total) * 100)

	// Only report if progress changed by at least 1%
	if progress != r.lastProgress {
		r.lastProgress = progress
		r.bus.PublishProgress(
			r.taskID,
			float64(r.current)/float64(r.total),
			fmt.Sprintf("Downloading: %d%%", progress),
		)
	}

	return n, err
}

// Helper functions for config access
func getConfigString(config map[string]any, key string) string {
	if v, ok := config[key].(string); ok {
		return v
	}
	return ""
}

func getConfigStringAny(config map[string]any, keys ...string) string {
	for _, key := range keys {
		if v, ok := config[key].(string); ok && v != "" {
			return v
		}
	}
	return ""
}

func getConfigInt(config map[string]any, key string, defaultVal int) int {
	switch v := config[key].(type) {
	case int:
		return v
	case int64:
		return int(v)
	case float64:
		return int(v)
	default:
		return defaultVal
	}
}

func getConfigIntAny(config map[string]any, defaultVal int, keys ...string) int {
	for _, key := range keys {
		if _, ok := config[key]; ok {
			return getConfigInt(config, key, defaultVal)
		}
	}
	return defaultVal
}

func getConfigBool(config map[string]any, key string) bool {
	if v, ok := config[key].(bool); ok {
		return v
	}
	return false
}

func getConfigBoolAny(config map[string]any, keys ...string) bool {
	for _, key := range keys {
		if _, ok := config[key]; ok {
			return getConfigBool(config, key)
		}
	}
	return false
}

func getConfigBoolDefault(config map[string]any, key string, defaultVal bool) bool {
	if _, ok := config[key]; ok {
		return getConfigBool(config, key)
	}
	return defaultVal
}

func getConfigStringSlice(config map[string]any, key string) []string {
	switch v := config[key].(type) {
	case []string:
		return v
	case []interface{}:
		result := make([]string, 0, len(v))
		for _, item := range v {
			if s, ok := item.(string); ok {
				result = append(result, s)
			}
		}
		return result
	default:
		return nil
	}
}

func parseFileMode(value any) int {
	switch v := value.(type) {
	case int:
		return v
	case int64:
		return int(v)
	case float64:
		return int(v)
	case string:
		if v == "" {
			return 0
		}
		if parsed, err := strconv.ParseInt(v, 8, 32); err == nil {
			return int(parsed)
		}
	}
	return 0
}
