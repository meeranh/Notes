# For Loops in Go
+ Go has only one iterative statement, the `for` loop.
+ Below is the basic syntax of a for loop:
```go
for i := 0; i < 10; i++ {
    // An operation
}
```
+ The structure of the `for` loop is similar to other languages, we have 3 components:
    1. **Initializer**:  `i := 0`
        + The iterative variable is defined here. It should use the `:=` operator.
    2. **Condition**: `i < 10`
        + The loop will continue to run as long as this condition is true.
    3. **Post statement**: `i++`
        + This is the operation that will be performed after each iteration.

# Go's While Loop
+ All three components of the `for` loop are optional.
+ You do not have to specify `i := 0` and `i < 10` in the `for` loop. An example of this is:
```go
for ; i < 10 ; {
    i++
}
```

+ This can be rewritten without the semicolons imitating a while loop:
```go
func main() {
    for i < 10 {
        i++
    }
}
```

# Go's Infinite Loop
+ If you do not specify any parameters to the `for` loop, you basically get an infinite loop:
```go
func main() {
	for {
	}
}
```

# For Loops With `range`
+ Lets say we want to perform an operation on each value of an array, this is a basic iteration:
```go
var num = []int{1, 2, 3, 4, 5}

func main() {
	for i := 0; i < len(num); i++ {
		fmt.Println(i)
	}
}
```

+ However, we can improve the syntax using the `range` keyword:
```go
var num = []int{1, 2, 3, 4, 5}

func main() {
	for index, value := range num {
		fmt.Println(index, value)
	}
}
```

+ In a for loop, the `range` keyword will act as a function which returns two values, the index, and the value at the current iteration.
+ If you want to omit a variable, the `_` variable name can be assigned and it will be ignored.
```go
for _, value := range num {
    fmt.Println(value)
}
```

# Iterating Maps With `range`
+ When iterating a map using the `for` keyword, the `range` keyword must be used to fetch both keys and their values at a given time in iteration.
+ The following is the syntax for iterating a map.
```go
func main() {
    varName := map[keyType]valueType{
        key1: value1,
        key2: value2,
        key3: value3,
    }

    for key, value := range varName {
        fmt.Println(key, value)
    }
}
```
