package main

import "fmt"

func main() {
	mySliceLiteral := [] int {1, 2, 3, 4, 6, 7}
	firstSlice := mySliceLiteral[0:5]

	fmt.Println(cap(firstSlice))
	fmt.Println(len(firstSlice))
}
