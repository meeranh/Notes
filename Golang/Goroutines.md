# Goroutines
+ This is basically the multi-threading mechanism of Golang.
+ Each Goroutine only takes up 2KB of memory while a typical OS thread takes up 1MB.
+ In simple words, you can run functions in an asynchronous manner using Goroutines.
+ The following is the syntax to use a Goroutine:
```go
func main() {
    go func() {
        fmt.Println("Hello from Goroutine!") // This will not get printed
    }()
}
```

+ The problem with Goroutines is that the main function will not wait for the Goroutine to finish.
+ Therefore, we have to use some sort of a mechanism to make the Go compiler stop and wait until all the Goroutines finish.
+ We can use `sync.WaitGroup` to achieve this

### How Does `sync.WaitGroup` Work?
+ This is basically like a counter which keeps track of the number of Goroutines in the program.
+ If we created 3 Goroutines, we have to set a counter of 3 using `sync.WaitGroup.Add(3)`.
+ Just having a counter is not enough. We also should call `sync.WaitGroup.Done()` every time a Goroutine finishes. This will force `WaitGroup` to decrement the counter by 1.
+ Finally, we can use `sync.WaitGroup.Wait()` to make the main function wait until all the Goroutines finish.
+ The following code snippet shows how to use `sync.WaitGroup`:
```go
func main() {
    var wg sync.WaitGroup
    wg.Add(3) // Letting WaitGroup know that we have 3 Goroutines

    go func() {
        fmt.Println("Hello from Goroutine 1!")
        wg.Done() // The counter reduces to 2
    }()

    go func() {
        fmt.Println("Hello from Goroutine 2!")
        wg.Done() // The counter reduces to 1
    }()

    go func() {
        fmt.Println("Hello from Goroutine 3!")
        wg.Done() // The counter reduces to 0
    }()

    // Only if the counter is 0, the function continues from here
    wg.Wait()
}
```
