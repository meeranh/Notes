package main

import (
	"fmt"
)

func randFunc(fn func(int, int) int) int {
	return fn(3, 5)
}

func add(x int, y int) int {
	return x + y
}

func main() {
	sum := randFunc(add)
	fmt.Println(sum)
}
