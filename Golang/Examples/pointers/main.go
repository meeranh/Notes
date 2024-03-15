package main

import "fmt"

func main() {
	a := 25 // The original value
	b := &a // A pointer to a

	fmt.Println(a)	// 25
	fmt.Println(b)	// 0xc0000b6010
	fmt.Println(*b) // 25

	// Changing the value of a through the pointer
	*b = 100
	fmt.Println(a) // 100

	// We can change the memory address of the pointer
	b = nil
	fmt.Println(b) // <nil>

	// Also, we can reassign a pointer to another pointer
	c := 69
	d := &c
	b = d
	fmt.Println(b) // 0xc0000b6018
}
