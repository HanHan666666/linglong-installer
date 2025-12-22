// Package builtin provides the dbus_service task implementation.
package builtin

import (
	"errors"
	"fmt"
	"strings"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// DbusServiceTask manages a D-Bus activated service via systemctl.
type DbusServiceTask struct {
	core.BaseTask
	Name             string
	Action           string
	RequirePrivilege bool
	UserScope        bool
}

// RegisterDbusServiceTask registers the dbus_service task factory.
func RegisterDbusServiceTask() {
	core.Tasks.Register("dbusService", func(config map[string]any, ctx *core.InstallContext) (core.Task, error) {
		requirePrivilege := getConfigBoolDefault(config, "requirePrivilege", true)
		task := &DbusServiceTask{
			BaseTask: core.BaseTask{
				TaskID:   getConfigString(config, "id"),
				TaskType: "dbusService",
				Config:   config,
			},
			Name:             ctx.Render(getConfigString(config, "name")),
			Action:           strings.ToLower(getConfigString(config, "action")),
			RequirePrivilege: requirePrivilege,
			UserScope:        getConfigBool(config, "user"),
		}

		if task.TaskID == "" {
			task.TaskID = fmt.Sprintf("dbus-%s", task.Name)
		}

		if task.UserScope {
			task.RequirePrivilege = false
		}

		return task, nil
	})
}

// Validate validates the dbus_service task configuration.
func (t *DbusServiceTask) Validate() error {
	if t.Name == "" {
		return errors.New("dbusService: name is required")
	}
	if t.Action == "" {
		return errors.New("dbusService: action is required")
	}
	return nil
}

// Execute performs the dbus action through systemctl.
func (t *DbusServiceTask) Execute(ctx *core.InstallContext, bus *core.EventBus) error {
	if err := ensurePrivilege(ctx, t.RequirePrivilege); err != nil {
		return err
	}

	unit := normalizeUnitName(t.Name)
	args := buildSystemctlArgs(t.Action, unit)
	if len(args) == 0 {
		return fmt.Errorf("dbusService: unsupported action %q", t.Action)
	}

	if t.UserScope {
		args = append([]string{"--user"}, args...)
	}

	ctx.AddLog(core.LogInfo, fmt.Sprintf("dbus %s %s", t.Action, unit))
	cmd := execCommand("systemctl", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("systemctl failed: %w: %s", err, strings.TrimSpace(string(output)))
	}
	return nil
}

// CanRollback returns false by default.
func (t *DbusServiceTask) CanRollback() bool {
	return false
}

// Rollback does nothing for dbus operations.
func (t *DbusServiceTask) Rollback(ctx *core.InstallContext, bus *core.EventBus) error {
	return nil
}
