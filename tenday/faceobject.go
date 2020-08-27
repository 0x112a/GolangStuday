//function
//package main
//
//import "fmt"
//
//type Rectangle struct {
//	width, height float64
//}
//
//func area(r Rectangle) float64{
//	return r.width * r.height
//}
//
//func main(){
//	r1 := Rectangle{12,2}
//	r2 := Rectangle{9,4}
//	fmt.Println("Area of r1 is: ",area(r1))
//	fmt.Println("Area of r2 is: ",area(r2))
//}
package main

import (
	"fmt"
	"math"
)

type Rectange struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectange) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectange{12, 2}
	r2 := Rectange{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}
