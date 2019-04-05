package main

import (
	"log"
	"net/http"

	"./localfiles"
	"./pixelcountapp"
	"github.com/gorilla/mux"
)

func main() {
	m := mux.NewRouter()
	m.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		file, err := localfiles.LookupImage("bp")
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "sad!")
		}
		xxxx := pixelcountapp.ProcessImage(file)
		respondWithJSON(w, http.StatusOK, xxxx)
	})
	log.Fatal(http.ListenAndServe(":8080", m))
}

// type xxx struct {
// 	Name string `json:"name"`
// }
