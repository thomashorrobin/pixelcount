package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"sort"

	"gopkg.in/go-playground/colors.v1"
)

func main() {

	infile, err := os.Open("images/bp.png")
	if err != nil {
		// replace this with real error handling
		panic("conldn't open file")
	}
	defer infile.Close()

	// Decode will figure out what type of image is in the file on its own.
	// We just have to be sure all the image packages we want are imported.
	src, err := png.Decode(infile)
	if err != nil {
		// replace this with real error handling
		log.Println(err)
		panic("conldn't decode file")
	}

	colorize(src)
}

func colorize(img image.Image) {
	b := image.NewRGBA(img.Bounds())
	draw.Draw(b, b.Bounds(), img, image.ZP, draw.Src)
	// var m map[color.RGBA]uint32

	m := make(map[color.Color]uint32)

	for x := 0; x < b.Bounds().Dx(); x++ {
		for y := 0; y < b.Bounds().Dy(); y++ {
			// log.Println(b.At(x, y).(color.RGBA))
			xx := b.At(x, y)
			m[xx] = m[xx] + 1
		}
	}

	type kv struct {
		Key   color.Color
		Value uint32
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss {
		fmt.Printf("%s %d\n", colorHex(kv.Key), kv.Value)
	}
}

func colorHex(c color.Color) string {
	newColor := colors.FromStdColor(c)
	return newColor.ToHEX().String()
	// return c.(color.RGBAColor).ToHex() //"#" + fmt.Sprintf("%x", c.R.(int)*c.G.(int)*c.B.(int)) // + fmt.Sprintf("%x", c.G) + fmt.Sprintf("%x", c.B)
}
