// Package core provides the workflow engine, context, and runtime for the installer framework.
package core

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

// InstallContext holds all state for an installation session.
// It is thread-safe and can be accessed from multiple goroutines.
type InstallContext struct {
	mu sync.RWMutex

	// Env contains environment detection results (read-only after init)
	Env EnvInfo

	// UserInput contains user-provided values from forms
	UserInput map[string]any

	// Plan contains the parsed task plan for display
	Plan *TaskPlan

	// Runtime contains current execution state
	Runtime RuntimeState

	// Meta contains arbitrary metadata from config
	Meta map[string]any

	// Event bus for log propagation
	bus *EventBus

	// Log output file
	logFile *os.File
	logPath string
}

// EnvInfo contains detected environment information.
type EnvInfo struct {
	// OS information
	Distro        string // e.g., "ubuntu", "fedora", "arch"
	DistroVersion string // e.g., "22.04", "39"
	Arch          string // e.g., "x86_64", "arm64"
	Desktop       string // e.g., "gnome", "kde", "unknown"

	// Permission state
	IsRoot    bool
	HasPolkit bool
	HasSudo   bool

	// System resources
	DiskFreeMB int64

	// Installation detection
	InstalledVersion string
	InstallDir       string
}

// TaskPlan represents the planned tasks for display in summary.
type TaskPlan struct {
	Tasks []TaskSummary
}

// TaskSummary is a human-readable summary of a task.
type TaskSummary struct {
	Type         string
	Description  string
	RequiresRoot bool
}

// RuntimeState contains current execution state.
type RuntimeState struct {
	Action      string // "install", "uninstall", "upgrade", "repair"
	CurrentStep string
	FlowID      string
	Progress    float64
	Logs        []LogEntry
	Errors      []error
	StartTime   int64
	Completed   bool
}

// LogEntry represents a single log message.
type LogEntry struct {
	Level   LogLevel
	Message string
	Time    int64
}

// LogLevel represents log severity.
type LogLevel int

const (
	LogInfo LogLevel = iota
	LogWarn
	LogError
)

// NewInstallContext creates a new empty InstallContext.
func NewInstallContext() *InstallContext {
	return &InstallContext{
		UserInput: make(map[string]any),
		Meta:      make(map[string]any),
		Runtime: RuntimeState{
			Logs:   make([]LogEntry, 0),
			Errors: make([]error, 0),
		},
	}
}

// SetEventBus attaches an EventBus to the context for log propagation.
func (c *InstallContext) SetEventBus(bus *EventBus) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.bus = bus
}

// SetLogFile configures a log file path for persistent logs.
func (c *InstallContext) SetLogFile(path string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.logFile != nil {
		_ = c.logFile.Close()
		c.logFile = nil
	}

	if path == "" {
		c.logPath = ""
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	c.logFile = file
	c.logPath = path
	return nil
}

// LogPath returns the active log file path, if any.
func (c *InstallContext) LogPath() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.logPath
}

// CloseLogFile closes the log file if configured.
func (c *InstallContext) CloseLogFile() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.logFile != nil {
		_ = c.logFile.Close()
		c.logFile = nil
	}
}

// Get retrieves a value by dot-notation path (e.g., "install.dir", "license.accepted").
// It searches in order: UserInput, Meta, Env (as map).
func (c *InstallContext) Get(path string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.getUnlocked(path)
}

func (c *InstallContext) getUnlocked(path string) (any, bool) {
	// Try UserInput first
	if val, ok := getNestedValue(c.UserInput, path); ok {
		return val, true
	}

	// Try Meta
	if val, ok := getNestedValue(c.Meta, path); ok {
		return val, true
	}

	// Try Env fields
	if val, ok := c.getEnvField(path); ok {
		return val, true
	}

	return nil, false
}

// GetString retrieves a string value by path, returning empty string if not found or not a string.
func (c *InstallContext) GetString(path string) string {
	val, ok := c.Get(path)
	if !ok {
		return ""
	}
	if s, ok := val.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", val)
}

// GetBool retrieves a boolean value by path, returning false if not found or not a bool.
func (c *InstallContext) GetBool(path string) bool {
	val, ok := c.Get(path)
	if !ok {
		return false
	}
	if b, ok := val.(bool); ok {
		return b
	}
	return false
}

// GetInt retrieves an integer value by path, returning 0 if not found or not numeric.
func (c *InstallContext) GetInt(path string) int64 {
	val, ok := c.Get(path)
	if !ok {
		return 0
	}
	switch v := val.(type) {
	case int:
		return int64(v)
	case int64:
		return v
	case float64:
		return int64(v)
	default:
		return 0
	}
}

// Set sets a value by dot-notation path in UserInput.
func (c *InstallContext) Set(path string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	setNestedValue(c.UserInput, path, value)
}

// SetMeta sets a value in Meta by dot-notation path.
func (c *InstallContext) SetMeta(path string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	setNestedValue(c.Meta, path, value)
}

// Render replaces ${path} placeholders in a template string with context values.
func (c *InstallContext) Render(template string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := template

	// Support ${path} syntax
	reDollar := regexp.MustCompile(`\$\{([^}]+)\}`)
	result = reDollar.ReplaceAllStringFunc(result, func(match string) string {
		path := match[2 : len(match)-1]
		if val, ok := c.getUnlocked(path); ok {
			return fmt.Sprintf("%v", val)
		}
		return match
	})

	// Support {{.path}} syntax (Go template style)
	reMustache := regexp.MustCompile(`\{\{\.([^}]+)\}\}`)
	result = reMustache.ReplaceAllStringFunc(result, func(match string) string {
		path := match[3 : len(match)-2]
		if val, ok := c.getUnlocked(path); ok {
			return fmt.Sprintf("%v", val)
		}
		return match
	})

	return result
}

// RenderOrDefault returns the rendered value at path, or the default if not found.
func (c *InstallContext) RenderOrDefault(path string, defaultVal string) string {
	val, ok := c.Get(path)
	if !ok || val == nil {
		return defaultVal
	}
	return fmt.Sprintf("%v", val)
}

// AddLog adds a log entry to the runtime state.
func (c *InstallContext) AddLog(level LogLevel, message string) {
	c.mu.Lock()
	bus := c.bus
	writer := c.logFile
	c.mu.Unlock()

	entry := LogEntry{
		Level:   level,
		Message: message,
		Time:    currentTimeMillis(),
	}

	c.mu.Lock()
	c.Runtime.Logs = append(c.Runtime.Logs, entry)
	c.mu.Unlock()

	if bus != nil {
		bus.PublishLog(level, message)
	}

	if writer != nil {
		writeLogEntry(writer, entry)
	}
}

// AddError adds an error to the runtime state.
func (c *InstallContext) AddError(err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Runtime.Errors = append(c.Runtime.Errors, err)
}

// SetProgress updates the current progress (0.0 to 1.0).
func (c *InstallContext) SetProgress(progress float64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Runtime.Progress = progress
}

// getEnvField retrieves a field from EnvInfo by name.
func (c *InstallContext) getEnvField(path string) (any, bool) {
	// Handle env.* prefix
	if strings.HasPrefix(path, "env.") {
		path = path[4:]
	}

	switch path {
	case "distro":
		return c.Env.Distro, true
	case "distroVersion":
		return c.Env.DistroVersion, true
	case "arch":
		return c.Env.Arch, true
	case "desktop":
		return c.Env.Desktop, true
	case "isRoot":
		return c.Env.IsRoot, true
	case "hasPolkit":
		return c.Env.HasPolkit, true
	case "hasSudo":
		return c.Env.HasSudo, true
	case "diskFreeMB":
		return c.Env.DiskFreeMB, true
	case "installedVersion":
		return c.Env.InstalledVersion, true
	case "installDir":
		return c.Env.InstallDir, true
	}
	return nil, false
}

func writeLogEntry(w io.Writer, entry LogEntry) {
	level := "INFO"
	switch entry.Level {
	case LogWarn:
		level = "WARN"
	case LogError:
		level = "ERROR"
	}
	_, _ = fmt.Fprintf(w, "[%s] %s\n", level, entry.Message)
}

// getNestedValue retrieves a value from a nested map using dot notation.
func getNestedValue(m map[string]any, path string) (any, bool) {
	parts := strings.Split(path, ".")
	current := any(m)

	for _, part := range parts {
		switch v := current.(type) {
		case map[string]any:
			val, ok := v[part]
			if !ok {
				return nil, false
			}
			current = val
		default:
			return nil, false
		}
	}
	return current, true
}

// setNestedValue sets a value in a nested map using dot notation.
func setNestedValue(m map[string]any, path string, value any) {
	parts := strings.Split(path, ".")

	// Navigate/create nested maps
	current := m
	for i := 0; i < len(parts)-1; i++ {
		part := parts[i]
		if next, ok := current[part].(map[string]any); ok {
			current = next
		} else {
			// Create new nested map
			next := make(map[string]any)
			current[part] = next
			current = next
		}
	}

	// Set the final value
	current[parts[len(parts)-1]] = value
}

// currentTimeMillis returns current time in milliseconds (stub for now).
func currentTimeMillis() int64 {
	// Will be replaced with actual time implementation
	return 0
}
