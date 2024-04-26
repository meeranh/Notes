package main

import (
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// Writing to a file
	dataWrite := []byte("My name is Meeran :)\nThis is me learning Go")
	err := os.WriteFile("myFile.txt", dataWrite, 0644)
	check(err)

	// Reading a file
	dataRead, err := os.ReadFile("myFile.txt")
	check(err)
	fmt.Println(string(dataRead))
	}
