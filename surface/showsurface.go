package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	//fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
	wf, _ := os.Open("surface.svg")
	wf = bufio.NewScanner(wf)
	wf.
		w = *wf
}
