package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
)

func checkerror(err error) {
	fmt.Println(err)
}

func bufioRead() {
	scanner := bufio.NewScanner(os.Stdin)

	// Scan until user provides a .
	for scanner.Scan() {
		if scanner.Text() == "." {
			break
		}
		fmt.Println(scanner.Text())
	}
}

func ioRead() {
	reader := io.Reader(os.Stdin)
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		checkerror(err)
		fmt.Println(string(buf[:n]))
	}
}

func main() {
	ioRead()
}
