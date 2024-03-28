package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func bytesBuffers() {

	// This is a buffer that can grow in size
	var buffer bytes.Buffer

	buffer.WriteString("Hello, this is Meeran!")
	fmt.Println(buffer.String())

	buffer.WriteString("\nThis is another line")
	fmt.Println(buffer.String())
}

func bufioBuffers() {

	// Writing to a file
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error creating file")
		return
	}

	defer file.Close()

	file.WriteString("Hello, this is Meeran!")
	file.WriteString("\nThis is another line")

	// Reading from a file
	file, err = os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
}
