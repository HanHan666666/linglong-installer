// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "libtcl9.0.so"
	tkBin  = "libtcl9tk9.0.so"
)

var moreDLLs = []*dllInfo{
	{"libtcl9tkimg2.0.1.so", "Tkimg_Init"},
	{"libtcl9jpegtcl9.6.0.so", "Jpegtcl_Init"},
	{"libtcl9tkimgjpeg2.0.1.so", "Tkimgjpeg_Init"},
	{"libtcl9tkimgbmp2.0.1.so", "Tkimgbmp_Init"},
	{"libtcl9tkimgico2.0.1.so", "Tkimgico_Init"},
	{"libtcl9tkimgpcx2.0.1.so", "Tkimgpcx_Init"},
	{"libtcl9tkimgxpm2.0.1.so", "Tkimgxpm_Init"},
	{"libtcl9zlibtcl1.3.1.so", "Zlibtcl_Init"},
	{"libtcl9pngtcl1.6.44.so", "Pngtcl_Init"},
	{"libtcl9tkimgpng2.0.1.so", "Tkimgpng_Init"},
	{"libtcl9tkimgppm2.0.1.so", "Tkimgppm_Init"},
	{"libtcl9tkimgtga2.0.1.so", "Tkimgtga_Init"},
	{"libtcl9tifftcl4.7.0.so", "Tifftcl_Init"},
	{"libtcl9tkimgtiff2.0.1.so", "Tkimgtiff_Init"},
	{"libtcl9tkimgxbm2.0.1.so", "Tkimgxbm_Init"},
}

//go:embed embed/linux/amd64/lib.zip
var libZip []byte

// Keep last for internal/shasig.go to update.
var shasig = map[string]string{
	"libtcl9.0.so":               "e9a16bfe701f2e2ba78c6e6a76a3da1d2dbc532b7d047a2aaa65c729c1e13d47",
	"libtcl9jpegtcl9.6.0.so":     "63f76dd8c85d7d5df1467daeb8a5ed9a194453952de2770822b08d1c1fcbb183",
	"libtcl9pngtcl1.6.44.so":     "f103ac1c4ef270ec3c87bbd7b6f0f9f20f2fa1bd31355ee4ad3f8435660f371d",
	"libtcl9tifftcl4.7.0.so":     "d6118b52179266a8d754058ef2e6325d1e8533d1f8b75c5b00b557b0285be409",
	"libtcl9tk9.0.so":            "4d487c7929f934240246542d1d512d6bb5b92e1b4bdb66992c464803d673cac9",
	"libtcl9tkimg2.0.1.so":       "e169990a08bec5f77ae2d179930014d76e0d4665e40235d4a5ea2e0242b2e4de",
	"libtcl9tkimgbmp2.0.1.so":    "b0cbcb898628d6e484fc8703dcb62970084d80998bcdda8cac518c1f33dae7a9",
	"libtcl9tkimgdted2.0.1.so":   "c20fd2609540c6b653d72278be918f4c413c5823f8665f39ad3ec74dda98af3b",
	"libtcl9tkimgflir2.0.1.so":   "5cc806c3b60ca62af263cc531e7fec255041a25fabeadb00bb6425702fdee9f9",
	"libtcl9tkimggif2.0.1.so":    "1f8c067a4e4c678e7c1f6eb51f78e747e855f6d267e0c716de3e9653489e7781",
	"libtcl9tkimgico2.0.1.so":    "be1dfd3944bda21da8e7aef4753feb7caa8739f4e6e57340a9c1997a4f75626f",
	"libtcl9tkimgjpeg2.0.1.so":   "1374c2f1a7764add255f53b4467cf06724fffe49c17b10b9506cc20b3fe6309b",
	"libtcl9tkimgpcx2.0.1.so":    "9c123162baffd848c3ca96ca00eaf7aea867d86dae40553093f891ac281405ba",
	"libtcl9tkimgpixmap2.0.1.so": "c90853bd237b41738a57f679498fb3915a75cd8fa6d7adf77cbd6701f898f624",
	"libtcl9tkimgpng2.0.1.so":    "0c951db84dc5c87740c0b8c761a0078bc17c975ba5c6533c2f9500d04f4f9d97",
	"libtcl9tkimgppm2.0.1.so":    "54273bc22d9c62e8574a36103a4f9a602c1c0378ce32aa4e50716ef82d60b34d",
	"libtcl9tkimgps2.0.1.so":     "f3b8a44bb5b04e49aa13f6e11970609b31fb030297b927d0a2247b8d193de4ae",
	"libtcl9tkimgraw2.0.1.so":    "935539955741ef1427fbbed5ec8b341c87777bb936174bcded2be7aa31211062",
	"libtcl9tkimgsgi2.0.1.so":    "9c32db1fbe04c8c94cd9092169f7f91adec8f4f91acb0848234cd6d8b78eefdf",
	"libtcl9tkimgsun2.0.1.so":    "d192cec47d2fcd132739169513938f6a5755c05843410fbda6500e943d4de469",
	"libtcl9tkimgtga2.0.1.so":    "f5ea5f26f553d824ee7f70c79c2d647776caedb7eb2ba39d441be0f9a629bb19",
	"libtcl9tkimgtiff2.0.1.so":   "09a4582625cb0adadb40cc080636ba8b34926781d8de8c6d62eebac2dd41b4f1",
	"libtcl9tkimgwindow2.0.1.so": "a7389d273a93f14c02551523cf46f597246634d7496360171d7cb46568b0ab95",
	"libtcl9tkimgxbm2.0.1.so":    "af5f1adc39684dcf0fa34830b5fbe9243be1312c1bbcd9a01bb17a2539e0bfde",
	"libtcl9tkimgxpm2.0.1.so":    "3500fc2e7c95eba10af0eb72c4a3f380ab7214a2e8aa5dace329140809e70f3f",
	"libtcl9zlibtcl1.3.1.so":     "0f76a0fabdbcefed4809f7e880813036a82900d356eb252ca204e0e9bd347d32",
	"libtk9.0.1.zip":             "df0ac2f2c47c77d5247919b06b93f9efd60a94ce3aef94db239ce0201ae4abc1",
}
