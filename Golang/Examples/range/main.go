package main

import "fmt"

var num = []int{1, 2, 3, 4, 5}

func main() {
	// This is a basic iteration
	for i := 0; i < len(num); i++ {
		fmt.Println(i)
	}

	// The same function with a range
	for index, number := range num {
		fmt.Println(index, number)
	}
}
