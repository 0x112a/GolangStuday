package main

import "fmt"

func main() {
	items := make([]map[int]int, 10)
	for _, item := range items {
		item = make(map[int]int, 1) // Oops! item is only a copy of the slice element.
		item[1] = 2                 // This 'item' will be lost on the next iteration.
	}
	fmt.Println(items)
	item := make([]map[int]int, 10)
	for i := range items {
		item[i] = make(map[int]int, 1)
		item[i][1] = 2
	}
	fmt.Println(item)
}
