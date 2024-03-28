package main

import "fmt"

// A function that returns a function that prints its parameters
func catParam(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	// Calling the closure
	catParam("This is a closure")()
}
