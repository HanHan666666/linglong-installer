// Package core provides privilege helpers.
package core

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	PrivilegeNone   = "none"
	PrivilegeSudo   = "sudo"
	PrivilegePkexec = "pkexec"
)

// GetPrivilegeStrategy resolves the configured privilege strategy.
func GetPrivilegeStrategy(ctx *InstallContext) string {
	keys := []string{
		"privilege.strategy",
		"meta.privilegeStrategy",
		"meta.privilege_strategy",
	}

	for _, key := range keys {
		if val, ok := ctx.Get(key); ok && val != nil {
			return normalizeStrategy(fmt.Sprintf("%v", val))
		}
	}

	return PrivilegeNone
}

func normalizeStrategy(strategy string) string {
	switch strings.ToLower(strings.TrimSpace(strategy)) {
	case PrivilegeSudo:
		return PrivilegeSudo
	case PrivilegePkexec:
		return PrivilegePkexec
	default:
		return PrivilegeNone
	}
}

// EnsurePrivilege checks privilege requirement against environment state.
func EnsurePrivilege(ctx *InstallContext, required bool) error {
	if !required {
		return nil
	}
	if ctx.Env.IsRoot {
		return nil
	}

	strategy := GetPrivilegeStrategy(ctx)
	msg := fmt.Sprintf("administrator privileges required (strategy: %s)", strategy)
	ctx.AddLog(LogError, msg)
	return errors.New(msg)
}

// NeedsPrivilege returns true when any task config marks requirePrivilege.
func NeedsPrivilege(cfg *Config, flowID string) bool {
	if cfg == nil {
		return false
	}

	if cfg.Flows == nil {
		return false
	}

	flow := cfg.Flows[flowID]
	if flow == nil {
		return false
	}

	for _, step := range flow.Steps {
		for _, task := range step.Tasks {
			if taskRequiresPrivilege(task) {
				return true
			}
		}
	}

	return false
}

func requiresPrivilege(params map[string]any) bool {
	val, ok := params["requirePrivilege"]
	if !ok {
		return false
	}
	switch v := val.(type) {
	case bool:
		return v
	case string:
		return strings.EqualFold(v, "true") || v == "1"
	default:
		return false
	}
}

func taskRequiresPrivilege(task TaskConfig) bool {
	if requiresPrivilege(task.Params) {
		return true
	}

	switch StripGoPrefix(task.Type) {
	case "systemdService", "dbusService", "permission":
		if StripGoPrefix(task.Type) != "permission" && isUserScope(task.Params) {
			return false
		}
		return true
	default:
		return false
	}
}

func isUserScope(params map[string]any) bool {
	val, ok := params["user"]
	if !ok {
		return false
	}
	switch v := val.(type) {
	case bool:
		return v
	case string:
		return strings.EqualFold(v, "true") || v == "1"
	default:
		return false
	}
}

// Elevated returns true if the process is already elevated.
func Elevated() bool {
	return os.Getenv("GPKI_ELEVATED") == "1"
}
