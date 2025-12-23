// Copyright 2024 The libXdmcp-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generator.go

// Package libx11 is a ccgo/v4 version of libX11.a, a library implementing the
// Core X11 client protocol.
package libx11 // import "modernc.org/libx11"

import (
	"os"
	"unsafe"

	"modernc.org/libc"
)

// see https://gitlab.com/cznic/libx11/-/issues/2
// This is not ideal, but will have to do for now.
func init() {
	const localeDir = "XLOCALEDIR"
	if os.Getenv(localeDir) != "" {
		// Respect user setting, if any.
		return
	}

	possibleX11LocaleDirs := []string{
		"/usr/share/X11/locale",
		"/usr/local/share/X11/locale",
	}

	for _, dir := range possibleX11LocaleDirs {
		if fi, err := os.Stat(dir); err == nil && fi.IsDir() {
			if err = os.Setenv(localeDir, dir); err == nil {
				return
			}

			panic(err)
		}
	}
}

// int __darwin_check_fd_set_overflow(int, void *, int);
var __darwin_check_fd_set_overflow func(*libc.TLS, int32, uintptr, int32) int32

var utf8 = [...]byte{'U', 'T', 'F', '-', '8', 0}

func __XkbGetCharset(tls *libc.TLS) (r uintptr) {
	// For Go we support utf-8 encoding only.
	return uintptr(unsafe.Pointer(&utf8))
}

func ____mb_cur_max(tls *libc.TLS) int32 {
	return 4
}

type socklen_t = uint32

func ___inet_ntop(tls *libc.TLS, af int32, a0 uintptr, s uintptr, l socklen_t) uintptr {
	return libc.Xinet_ntop(tls, af, a0, s, l)
}

func ___inet_pton(tls *libc.TLS, af int32, s uintptr, a0 uintptr) int32 {
	return libc.Xinet_pton(tls, af, s, a0)
}
