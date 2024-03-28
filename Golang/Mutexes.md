# Race Conditions
+ Imagine you were running 4 threads that write to memory.
+ By chance, imagine those 2 threads write to the memory location 0x0001 at the same time.
+ Most likely, you will get a corrupted value.
+ We want to prevent 2 threads from writing to the same memory location at the same time.
+ **Race conditions** are when 2 threads try to write to the same memory location at the same time.

# Mutexes
+ A Mutex (Mutual Exclusion) is basically like a lock that will prevent 2 threads from writing to the same memory location at the same time.
+ When you are accessing a part of the memory, you have to `Lock()` that memory address.
+ Now lets imagine another thread came across and tried to write to that memory location.
+ Due to that lock, the thread will have to wait until the first thread is done writing to that memory location.
+ The `sync` package in Go has a Mutex type that you can use. This is an example:
```go
func main() {
    var mutex sync.Mutex
    var x int
    go func() {
        for i := 0; i < 1000; i++ {
            mutex.Lock()
            x++ // This is now protected
            mutex.Unlock()
        }
    }()
    go func() {
        for i := 0; i < 1000; i++ {
            mutex.Lock()
            x++ // This is the same memory location as the previous one
            mutex.Unlock()
        }
    }()
    time.Sleep(100 * time.Millisecond)
    fmt.Println(x)
}
```
