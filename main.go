package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"./pixelcountapp"
	colors "gopkg.in/go-playground/colors.v1"
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

	pixels := pixelcountapp.ProcessImage(src)

	for _, pixel := range pixels {
		fmt.Printf("%s %d\n", colors.FromStdColor(pixel.PixelColor).ToHEX().String(), pixel.PixelCount)
	}
}
