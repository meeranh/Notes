## Variable Declaration

### Single Variable Declaration
+ Variables are declared using the `var` keyword.
+ The syntax it follows is `var nameOfVariable typeOfVariable`
+ An example of variable declaration is:
```go
var name string
var age uint32
var height float32
var chad bool
```

### Multiple Variable Declaration
+ Multiple variables can be declared in a single line.
+ The syntax it follows is `var nameOfVariable1, nameOfVariable2 typeOfVariable`
```go
var fname, lname string
```

## Multiple One-Line Declarations
+ Variables can also be declared in a single line.
```go
name, age := "Meeran", 21
```

## Zero Values
+ Variables are initialized with a default zero value.
+ Unlike in C, we don't have to explicitly initialize variables to zero.
```go
func main(){
    var x int
    var y float32
    var z string
    var a bool

    fmt.Println(x) // Prints 0
    fmt.Println(y) // Prints 0
    fmt.Println(z) // Prints ""
    fmt.Println(a) // Prints false
}
```

## Type Inference
+ Go is a **statically typed language**.
+ This means that **variables always have a specific type**.
+ However, **Go can infer the type of a variable** based on the value assigned to it.
```go
func main(){
    x := 5             // Integer
    y := 5.5           // Float
    z := "Hello World" // String

    // Also works with the var keyword if a type is not specified
    var a = 5               // Integer
    var b = 5.5             // Float
    var c = "Hello World"   // String
}
```

## Constants
+ These are variables whose values cannot be changed.
+ They are declared using the `const` keyword.
```go
func main(){
    const x int = 5
    const y float32 = 5.5
    const z string = "Hello World"
    const a bool = true
}
```

## Integers & 32/64-bit Systems
+ The size of an integer **depends on the system architecture**.
+ On a **32-bit system**, an **integer is 32 bits** long.
+ On a **64-bit system**, an **integer is 64 bits** long.
+ If Go compiled a program on a 64-bit system, it **will convert all `int` data types to `uint`**.
+ This is because **`int` is 32 bits** in size, but **`uint` is 64 bits** in size.
+ It is better to use `int` instead of `uint` as it will make the code **more portable**.

## Variables With Initializers
+ If you want to initialize multiple variables with values, this is how you do it:
```go
var x, y, z int = 5, 10, 15
```
+ The compiler also can automatically identify the datatype:
```go
var x, y, z = "Cat", true, 1
```

## Short Assignment
+ The *short assignment* operator denoted by *:=* can be used inside a function. Anything that is declared outside a function will require a keyword such as `var` or `func`.
+ The short assignment operator does not require an explicit type.
```go
func main() {
    k := 5
    x, y, z := 1, 2, "Hello"
}
```

## IOTA (Incremental Constants)
+ `iota` is a feature where after you declare one constant:
    + The next constant will be incremented by 1
    + The following constant will be incremented by 2
    + The following constant will be incremented by 3
    + And so on...
```go
const (
    zero int = iota
    one
    two
)

func main(){
    fmt.Println(zero) // Prints 0
    fmt.Println(one)  // Prints 1
    fmt.Println(two)  // Prints 2
}
```

# Numeric Constants via Binary Shifts
+ In Go, you can declare constants using [[Binary Shifts]]
```go
const (
    KB = 1 << (100)
    MB = 1 << (200)
    GB = 1 << (300)
)

func main() {
    fmt.Println(KB) // Prints 1024
    fmt.Println(MB) // Prints 1048576
    fmt.Println(GB) // Prints 1073741824
}
```
+ Its basically a binary shift (multiplication of 2), In other words, the `<<` operator will add 100 zeroes to KB, 200 zeroes to MB, and 300 zeroes to GB.
