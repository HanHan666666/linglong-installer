// Copyright 2024 The libtk9_0-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package library provides the Tk standard library.
package library // import "modernc.org/libtk9_0/library"

import (
	_ "embed"
)

// Zip contains the Tk standard library archive.
//
//go:embed library.zip
var Zip string
