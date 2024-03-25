## Main Function
+ All Go programs **must have a `main` function**. The code that sits inside the `main` function will be **executed when the program is run**.

```go
func main(){
    // Code goes here
}
```

## Function Syntax
```go
func functionName(parameter type)(returnValue type){
    return value
}
```
+ A function is declared using the `func` keyword.
+ If **inputs** are required, a **parameter name** should be given **with its [[Types|type]]**.
+ If a **return value** is required, a **return type** should be given **with its [[Types|type]]**.

### Function Parameters With Similar Types
+ If for example, you create a function with two or more parameters that take in the same data type, you do not have to explicitly state the type for all those parameters, just mentioning the type of the last parameter will be enough. In this addition function, we do not explicitly specify all parameter types:
```go
func add(x, y int) int {
    return x + y
}
```

### Multiple Return Values
+ If you want to return two or more values and save it to a variable, this is how you do it:
+ Assume you create function that returns two values:
```go
func swapStrings(x, y string) (string, string) {
    return y, x
}
```

+ If you want to store these values into two variables, you need to use the **`:=`** operator.
+ Another name for this is **multiple assignment**.
```go
x, y := swapStrings("Hello", "World")
```

### Named Return Values
+ In the above example, we just returned the variables as parameters to the `return` keyword.
+ There is an alternative method to return values, that is using the function declaration itself.
```go
func AddTen(sum int) (x, y int) {
    x = sum + 10
    y = x + 10
    return
}
```
+ This function then can be called like a normal multiple return value function, and even multiple assignment can be done as usual.

### Functions As Values
+ You can also pass function names as parameters to a function, for example, lets say we have declared an `add()` function:
```go
func add(x, y int) int {
    return x + y
}
```

+ We can pass this function as a parameter to another function:
```go
func compute(fn func(int, int) int) {
    fmt.Println(fn(3, 4))
}

fmt.Println(compute(add))
```
