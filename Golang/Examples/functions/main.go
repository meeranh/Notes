package main

import "fmt"

// Multiple return values
func swapStrings(x, y string) (string, string) {
	return y, x
}

// Named return values
func addTen(sum int) (x, y int) {
    x = sum + 10
    y = x + 10
		return
}

func main() {
	// Multiple assignment
	x, y := swapStrings("Hello", "World")
	a, b := addTen(20)

	fmt.Println("Multiple return values: ", x, y)
	fmt.Println("Named return values: ", a, b)
}
