// Package builtin provides the systemd_service task implementation.
package builtin

import (
	"errors"
	"fmt"
	"strings"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// SystemdServiceTask manages systemd services.
type SystemdServiceTask struct {
	core.BaseTask
	Name             string
	Action           string
	RequirePrivilege bool
	UserScope        bool
}

// RegisterSystemdServiceTask registers the systemd_service task factory.
func RegisterSystemdServiceTask() {
	core.Tasks.Register("systemdService", func(config map[string]any, ctx *core.InstallContext) (core.Task, error) {
		requirePrivilege := getConfigBoolDefault(config, "requirePrivilege", true)
		task := &SystemdServiceTask{
			BaseTask: core.BaseTask{
				TaskID:   getConfigString(config, "id"),
				TaskType: "systemdService",
				Config:   config,
			},
			Name:             ctx.Render(getConfigString(config, "name")),
			Action:           strings.ToLower(getConfigString(config, "action")),
			RequirePrivilege: requirePrivilege,
			UserScope:        getConfigBool(config, "user"),
		}

		if task.TaskID == "" {
			task.TaskID = fmt.Sprintf("systemd-%s", task.Name)
		}

		if task.UserScope {
			task.RequirePrivilege = false
		}

		return task, nil
	})
}

// Validate validates the systemd_service task configuration.
func (t *SystemdServiceTask) Validate() error {
	if t.Name == "" {
		return errors.New("systemdService: name is required")
	}
	if t.Action == "" {
		return errors.New("systemdService: action is required")
	}
	return nil
}

// Execute performs the systemd action.
func (t *SystemdServiceTask) Execute(ctx *core.InstallContext, bus *core.EventBus) error {
	if err := ensurePrivilege(ctx, t.RequirePrivilege); err != nil {
		return err
	}

	unit := normalizeUnitName(t.Name)
	args := buildSystemctlArgs(t.Action, unit)
	if len(args) == 0 {
		return fmt.Errorf("systemdService: unsupported action %q", t.Action)
	}

	if t.UserScope {
		args = append([]string{"--user"}, args...)
	}

	ctx.AddLog(core.LogInfo, fmt.Sprintf("systemd %s %s", t.Action, unit))
	cmd := execCommand("systemctl", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("systemctl failed: %w: %s", err, strings.TrimSpace(string(output)))
	}
	return nil
}

// CanRollback returns false by default.
func (t *SystemdServiceTask) CanRollback() bool {
	return false
}

// Rollback does nothing for systemd operations.
func (t *SystemdServiceTask) Rollback(ctx *core.InstallContext, bus *core.EventBus) error {
	return nil
}

func normalizeUnitName(name string) string {
	if name == "" {
		return name
	}
	if strings.Contains(name, ".") {
		return name
	}
	return name + ".service"
}

func buildSystemctlArgs(action, unit string) []string {
	switch action {
	case "enable_and_start":
		return []string{"enable", "--now", unit}
	case "stop_and_disable":
		return []string{"disable", "--now", unit}
	case "start", "stop", "restart", "enable", "disable":
		return []string{action, unit}
	default:
		return nil
	}
}
