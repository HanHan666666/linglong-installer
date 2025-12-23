// Package builtin provides the permission task implementation.
package builtin

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

// PermissionTask updates permissions and ownership.
type PermissionTask struct {
	core.BaseTask
	Path             string
	Mode             os.FileMode
	Owner            string
	Group            string
	Recursive        bool
	RequirePrivilege bool
}

// RegisterPermissionTask registers the permission task factory.
func RegisterPermissionTask() {
	core.Tasks.Register("permission", func(config map[string]any, ctx *core.InstallContext) (core.Task, error) {
		mode := os.FileMode(0)
		if modeVal := parseFileMode(config["mode"]); modeVal != 0 {
			mode = os.FileMode(modeVal)
		}

		task := &PermissionTask{
			BaseTask: core.BaseTask{
				TaskID:   getConfigString(config, "id"),
				TaskType: "permission",
				Config:   config,
			},
			Path:             ctx.Render(getConfigString(config, "path")),
			Mode:             mode,
			Owner:            ctx.Render(getConfigString(config, "owner")),
			Group:            ctx.Render(getConfigString(config, "group")),
			Recursive:        getConfigBool(config, "recursive"),
			RequirePrivilege: getConfigBoolDefault(config, "requirePrivilege", true),
		}

		if task.TaskID == "" {
			task.TaskID = fmt.Sprintf("permission-%s", filepath.Base(task.Path))
		}

		return task, nil
	})
}

// Validate validates the permission task configuration.
func (t *PermissionTask) Validate() error {
	if t.Path == "" {
		return errors.New("permission: path is required")
	}
	if t.Mode == 0 && t.Owner == "" && t.Group == "" {
		return errors.New("permission: mode or owner/group is required")
	}
	return nil
}

// Execute updates file permissions/ownership.
func (t *PermissionTask) Execute(ctx *core.InstallContext, bus *core.EventBus) error {
	if err := ensurePrivilege(ctx, t.RequirePrivilege); err != nil {
		return err
	}

	ctx.AddLog(core.LogInfo, fmt.Sprintf("Updating permissions: %s", t.Path))

	if t.Recursive {
		return filepath.Walk(t.Path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			return t.apply(path)
		})
	}

	return t.apply(t.Path)
}

func (t *PermissionTask) apply(path string) error {
	if t.Mode != 0 {
		if err := os.Chmod(path, t.Mode); err != nil {
			return fmt.Errorf("chmod failed for %s: %w", path, err)
		}
	}

	if t.Owner != "" || t.Group != "" {
		uid, gid, err := resolveOwnerGroup(t.Owner, t.Group)
		if err != nil {
			return err
		}
		if err := os.Chown(path, uid, gid); err != nil {
			return fmt.Errorf("chown failed for %s: %w", path, err)
		}
	}

	return nil
}

// CanRollback returns false by default.
func (t *PermissionTask) CanRollback() bool {
	return false
}

// Rollback does nothing for permission operations.
func (t *PermissionTask) Rollback(ctx *core.InstallContext, bus *core.EventBus) error {
	return nil
}

func resolveOwnerGroup(owner, group string) (int, int, error) {
	uid := -1
	gid := -1

	if owner != "" {
		if parsed, err := strconv.Atoi(owner); err == nil {
			uid = parsed
		} else {
			u, err := user.Lookup(owner)
			if err != nil {
				return -1, -1, fmt.Errorf("unknown owner: %s", owner)
			}
			parsed, err := strconv.Atoi(u.Uid)
			if err != nil {
				return -1, -1, fmt.Errorf("invalid uid: %s", u.Uid)
			}
			uid = parsed
		}
	}

	if group != "" {
		if parsed, err := strconv.Atoi(group); err == nil {
			gid = parsed
		} else {
			g, err := user.LookupGroup(group)
			if err != nil {
				return -1, -1, fmt.Errorf("unknown group: %s", group)
			}
			parsed, err := strconv.Atoi(g.Gid)
			if err != nil {
				return -1, -1, fmt.Errorf("invalid gid: %s", g.Gid)
			}
			gid = parsed
		}
	}

	if uid == -1 && gid == -1 {
		return -1, -1, errors.New("owner/group not resolved")
	}

	if uid == -1 {
		uid = os.Geteuid()
	}
	if gid == -1 {
		gid = os.Getegid()
	}

	return uid, gid, nil
}
