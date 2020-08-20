package main

import (
	"fmt"
)

func sun(a []int, c chan int) {
	total := 0
	for _, V := range a {
		total += V
	}
	c <- total
	fmt.Println(a)
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sun(a[:len(a)/2], c)
	go sun(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
