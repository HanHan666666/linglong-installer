// Copyright 2023 The Knuth Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tds provides a minimal TeX Directory Structure.
package tds // import "modernc.org/knuth/internal/tds"

import (
	"embed"
)

// FS holds the contents of a minimal TeX Directory Structure.
//
//go:embed fonts/afm/public/amsfonts/cm/*.afm
//go:embed fonts/pk/ljfour/public/cm/dpi600/*.pk
//go:embed fonts/tfm/public/cm/*.tfm
//go:embed fonts/type1/public/amsfonts/cm/*pfb
//go:embed tex/generic/hyphen/hyphen.tex
//go:embed tex/latex/graphics
//go:embed tex/latex/xcolor
//go:embed tex/plain/base/plain.tex
var FS embed.FS
