package main

import "fmt"
import mat "github.com/meeranh/mathematics"

func main() {
	// Variable declaration
	var num1 int = 10
	var num2 int = 4

	// Preparing variables for storing arithmetic results
	var addSum int
	var subSum int
	var mulSum int
	var divSum int

	// Performing the arithmetic
	addSum = mat.Add(num1, num2)
	subSum = mat.Subtract(num1, num2)
	mulSum = mat.Multiply(num1, num2)
	divSum = mat.Divide(num1, num2)
	
	// Printing out the sums
	fmt.Println("Addition: ", addSum)
	fmt.Println("Subtraction: ", subSum)
	fmt.Println("Multiplication: ", mulSum)
	fmt.Println("Division: ", divSum)

}
