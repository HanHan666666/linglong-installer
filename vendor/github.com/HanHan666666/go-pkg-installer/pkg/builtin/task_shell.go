// Package builtin provides the shell task implementation.
package builtin

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// ShellTask executes a shell command.
type ShellTask struct {
	core.BaseTask
	Command          string
	Args             []string
	WorkDir          string
	Env              map[string]string
	Timeout          time.Duration
	RequirePrivilege bool

	// Rollback command (optional)
	RollbackCmd string
}

// RegisterShellTask registers the shell task factory.
func RegisterShellTask() {
	core.Tasks.Register("shell", func(config map[string]any, ctx *core.InstallContext) (core.Task, error) {
		// Parse environment variables
		env := make(map[string]string)
		if envMap, ok := config["env"].(map[string]any); ok {
			for k, v := range envMap {
				if s, ok := v.(string); ok {
					env[k] = ctx.Render(s)
				}
			}
		}

		task := &ShellTask{
			BaseTask: core.BaseTask{
				TaskID:   getConfigString(config, "id"),
				TaskType: "shell",
				Config:   config,
			},
			Command:          ctx.Render(getConfigStringAny(config, "command", "script")),
			Args:             renderStringSlice(ctx, getConfigStringSlice(config, "args")),
			WorkDir:          ctx.Render(getConfigStringAny(config, "workDir", "workdir")),
			Env:              env,
			Timeout:          time.Duration(getConfigIntAny(config, 300, "timeoutSec", "timeout")) * time.Second,
			RollbackCmd:      ctx.Render(getConfigStringAny(config, "rollbackCommand", "rollback_command")),
			RequirePrivilege: getConfigBool(config, "requirePrivilege"),
		}

		if task.TaskID == "" {
			// Use first word of command as ID
			parts := strings.Fields(task.Command)
			if len(parts) > 0 {
				task.TaskID = fmt.Sprintf("shell-%s", parts[0])
			} else {
				task.TaskID = "shell-cmd"
			}
		}

		return task, nil
	})
}

// renderStringSlice applies template rendering to a slice of strings.
func renderStringSlice(ctx *core.InstallContext, slice []string) []string {
	result := make([]string, len(slice))
	for i, s := range slice {
		result[i] = ctx.Render(s)
	}
	return result
}

// Validate validates the shell task configuration.
func (t *ShellTask) Validate() error {
	if t.Command == "" {
		return errors.New("shell: command is required")
	}
	return nil
}

// Execute runs the shell command.
func (t *ShellTask) Execute(ctx *core.InstallContext, bus *core.EventBus) error {
	if err := ensurePrivilege(ctx, t.RequirePrivilege); err != nil {
		return err
	}

	ctx.AddLog(core.LogInfo, fmt.Sprintf("Executing: %s %s", t.Command, strings.Join(t.Args, " ")))

	// Default timeout if not set
	timeout := t.Timeout
	if timeout == 0 {
		timeout = 300 * time.Second
	}

	// Build command
	var cmd *exec.Cmd
	ctxTimeout, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if len(t.Args) > 0 {
		cmd = exec.CommandContext(ctxTimeout, t.Command, t.Args...)
	} else {
		// Use shell to execute the command
		cmd = exec.CommandContext(ctxTimeout, "sh", "-c", t.Command)
	}

	// Set working directory
	if t.WorkDir != "" {
		cmd.Dir = t.WorkDir
	}

	// Set environment
	if len(t.Env) > 0 {
		cmd.Env = os.Environ()
		for k, v := range t.Env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
		}
	}

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("shell: failed to capture stdout: %w", err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("shell: failed to capture stderr: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("shell: failed to start command: %w", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go streamOutput(ctx, stdoutPipe, core.LogInfo, "", &wg)
	go streamOutput(ctx, stderrPipe, core.LogError, "", &wg)

	err = cmd.Wait()
	wg.Wait()

	if ctxTimeout.Err() == context.DeadlineExceeded {
		ctx.AddLog(core.LogError, fmt.Sprintf("Command timed out after %v", timeout))
		return fmt.Errorf("command timed out after %v", timeout)
	}
	if err != nil {
		ctx.AddLog(core.LogError, fmt.Sprintf("Command failed: %v", err))
		return fmt.Errorf("command failed: %w", err)
	}

	return nil
}

// CanRollback returns true if a rollback command is configured.
func (t *ShellTask) CanRollback() bool {
	return t.RollbackCmd != ""
}

// Rollback executes the rollback command.
func (t *ShellTask) Rollback(ctx *core.InstallContext, bus *core.EventBus) error {
	if t.RollbackCmd == "" {
		return nil
	}

	ctx.AddLog(core.LogInfo, fmt.Sprintf("Rolling back: %s", t.RollbackCmd))

	cmd := exec.Command("sh", "-c", t.RollbackCmd)
	if t.WorkDir != "" {
		cmd.Dir = t.WorkDir
	}

	if len(t.Env) > 0 {
		cmd.Env = os.Environ()
		for k, v := range t.Env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
		}
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		ctx.AddLog(core.LogWarn, fmt.Sprintf("Rollback command failed: %v\nstderr: %s", err, stderr.String()))
		return fmt.Errorf("rollback command failed: %w", err)
	}

	return nil
}

func streamOutput(ctx *core.InstallContext, reader io.Reader, level core.LogLevel, label string, wg *sync.WaitGroup) {
	defer wg.Done()

	scanner := bufio.NewScanner(reader)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\r")
		if line == "" {
			continue
		}
		if label != "" {
			ctx.AddLog(level, fmt.Sprintf("%s: %s", label, line))
		} else {
			ctx.AddLog(level, line)
		}
	}

	if err := scanner.Err(); err != nil {
		ctx.AddLog(core.LogWarn, fmt.Sprintf("stream read error (%s): %v", label, err))
	}
}
