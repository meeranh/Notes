package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 1; i <= 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	p := 2.0
	fmt.Println(Sqrt(p))
}
