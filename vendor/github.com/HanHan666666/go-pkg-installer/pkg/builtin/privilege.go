// Package builtin provides shared helpers for builtin tasks.
package builtin

import "github.com/HanHan666666/go-pkg-installer/pkg/core"

func ensurePrivilege(ctx *core.InstallContext, required bool) error {
	return core.EnsurePrivilege(ctx, required)
}
