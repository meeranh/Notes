# Deadlocks
+ To understand channels, we first need to understand deadlocks.
+ In [[Mutexes]], we covered that if a thread tries to access a memory location:
    1. The memory location is first locked.
    2. The thread makes the changes to the memory location.
    3. Finally the thread releases the mutex lock.
    4. While locked, no other thread can access the memory location.

+ Imagine the following situation:
    1. Thread 1 locks memory location A.
    2. Thread 2 locks memory location B.
    3. Thread 1 tries to access memory location B.
    4. Thread 2 tries to access memory location A.

+ Both threads are now waiting for each other to release the mutex lock. This is called a **deadlock**.

# Channels
+ Channels are a convenient way to share data between [[Goroutines]].
+ If you create a channel in Goroutine A, you can pass data to Goroutine B through that channel.
+ The catch here is that channels should always be created inside a Goroutine. If you don't, you will get a deadlock.

### Channel Senders & Receivers
+ Channel senders are usually the Goroutines that create the channel. The syntax for a channel sender is `ch <- 42`.
+ Channel receivers are usually the Goroutines that receive the data from the channel. The syntax for a channel receiver is `fmt.Println(<-ch)`.
+ While **channel receivers have the arrow pointing towards the channel**, **channel senders have the arrow pointing away from the channel**.
+ You can imaging channels as a box where you can put data in one end and take it out from the other end.

+ The following is a basic usage of channels:
```go
func main() {
    ch := make(chan int)
    go func() {
        ch <- 42
    }()
    fmt.Println(<-ch) // 42
}
```

 + This will not work:
```go
func main() {
    ch := make(chan int)
    ch <- 42
    fmt.Println(<-ch) // deadlock
}
```

### Buffered Channels
+ Buffered channels are channels that can store a certain number of values.
+ If you create a buffered channel with a capacity of 5, you can store 5 values in the channel without blocking the Goroutine.
+ This is how you create a buffered channel:
```go
ch := make(chan int, 5) // The second argument is the capacity of the channel.
```

### Is A Buffer Open/Closed?
+ You can check if a channel is open or closed by using the following syntax:
```go
func main() {
    ch := make(chan int)
    close(ch)
    _, ok := <-ch
    fmt.Println(ok) // false
}
```
