// Copyright 2024 The libtcl9_0-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generator.go

// Package libtcl9_0 is a ccgo/v4 version the Tool Command Language (Tcl).
package libtcl9_0 // import "modernc.org/libtcl9_0"

import (
	"modernc.org/libc"
)

const (
	Version = "tcl9.0.1"
)

func ___darwin_check_fd_set(t *libc.TLS, a int32, b uintptr) int32 {
	return 1
}

func _mkdtemp(tls *libc.TLS, template uintptr) (r uintptr) {
	panic(todo(""))
}

func ___ccgo_notrap(tls *libc.TLS) {
}
