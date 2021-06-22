package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		cherkErr(err)
		//b, err := ioutil.ReadAll(resp.Body)
		i, err := io.Copy(os.Stdout, resp.Body)
		fmt.Println(i)
		resp.Body.Close()
		cherkErr(err)

		fmt.Printf("\n%v\n", resp.Status)

	}
}

func cherkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
}
