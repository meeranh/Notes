# What are buffers?
+ Imagine you were dealing with a slow I/O medium such as:
    1. A slow network connection
    2. A slow HDD disk

+ The very basic fundamental buffer that all buffer packages in Go use is `make([]byte, size)`.

+ Buffers allow you to collect data and store it in the RAM.
+ This will reduce the need to access the slow I/O medium every time you need to read or write data.

# Buffer handling in `bytes`
+ The `bytes` package provides a buffer implementation.
```go
func main() {
	// This is a buffer that can grow in size
	var buffer bytes.Buffer

	buffer.WriteString("Hello!")
	fmt.Println(buffer.String()) // Hello!

	buffer.WriteString(" How are you?")
	fmt.Println(buffer.String()) // Hello! How are you?
}
```

# Buffer handling in `bufio`
+ The `bufio` package provides a buffered I/O implementation.
+ The following example shows how you can use `bufio` to read from a file:
```go
func main() {
    file, _ := os.Open("file.txt")

    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text()) // Print each line of the file
    }
}
```

+ You can also read stdin using `bufio`:
```go
func main() {
    // Stdin is treated as a buffer
	scanner := bufio.NewScanner(os.Stdin)

	// Scan until user provides a STOP
	for scanner.Scan() {
		if scanner.Text() == "STOP" {
			break
		}
		fmt.Println(scanner.Text())
	}
}
```
