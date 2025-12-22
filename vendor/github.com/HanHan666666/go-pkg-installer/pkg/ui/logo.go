package ui

import (
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
	"path/filepath"
	"strings"

	. "modernc.org/tk9.0"

	"github.com/HanHan666666/go-pkg-installer/pkg/core"
)

func logoKeyFromContext(ctx *core.InstallContext) string {
	if ctx == nil {
		return ""
	}
	if data := embeddedLogoBytes(ctx); len(data) > 0 {
		return "embedded"
	}
	path := resolveAssetPath(ctx, ctx.RenderOrDefault("product.logo", ""))
	if strings.TrimSpace(path) == "" {
		return ""
	}
	return path
}

func loadLogoForKey(ctx *core.InstallContext, key string) *Img {
	if key == "" {
		return nil
	}
	if key == "embedded" {
		return loadLogoImageFromData(embeddedLogoBytes(ctx))
	}
	return loadLogoImage(key)
}

func resolveAssetPath(ctx *core.InstallContext, logoPath string) string {
	path := strings.TrimSpace(logoPath)
	if path == "" || filepath.IsAbs(path) {
		return path
	}
	if ctx == nil {
		return path
	}
	base := strings.TrimSpace(ctx.RenderOrDefault("config.dir", ""))
	if base == "" {
		return path
	}
	candidate := filepath.Join(base, path)
	if _, err := os.Stat(candidate); err == nil {
		return candidate
	}
	return path
}

func loadLogoImage(path string) *Img {
	if path == "" {
		return nil
	}
	if _, err := os.Stat(path); err != nil {
		return nil
	}
	var img *Img
	func() {
		defer func() {
			if recover() != nil {
				img = nil
			}
		}()
		if scaled := loadScaledLogo(path, sidebarLogoSize); scaled != nil {
			img = scaled
			return
		}
		img = NewPhoto(File(path))
	}()
	return img
}

func embeddedLogoBytes(ctx *core.InstallContext) []byte {
	if ctx == nil {
		return nil
	}
	embedData, ok := ctx.Get("product.logo.bytes")
	if !ok {
		return nil
	}
	data, ok := embedData.([]byte)
	if !ok || len(data) == 0 {
		return nil
	}
	return data
}

// loadLogoImageFromData loads a logo image from byte data
func loadLogoImageFromData(data []byte) *Img {
	if len(data) == 0 {
		return nil
	}
	var img *Img
	func() {
		defer func() {
			if recover() != nil {
				img = nil
			}
		}()
		if scaled := loadScaledLogoFromData(data, sidebarLogoSize); scaled != nil {
			img = scaled
			return
		}
		img = NewPhoto(Data(data))
	}()
	return img
}

// loadScaledLogoFromData loads and scales a logo from byte data
func loadScaledLogoFromData(data []byte, target int) *Img {
	if target <= 0 || len(data) == 0 {
		return nil
	}

	reader := bytes.NewReader(data)
	src, _, err := image.Decode(reader)
	if err != nil {
		return nil
	}

	scaled := scaleToFit(src, target)
	if scaled == nil {
		return nil
	}

	return NewPhoto(Data(scaled))
}

func loadScaledLogo(path string, target int) *Img {
	if target <= 0 {
		return nil
	}
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	src, _, err := image.Decode(file)
	if err != nil {
		return nil
	}

	scaled := scaleToFit(src, target)
	if scaled == nil {
		return nil
	}

	return NewPhoto(Data(scaled))
}

func scaleToFit(src image.Image, target int) image.Image {
	if src == nil || target <= 0 {
		return nil
	}
	bounds := src.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	if width <= 0 || height <= 0 {
		return src
	}
	maxDim := width
	if height > maxDim {
		maxDim = height
	}
	if maxDim <= target {
		return src
	}

	scale := float64(target) / float64(maxDim)
	newW := int(math.Round(float64(width) * scale))
	newH := int(math.Round(float64(height) * scale))
	if newW < 1 {
		newW = 1
	}
	if newH < 1 {
		newH = 1
	}

	dst := image.NewRGBA(image.Rect(0, 0, newW, newH))
	xRatio := float64(width) / float64(newW)
	yRatio := float64(height) / float64(newH)
	for y := 0; y < newH; y++ {
		srcY := int(float64(y) * yRatio)
		if srcY >= height {
			srcY = height - 1
		}
		for x := 0; x < newW; x++ {
			srcX := int(float64(x) * xRatio)
			if srcX >= width {
				srcX = width - 1
			}
			dst.Set(x, y, src.At(bounds.Min.X+srcX, bounds.Min.Y+srcY))
		}
	}
	return dst
}
