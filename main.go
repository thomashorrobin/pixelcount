package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
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

	m := make(map[color.RGBA]uint32)

	for x := 0; x < b.Bounds().Dx(); x++ {
		for y := 0; y < b.Bounds().Dy(); y++ {
			// log.Println(b.At(x, y).(color.RGBA))
			xx := b.At(x, y).(color.RGBA)
			m[xx] = m[xx] + 1
		}
	}
	log.Println(m)
}
