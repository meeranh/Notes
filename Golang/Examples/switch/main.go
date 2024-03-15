package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch runtime.GOOS {
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Other OS")
	}
}
