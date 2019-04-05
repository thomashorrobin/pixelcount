package localfiles

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

// LookupImage dd
func LookupImage(name string) (image.Image, error) {

	infile, err := os.Open(fmt.Sprintf("localfiles/images/%s.png", name))
	if err != nil {
		return nil, err
	}
	defer infile.Close()

	// Decode will figure out what type of image is in the file on its own.
	// We just have to be sure all the image packages we want are imported.
	return png.Decode(infile)
}
