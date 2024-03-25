package main

import "fmt"

func main() {
	sumOfVals := proAdder(1, 5, 32, 6)
	fmt.Println(sumOfVals)
}

func proAdder(values ...int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum
}
