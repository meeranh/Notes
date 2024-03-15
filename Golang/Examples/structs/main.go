package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func main() {
	var myRectangle Rectangle = Rectangle{width: 10, height: 5} // Explicit syntax
	anotherRectangle := Rectangle{10, 5} // Shorthand syntax

	fmt.Println(myRectangle)
	fmt.Println(anotherRectangle.width)
}
