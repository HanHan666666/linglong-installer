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

//go:embed embed/linux/arm64/lib.zip
var libZip []byte

// Keep last for internal/shasig.go to update.
var shasig = map[string]string{
	"libtcl9.0.so":               "3999da239341cfabb01d3c53e532410d9ff88e0bd8e2075430286c1f7c185a7d",
	"libtcl9jpegtcl9.6.0.so":     "d750cda246c00b1da6cfd3c9440dd3a599f26357e60da92464699e735b675368",
	"libtcl9pngtcl1.6.44.so":     "d83489190c1847ab1f7ea5a0d1db060aab76402352dec70e656a78990992331c",
	"libtcl9tifftcl4.7.0.so":     "ff9670a330d4901bc22c6bdd6ae416f2440819a55a706eb018eb3419436c76b6",
	"libtcl9tk9.0.so":            "75106b0a1cfb49089483cb2c07eebb811e68a7561198bed01ea5ea49c0a07e2b",
	"libtcl9tkimg2.0.1.so":       "8767b782c6ae964a3a41c50d507f49012f30ad974e8e8dc15cdd17933f168a5a",
	"libtcl9tkimgbmp2.0.1.so":    "ac5012b089986c9189733f5da327e8d49d3ada954ee24a5e2b8a41f2560c2566",
	"libtcl9tkimgdted2.0.1.so":   "0c52f670784cda34cb006696bd9201daab85001a59137c6e865e29bf0d2543fc",
	"libtcl9tkimgflir2.0.1.so":   "5052aa5046507a7cfc97030c5a82621de14266fd80b3770a256650c89554b1f2",
	"libtcl9tkimggif2.0.1.so":    "eba6b783a838acebb904b15216b9e731dfddc1c374432fda0d0d79dafcb1a58d",
	"libtcl9tkimgico2.0.1.so":    "3b4e5ec3ee8d6630f267e4af6a110d845f759c21d97238c284b48fe0cf72048b",
	"libtcl9tkimgjpeg2.0.1.so":   "18ce28def6390bb95d1a753aa95984e5b1dc19ca90a356d6ebf9d69a33cb6efd",
	"libtcl9tkimgpcx2.0.1.so":    "194c6f6863ce09de3abaa0e6522c9a42b4bffe231f26a1e871b54177887acee9",
	"libtcl9tkimgpixmap2.0.1.so": "e794ac89be42892fd9e37f9781a06c6598a078899962ceda316f7c0a6bf0c625",
	"libtcl9tkimgpng2.0.1.so":    "8d9cf6fef2ff76d505f42ce1ab4dbedf871b1524b234e6bf356a02168997cc33",
	"libtcl9tkimgppm2.0.1.so":    "b5785ff03785590b5b601b9b4307d0096e5f25db42b140aef79aa9f227c14ca8",
	"libtcl9tkimgps2.0.1.so":     "9a14c28fc624c4595a47bb5c9d475c0648aea041b7dc0d70680378fc69729c48",
	"libtcl9tkimgraw2.0.1.so":    "04eafa1e8d0382321f37ccc38f38d28c067d91c06b545f3db5c94bbe99b85dae",
	"libtcl9tkimgsgi2.0.1.so":    "742e61a26d5b1c8036d0a7e9e8ccc0f2617fcd85db8a142068a8fb13938550fb",
	"libtcl9tkimgsun2.0.1.so":    "fcee90c25a473c6277fa679fc95fa4e8f87dfba59b73c4ac873b49d38213f80a",
	"libtcl9tkimgtga2.0.1.so":    "366e274719d8acad9acca1955bffb8c3480511e8a428deec48973afc0d842306",
	"libtcl9tkimgtiff2.0.1.so":   "9eda1a2b03835795b97c51bd99c15f890fb2c4330895aa6c373fb0848b30a0db",
	"libtcl9tkimgwindow2.0.1.so": "a7eb5f31c1c3052f9afbf59eb8a88e3d616255e28a2857c509f86d6c56a9aca9",
	"libtcl9tkimgxbm2.0.1.so":    "fa6d31f73ad0577bd8224a480f61e4323915d907466b50ed9bc1cd34ef28286b",
	"libtcl9tkimgxpm2.0.1.so":    "44b1abbbf34a4960e71f1dffe55d354d7953dfaf7387358518bb844865c3d7a3",
	"libtcl9zlibtcl1.3.1.so":     "972a92f97932013671ddb561c088759e99ab0b4e045b4804cf0d223f6f318dec",
	"libtk9.0.1.zip":             "0be49823ab336786168ff8dafa0b1c5ed5f0a5cb51b3651021a94d422bb94274",
}
