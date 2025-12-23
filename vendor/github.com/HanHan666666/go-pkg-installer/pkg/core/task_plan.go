// Package core provides task plan builders.
package core

import "fmt"

// BuildTaskPlan builds a plan from a flow configuration.
func BuildTaskPlan(flow *FlowConfig) *TaskPlan {
	if flow == nil {
		return nil
	}

	plan := &TaskPlan{Tasks: make([]TaskSummary, 0)}
	for _, step := range flow.Steps {
		for _, task := range step.Tasks {
			desc := task.ID
			if desc == "" {
				desc = fmt.Sprintf("%s task", task.Type)
			}
			plan.Tasks = append(plan.Tasks, TaskSummary{
				Type:         task.Type,
				Description:  desc,
				RequiresRoot: requiresPrivilege(task.Params),
			})
		}
	}
	return plan
}
