# Pointers
## Pointer Creation & Dereferencing
+ These are memory addresses of variables.
+ There are two special functions of pointers:
    1. `&` - This is used to get the memory address of a variable.
    2. `*` - This is used to **dereference** a memory address.

+ Dereferencing a pointer is simply to get the value of the variable that the pointer is pointing to.

+ The below example shows how to create & dereference a pointer:
```go
func main() {
	a := 25 // The original value
	b := &a // A pointer to a

	fmt.Println(a)	// 25
	fmt.Println(b)	// 0xc0000b6010
	fmt.Println(*b) // 25
}
```

+ We can change the value of a variable through one of its pointers:
```go
func main() {
	a := 25 // The original value
	b := &a // A pointer to a

	*b = 100
	fmt.Println(a) // 100
}
```

+ Finally, we can assign custom memory locations to reassign pointers:
```go
func main() {
	a := 25 // The original value
	b := &a // A pointer to a

	// We can change the memory address of the pointer
	b = nil
	fmt.Println(b) // <nil>

	// Also, we can reassign a pointer to another pointer
	c := 69
	d := &c
	b = d
	fmt.Println(b) // 0xc0000b6018
}
```
