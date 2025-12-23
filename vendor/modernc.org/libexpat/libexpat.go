// Copyright 2023 The libexpat-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generator.go

// Package libexpat is a ccgo/v4 version of liexapat.a, a stream-oriented XML
// parser library.
package libexpat // import "modernc.org/libexpat"

import "modernc.org/libc"

func _arc4random_buf(tls *libc.TLS, buf uintptr, n Tsize_t) {
	panic(todo(""))
}
