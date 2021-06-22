package main

import (
	"GolangStuday/NewBegin/0x01/lissajous"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("b0x112:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.
		lissajous.Lissajous(w)

}
