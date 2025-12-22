// Copyright 2023 The libXau-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generator.go

// Package libxau is a ccgo/v4 version of libXau.a, a library implementing the
// X11 Authorization Protocol.
package libxau // import "modernc.org/libxau"

import (
	"modernc.org/libc"
	"modernc.org/libc/sys/types"
)

func _explicit_bzero(tls *libc.TLS, d uintptr, n types.Size_t) {
	libc.Xmemset(tls, d, 0, n)
}
