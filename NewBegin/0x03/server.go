package main

import (
	"GolangStuday/NewBegin/0x03/surface"
	"net/http"
)

func main() {

	http.HandleFunc("/", surfaceHandler)

	http.ListenAndServe("localhost:8000", nil)
}

func surfaceHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "image/svg+xml")
	surface.Surface(w)

}
