package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func main() {
	r := Rectangle{width: 10, height: 5}
	p := &r

	fmt.Println("Area: ", p.width)
}
