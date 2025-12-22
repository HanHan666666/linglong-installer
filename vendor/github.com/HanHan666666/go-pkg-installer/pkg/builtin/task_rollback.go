// Package builtin provides the rollback task implementation.
package builtin

import (
	"errors"
	"fmt"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// RollbackTask holds tasks to run during rollback.
type RollbackTask struct {
	core.BaseTask
	TasksCfg []map[string]any
}

// RegisterRollbackTask registers the rollback task factory.
func RegisterRollbackTask() {
	core.Tasks.Register("rollback", func(config map[string]any, ctx *core.InstallContext) (core.Task, error) {
		task := &RollbackTask{
			BaseTask: core.BaseTask{
				TaskID:   getConfigString(config, "id"),
				TaskType: "rollback",
				Config:   config,
			},
			TasksCfg: parseTaskConfigs(config["tasks"]),
		}

		if task.TaskID == "" {
			task.TaskID = "rollback"
		}

		return task, nil
	})
}

// Validate validates the rollback task configuration.
func (t *RollbackTask) Validate() error {
	if len(t.TasksCfg) == 0 {
		return errors.New("rollback: tasks are required")
	}
	return nil
}

// Execute is a no-op for rollback tasks.
func (t *RollbackTask) Execute(ctx *core.InstallContext, bus *core.EventBus) error {
	ctx.AddLog(core.LogInfo, "Rollback task registered")
	return nil
}

// CanRollback returns true when tasks are configured.
func (t *RollbackTask) CanRollback() bool {
	return len(t.TasksCfg) > 0
}

// Rollback executes the configured rollback tasks.
func (t *RollbackTask) Rollback(ctx *core.InstallContext, bus *core.EventBus) error {
	for _, cfg := range t.TasksCfg {
		typeName, ok := cfg["type"].(string)
		if !ok || typeName == "" {
			return errors.New("rollback: task type is required")
		}

		factory, ok := core.Tasks.Get(typeName)
		if !ok && core.IsGoExtension(typeName) {
			factory, ok = core.Tasks.Get(core.StripGoPrefix(typeName))
		}
		if !ok {
			return fmt.Errorf("rollback: unknown task type %s", typeName)
		}

		task, err := factory(cfg, ctx)
		if err != nil {
			return fmt.Errorf("rollback: failed to create task %s: %w", typeName, err)
		}

		if err := task.Validate(); err != nil {
			return fmt.Errorf("rollback: validation failed for %s: %w", typeName, err)
		}

		if err := task.Execute(ctx, bus); err != nil {
			return fmt.Errorf("rollback: task %s failed: %w", typeName, err)
		}
	}

	return nil
}

func parseTaskConfigs(value any) []map[string]any {
	list, ok := value.([]any)
	if !ok {
		return nil
	}

	result := make([]map[string]any, 0, len(list))
	for _, item := range list {
		if cfg, ok := item.(map[string]any); ok {
			result = append(result, cfg)
		}
	}
	return result
}
