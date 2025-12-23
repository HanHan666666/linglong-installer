// Copyright 2023 The libbsd-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generator.go

// Package libsd is a ccgo/v4 version of libbsd.a: Utility functions from BSD systems.
package libbsd // import "modernc.org/libbsd"

import (
	"modernc.org/libc"
)

func __rs_forkhandler(tls *libc.TLS) {
	// nop
}

func __rs_forkdetect(tls *libc.TLS) {
	// nop
}

func _warnx(tls *libc.TLS, format uintptr, va uintptr) {
	panic(todo(""))
	// Xwarnc(tls, -1, format, va)
}

func _dl_iterate_phdr(...any) {
	panic(todo(""))
}
func _statfs(...any) int32 {
	panic(todo(""))
}

func _strtonum(tls *libc.TLS, nptr uintptr, minval int64, maxval int64, errstr uintptr) (r int64) {
	panic(todo(""))
}

// int dirfd(DIR *dirp);
func _dirfd(tls *libc.TLS, dir uintptr) (r int32) {
	panic(todo(""))
}

// int openat(int fd, const char *path, int oflag, ...);
func _openat(tls *libc.TLS, fd int32, path uintptr, oflag int32, va uintptr) (r int32) {
	panic(todo(""))
}

// int fstatat(int fd, const char *path, struct stat *buf, int flag);
func _fstatat(tls *libc.TLS, fd int32, path, buf uintptr, flag int32) (r int32) {
	panic(todo(""))
}

// int asprintf(char **ret, const char *format, ...);
func _asprintf(tls *libc.TLS, ret, format uintptr, va uintptr) (r int32) {
	panic(todo(""))
}

// void setprogname(const char *progname);
func _setprogname(tls *libc.TLS, progname uintptr) {
	panic(todo(""))
}

// int getpagesize(void);
func _getpagesize(tls *libc.TLS) int32 {
	panic(todo(""))
}
