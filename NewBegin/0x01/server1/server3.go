package main

import (
	"GolangStuday/NewBegin/0x01/lissajous"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	http.HandleFunc("/test", liss)
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Printf("aa")
		fmt.Fprintf(w, "Form[%T] = %s\n", k, strings.Join(v, ""))
	}

}

func liss(w http.ResponseWriter, r *http.Request) {
	var cycles float64

	r.ParseForm()
	for k, v := range r.Form {
		if k == "cycles" {
			t := strings.Join(v, "")
			fmt.Println(t)
			cycles, _ = strconv.ParseFloat(t, 64)
		}
	}

	lissajous.Lissajous(w, cycles)

}
