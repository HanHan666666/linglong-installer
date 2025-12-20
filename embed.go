// Package main provides embedded resources for the installer
package main

import (
	_ "embed"
)

//go:embed installer.yaml
var embeddedConfig []byte

//go:embed assets/logo.png
var embeddedLogo []byte
