---
id: "Types"
aliases:
  - "Why Types Are Important"
  - "Why Types Are Important In Go"
tags: []
---

## Why Types Are Important In Go
+ Types are important because they **tell the compiler how to interpret the data**.
+ For example, the following code will **not compile**:
```go
package main

import "fmt"

func main(){
    var x = 5
    var y = 5.5
    fmt.Println(x + y)
}
```
+ The compiler will throw an error because **Go is a strongly typed language**.

## All Types In Go
1. **Numeric Types**
    + **Integers**: These are whole numbers.
    ```go
    func main(){
        // Integers can be positive or negative

        var x int = 5                       // An ordinary integer
        var y int8 = 127                    // 8 bit integer
        var z int16 = 32767                 // 16 bit integer
        var a int32 = 2147483647            // 32 bit integer
        var b int64 = 9223372036854775807   // 64 bit integer
        }
    ```

    + **Unsigned Integers**: These are whole numbers that can only be positive.
    ```go
    func main(){
        // Unsigned integers can only be positive

        var x uint = 5                      // An ordinary unsigned integer
        var y uint8 = 255                   // 8 bit unsigned integer
        var z uint16 = 65535                // 16 bit unsigned integer
        var a uint32 = 4294967295           // 32 bit unsigned integer
        var b uint64 = 18446744073709551615 // 64 bit unsigned integer
        }
    ```

    + **Floating Point Numbers**: These are numbers with a decimal point.
    ```go
    func main(){
        // Floating point numbers have decimal points

        var x float32 = 5.5 // 32 bit floating point number
        var y float64 = 5.5 // 64 bit floating point number
    }
    ```

    + **Complex Numbers**: These are numbers with a real and imaginary part.
    ```go
    func main(){
        // Complex numbers contain a real and imaginary part

        var x complex64 = 1 + 2i // 32 bit complex number
        var y complex128 = 1 + 2i // 64 bit complex number
    }
    ```
2. **Booleans**
    + These are values that can be either `true` or `false`.
    ```go
    func main(){
        var x bool = true
        var y bool = false
    }
    ```

3. **Strings**
    + These are sequences of characters.
    ```go
    func main(){
        var x string = "Hello World"
    }
    ```

4. **Derived Types**
    + These are types that are derived from other types.
    ```go
    func main(){
        // Arrays
        var x [5]int // An array of 5 integers

        // Slices
        var y []int // A slice of integers (a dynamic array)

        // Structs
        type Person struct{
            name string
            age int
        }
        var z Person // A struct of type Person
        z.name = "John"
        z.age = 30

        // Maps (Python dictionaries in Go)
        var a map[string]int // Key is a string, value is an integer

        // Pointer
        var b *int
        b = &x // The pointer b points to the address of x

        // Interfaces (Grouping of types with a common method)
        type Shape interface{
            area() float64 // A circle and a rectangle can both have an area
        }
        
        // A circle has an area
        func (c Circle) area() float64{
            return math.Pi * c.radius * c.radius
        }
        // A rectangle has an area
        func (r Rectangle) area() float64{
            return r.length * r.width
        }
        // So, a circle and a rectangle are both shapes
        func getArea(s Shape) float64{
            return s.area()
        }
    }
    ```

5. **Type Conversions**
    + These are conversions from one type to another.
    ```go
    func main(){
        // Converting an integer to a float
        var x int = 5
        var y float64 = float64(x)

        // Converting a float to an integer
        var z float64 = 5.5
        var a int = int(z)
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
