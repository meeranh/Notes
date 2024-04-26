package main

import "fmt"

func main() {
	defer fmt.Println("This prints last")
	defer fmt.Println("This prints first")
	fmt.Println("Hello World")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

}
