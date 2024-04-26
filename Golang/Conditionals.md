# If-Else Statements
+ If statements require a *condition*, and it does not require it to be wrapped in parentheses:
```go
func main() {
    youAreCool := true

    if youAreCool {
        fmt.Println("I am so cool :D")
    }
}
```

+ The following is how the `else` keyword can be used along with `if`:
```go
func main() {
    youAreCool := true

    if youAreCool {
        fmt.Println("I am so cool :D")
    } else {
        fmt.Println("Nobody loves you")
    }
}
```

# If With A Short Statement
+ You can execute a small statement before the if condition and save its output to a variable.
+ This variable then can be checked for a certain condition. This can reduce the amount of code you write:
```go
func main() {
    if x := 10; x > 5 {
        fmt.Println(x, " is greater than 5")
    } else {
        fmt.Println(x, " is not greater than 5")
    }
}
```
+ The first part of the if statement is basically the execution of the statement `x := 10`, and then the condition `x > 5` is checked.
+ The scope of the variable `x` is limited to the `if` and `else` block. Meaning, you cannot use the `x` variable outside of the `if` and `else` block.

# Switch Statements
+ Switch statement have the following structure:
```go
switch ValueToCheck {
    case 'SomeValue':
        // Do something
    case 'NotSomeValue':
        // Again, do something
    case default:
        // If nothing matches
}
```
+ If one of the `case`s resolves to true, the `case`s under it will not be executed.

## Switch With No Condition
+ This is basically a switch statement without a value provided to check such as the following:
```go
func main() {
    switch {
        case false:
            fmt.Println("This will not print")
        case true:
            fmt.Println("This will print")
    }
}
```
+ This can be useful to refactor multiple `if-else` statements into a single `switch` statement.

