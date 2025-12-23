// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "libtcl90.so.1.0"
	tkBin  = "libtcl9tk90.so.1.0"
)

var moreDLLs = []*dllInfo{
	{"libtcl9tkimg201.so", "Tkimg_Init"},
	{"libtcl9jpegtcl960.so", "Jpegtcl_Init"},
	{"libtcl9tkimgjpeg201.so", "Tkimgjpeg_Init"},
	{"libtcl9tkimgbmp201.so", "Tkimgbmp_Init"},
	{"libtcl9tkimgico201.so", "Tkimgico_Init"},
	{"libtcl9tkimgpcx201.so", "Tkimgpcx_Init"},
	{"libtcl9tkimgxpm201.so", "Tkimgxpm_Init"},
	{"libtcl9zlibtcl131.so", "Zlibtcl_Init"},
	{"libtcl9pngtcl1644.so", "Pngtcl_Init"},
	{"libtcl9tkimgpng201.so", "Tkimgpng_Init"},
	{"libtcl9tkimgppm201.so", "Tkimgppm_Init"},
	{"libtcl9tkimgtga201.so", "Tkimgtga_Init"},
	{"libtcl9tifftcl470.so", "Tifftcl_Init"},
	{"libtcl9tkimgtiff201.so", "Tkimgtiff_Init"},
	{"libtcl9tkimgxbm201.so", "Tkimgxbm_Init"},
}

//go:embed embed/openbsd/amd64/lib.zip
var libZip []byte

// Keep last for internal/shasig.go to update.
var shasig = map[string]string{
	"libtcl90.so.1.0":          "861ca32b8281500f02488f7fc13b9e2db8e9baa3f8a3aa17b8ce77b3a4240770",
	"libtcl9jpegtcl960.so":     "13b70b506b3bf4e44a3d8f6cc029096f7fbe799a05c63fa75fba67094743f069",
	"libtcl9pngtcl1644.so":     "307e6ff4a9fd04c8776100a4e9a613d9bc0c5f775700dd4d9b0153ea2f22f0d5",
	"libtcl9tifftcl470.so":     "181bbcfbb290b66a57b7fdf5d282089d9723e0166014d2597b103c6e618a28b8",
	"libtcl9tk90.so.1.0":       "4ad27f15f533b944c3ad011317e26609072f5ac6f311568bc202694386344100",
	"libtcl9tkimg201.so":       "aeb9c3d501e490e8324db149e571c2e37711633b7160393e62b9c0087b05c15c",
	"libtcl9tkimgbmp201.so":    "cc6b460109adaf559aeb5d606edf7e7f84c82eee7998c0ed4e7c37c8a3172dea",
	"libtcl9tkimgdted201.so":   "8e47f291d1c3c09e3de9a85c69a997fc898e972579ee12ecd2e0eeb944b1b3e3",
	"libtcl9tkimgflir201.so":   "54ff7b33cbe797aa434ff5776398e7c2316d50e574916c4ec36b73267b516dd6",
	"libtcl9tkimggif201.so":    "22c5b053bd2811307e0fe8780b59e9b5a618ab1454152a3be241a27aa9b1cd3a",
	"libtcl9tkimgico201.so":    "e5d825625a9f2e656a81fa4c3741c44c3c63f6096095c1d57d42876e2c0e866f",
	"libtcl9tkimgjpeg201.so":   "ebc34987d7c689f307229b4cc209a008acb3203665e689acfba025429364019c",
	"libtcl9tkimgpcx201.so":    "41dc1304fc75cc66180820cce5a8b650e29bce782cad0d1cb72cd351596a7f77",
	"libtcl9tkimgpixmap201.so": "ad038ae2f0002634a9c9fa5b03ec533b8ea7173d306c504fac128bbef982fb83",
	"libtcl9tkimgpng201.so":    "22bbae417268c6e1017fb7d5c56631f6d8f900ea58d54d42f2a00181b98966de",
	"libtcl9tkimgppm201.so":    "81b3d7d79f56fc414c182c62b6ad0a8f053a1f21363fbefc14760a59a040c355",
	"libtcl9tkimgps201.so":     "39ebfe798df4b93390f0b78698892b0e69f7c17216352100c8b8be270184d6fc",
	"libtcl9tkimgraw201.so":    "08d79c4901801605c09a701dcbd6598954d36d0774282aabd03b14191d9e5f6a",
	"libtcl9tkimgsgi201.so":    "16a30334f322f90cdea3cba337bd17e129f3835ed6642c3e1fc26f26e229dad2",
	"libtcl9tkimgsun201.so":    "0576680ff1fce6ebcb10e3b5f270a8f3547c3cc243f9f3f36eb13dc2c75e4599",
	"libtcl9tkimgtga201.so":    "3dd02f92b63e812704db74dc0de12f0e8473b860fd235423b04621e790cdd4c3",
	"libtcl9tkimgtiff201.so":   "2b32ad4bd802973604e385dd1a91e19276e0c1a01969abf27de021d8bd42c1db",
	"libtcl9tkimgwindow201.so": "5345194f2bbe8da3c7469bd3423531cce75307d13f86b6c9b8fb9f85e13a8acd",
	"libtcl9tkimgxbm201.so":    "fcf79977b68a8f7d2f82c2b65a061fb7d26b96aaaad5c9db663aec66b29323c5",
	"libtcl9tkimgxpm201.so":    "a66e8283c3e633d0efa33c23519da8bfcb3798623f8531739cccfeb8c3582e0f",
	"libtcl9zlibtcl131.so":     "34b4e1c55887b0b36428a7958316681fdff23ba2144b5af86ef9c05f13668032",
	"libtk9.0.1.zip":           "512e932b82594d805aff1377d57fb6b021b61187e2eee251825eb35d2950fcf3",
}
