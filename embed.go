// Package main provides embedded resources for the installer
package main

import "embed"

//go:embed installer.yaml
var embeddedConfig []byte

//go:embed assets/logo.png
var embeddedLogo []byte

//go:embed scripts/common.sh scripts/distros/*.sh scripts/user/*.sh
var embeddedScripts embed.FS
