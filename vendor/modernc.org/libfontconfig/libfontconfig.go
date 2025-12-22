// Copyright 2024 The libfontconfig-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generator.go

// Package libfontconfig is a ccgo/v4 version of libfontconfig.a, a library
// providing configuration, enumeration and substitution of fonts to other
// programs.
package libfontconfig // import "modernc.org/libfontconfig"

import (
	"sync/atomic"
	"unsafe"

	"modernc.org/libc"
)

func ___sync_bool_compare_and_swap(tls *libc.TLS, p, old, new uintptr) int32 {
	return libc.Bool32(atomic.CompareAndSwapUintptr((*uintptr)(unsafe.Pointer(p)), old, new))
}

func ___sync_fetch_and_add_impl(tls *libc.TLS, p uintptr, val int32) int32 {
	return atomic.AddInt32((*int32)(unsafe.Pointer(p)), val) - val
}
