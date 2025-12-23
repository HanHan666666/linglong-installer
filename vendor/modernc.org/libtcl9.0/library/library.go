// Copyright 2024 The libtcl9_0-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package library provides the Tcl standard library.
package library // import "modernc.org/libtcl9_0/library"

import (
	_ "embed"
)

// Zip contains the Tcl standard library archive.
//
//go:embed library.zip
var Zip string
