package main

import "fmt"

func main() {
	a := make([]int, 10)
	fmt.Println("a: ", a)

	a = append(a, 4)
	fmt.Println("a: ", a)
}
