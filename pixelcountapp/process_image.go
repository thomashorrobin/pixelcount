package pixelcountapp

import (
	"image"
	"image/color"
	"image/draw"
	"sort"
)

// ProcessImage takes image.Image object and returns pixel counts
func ProcessImage(img image.Image) []PixelInfo {
	b := image.NewRGBA(img.Bounds())
	draw.Draw(b, b.Bounds(), img, image.ZP, draw.Src)

	m := make(map[color.Color]uint32)

	for x := 0; x < b.Bounds().Dx(); x++ {
		for y := 0; y < b.Bounds().Dy(); y++ {
			xx := b.At(x, y)
			m[xx] = m[xx] + 1
		}
	}

	var ss []PixelInfo
	for k, v := range m {
		ss = append(ss, PixelInfo{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].PixelCount > ss[j].PixelCount
	})

	return ss
}

// PixelInfo represents a count of pixels of that color
type PixelInfo struct {
	PixelColor color.Color
	PixelCount uint32
}

// TotalPixelsInImage uhisbsdc
func TotalPixelsInImage(img image.Image) uint {
	rec := img.Bounds()
	return uint(rec.Dx()) * uint(rec.Dy())
}
