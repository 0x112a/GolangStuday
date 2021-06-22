package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// begin time
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
		fmt.Println()
		fmt.Println(<-ch)

	}
	fmt.Printf("%.dms elapsed\n", time.Since(start).Milliseconds())

}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	checkErr(err)
	nbytes, err := io.Copy(os.Stdout, resp.Body)
	checkErr(err)
	resp.Body.Close()
	secs := time.Since(start).Milliseconds()
	ch <- fmt.Sprintf("%dms %7d %s", secs, nbytes, url)

}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stdout, "fecth :%v", err)
		os.Exit(1)
	}
}
