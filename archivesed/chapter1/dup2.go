package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
<<<<<<< HEAD
=======

>>>>>>> 44b5b901f8793c81080aa5222187ca3e6f7d3cef
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
<<<<<<< HEAD
				fmt.Fprintf(os.Stderr, "dup2: %V\n", err)
=======
				fmt.Fprintf(os.Stderr, "dup2:%\n", err)
				continue
>>>>>>> 44b5b901f8793c81080aa5222187ca3e6f7d3cef
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
<<<<<<< HEAD

=======
>>>>>>> 44b5b901f8793c81080aa5222187ca3e6f7d3cef
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
<<<<<<< HEAD
=======

>>>>>>> 44b5b901f8793c81080aa5222187ca3e6f7d3cef
}
