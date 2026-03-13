package ui

import (
	"testing"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

func TestIsUOSDeveloperModeRootError(t *testing.T) {
	ctx := core.NewInstallContext()
	ctx.Env.Distro = "uos"

	msg := "No root privileges. Please request root access in developer mode in Control Center."
	if !isUOSDeveloperModeRootError(ctx, msg) {
		t.Fatal("expected UOS developer mode privilege error to be detected")
	}
}

func TestIsUOSDeveloperModeRootErrorRejectsOtherDistros(t *testing.T) {
	ctx := core.NewInstallContext()
	ctx.Env.Distro = "deepin"

	msg := "No root privileges. Please request root access in developer mode in Control Center."
	if isUOSDeveloperModeRootError(ctx, msg) {
		t.Fatal("expected non-UOS distro not to trigger UOS-specific dialog")
	}
}

func TestIsUOSDeveloperModeRootErrorRejectsGenericExitStatus(t *testing.T) {
	ctx := core.NewInstallContext()
	ctx.Env.Distro = "uos"

	msg := "command failed: exit status 127"
	if isUOSDeveloperModeRootError(ctx, msg) {
		t.Fatal("expected generic exit status not to trigger UOS-specific dialog")
	}
}
