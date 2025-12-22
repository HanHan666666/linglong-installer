// Copyright 2024 The libXdmcp-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generator.go

// Package libxdmcp is a ccgo/v4 version of libXdmcp.a, the X Display Manager
// Control Protocol library.
package libxdmcp // import "modernc.org/libxdmcp"

import "modernc.org/libc"

func _arc4random_buf(tls *libc.TLS, buf uintptr, n Tsize_t) {
	panic(todo(""))
}
