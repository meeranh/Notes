package main

import "fmt"
import "math"

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (c Circle) area() float64 {
	return 2 * c.radius * math.Pi
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func main() {
	myCircle := Circle{radius: 5}
	myRectangle := Rectangle{length: 6, width: 9}

	myShapes := []Shape{myCircle, myRectangle}

	for _, v := range myShapes {
		fmt.Println(v.area())
	}
}
