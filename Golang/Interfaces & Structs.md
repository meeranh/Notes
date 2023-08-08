---
id: "Interfaces & Structs"
aliases:
  - "Structs"
tags: []
---

# Structs

### What Are Structs?
+ Structs are **like classes in other languages**, but they are not classes.
+ Structs create a new type (like `int` or `string`) containing named fields.
+ The following are **structs for a rectangle and a circle**:
```go
type Rectangle struct {
    length float64
    width float64
}

type Circle struct {
    radius float64
}
```

+ On declaring these structs, **we have now created two new types**: `Rectangle` and `Circle`.
+ We can **use this type to create variables** of this type:
```go
var rect1 Rectangle
rect1.length = 5

var circ1 Circle
circ1.radius = 10
```
---

# Interfaces

### What Are Interfaces?
+ Interfaces are like structs, but instead of containing data, they contain **method signatures**.
+ Method signatures are **functions without a body**.
+ Lets assume that `Rectangle` and `Circle` both have an `Area()` method:
```go
type Rectangle struct {
    length float64
    width float64
}

type Circle struct {
    radius float64
}

func (r Rectangle) Area() float64 {
    return r.length * r.width
}

func (c Circle) Area() float64 {
    return math.Pi * math.Pow(c.radius, 2)
}
```

+ Since both a `Rectangle` and a `Circle` have an `Area()` method, we can **create an interface** that catagorises both these structs into an interface called `Shapes`:
```go
type Shapes interface {
    Area() float64
}

func main() {
    var s Shapes
    s = Rectangle{5, 10}
    fmt.Println(s.Area()) // This will print 50
    s = Circle{10}
    fmt.Println(s.Area()) // This will print 314.15
}
```
