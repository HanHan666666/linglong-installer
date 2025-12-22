// Package builtin provides built-in task implementations for the installer.
package builtin

// RegisterAll registers all built-in tasks with the global registry.
func RegisterAll() {
	RegisterDownloadTask()
	RegisterUnpackTask()
	RegisterCopyTask()
	RegisterSymlinkTask()
	RegisterShellTask()
	RegisterWriteConfigTask()
	RegisterDesktopEntryTask()
	RegisterRemovePathTask()
	RegisterSystemdServiceTask()
	RegisterDbusServiceTask()
	RegisterPermissionTask()
	RegisterNetScriptTask()
	RegisterRemoveDesktopEntryTask()
	RegisterRollbackTask()
}
