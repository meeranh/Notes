# The Defer Keyword

+ There are two thing `defer` does:
    1. In a function, if a `defer` statement is used, the function call will be moved to the end of the function. For example:
    ```go
    func main() {
        defer fmt.Println("This will be printed last")
        fmt.Println("This will be printed first")
    }
    ```
    2. Also, if two `defer` functions are called, the one that is called first will be executed last (LIFO). For example:
    ```go
    func main() {
        defer fmt.Println("This will be printed last")
        defer fmt.Println("This will be printed second")
        fmt.Println("This will be printed first")
    }
    ```
