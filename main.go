package main

import (
	"log"
	"net/http"

	"./localfiles"
	"./pixelcountapp"
	"github.com/gorilla/mux"
	colors "gopkg.in/go-playground/colors.v1"
)

func main() {
	m := mux.NewRouter()
	m.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		file, err := localfiles.LookupImage("bp")
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "sad!")
		}
		pixelData := pixelcountapp.ProcessImage(file)
		pixelCount := pixelcountapp.TotalPixelsInImage(file)
		respondWithJSON(w, http.StatusOK, ConvertToJSON(pixelData, pixelCount))
	})
	log.Fatal(http.ListenAndServe(":8080", m))
}

// PixelInfoJSON represents a count of pixels of that color
type PixelInfoJSON struct {
	Color   string  `json:"color"`
	Count   uint32  `json:"count"`
	Percent float32 `json:"percent"`
}

// ConvertToJSON fshjcijnsc
func ConvertToJSON(pixels []pixelcountapp.PixelInfo, totalPixels uint) []PixelInfoJSON {
	var exportablePixels []PixelInfoJSON
	for _, pixel := range pixels {
		exportablePixels = append(exportablePixels, PixelInfoJSON{Color: colors.FromStdColor(pixel.PixelColor).ToHEX().String(), Count: pixel.PixelCount, Percent: 100 * float32(pixel.PixelCount) / float32(totalPixels)})
	}
	return exportablePixels
}
