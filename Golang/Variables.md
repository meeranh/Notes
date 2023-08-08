---
id: "Variables"
aliases:
  - "Variable Declaration"
tags: []
---

## Variable Declaration

+ Variables are declared using the `var` keyword.
+ The syntax it follows is `var nameOfVariable typeOfVariable`
+ An example of variable declaration is:
```go
var name string
var age uint32
var height float32
var chad bool
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
}
```
