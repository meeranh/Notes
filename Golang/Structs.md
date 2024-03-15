# Structs

## Creating Structs
+ Structs are like **classes** in other languages. This is how you declare a struct:
```go
type Rectangle struct {
	Width  float64
	Height float64
}
```
+ Like with functions, the first letter of the struct name/field should be capitalized if you want to export it from the package.

## Struct Literals
+ A struct literal denotes a newly allocated struct value by listing the values of its fields using the `{}` syntax.
+ The `{}` and the values inside it are called a **struct literal**.
```go
func main() {
    // This is a struct literal
    r := Rectangle{10, 5}
    fmt.Println(r) // {10 5}
}
```

+ When creating an instance (or object) of a struct, we can do this either of the following ways:
```go
func main() {
    // Explicit syntax
	var myRectangle Rectangle = Rectangle{width: 10, height: 5}

    // Shorthand syntax
	anotherRectangle := Rectangle{10, 5}

    // Struct fields are accessed using a dot
    fmt.Println(myRectangle.width) // 10
    fmt.Println(anotherRectangle.height) // 5
    fmt.Println(myRectangle) // {10 5}
}
```

## Structs & Pointers
+ If we had a struct `v`, and a pointer `p` pointing to it, to to dereference the pointer and access the struct's fields, we would use the `(*p).field` syntax.
+ However, Go does not require to write such cumbersome code for dereferencing a struct pointer.
+ Instead of using `(*p).field`, we can use `p.field` to access the fields of the struct as Go will understand that we are trying to access the field of the struct that the pointer is pointing to.
+ The following syntax can be used:
```go
func main() {
    r := Rectangle{10, 5}
    p := &r

    // This is the same as (*p).width
    fmt.Println(p.width) // 10
}
```
