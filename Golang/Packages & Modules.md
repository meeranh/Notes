## What Are Packages In Go
+ Packages are **very similar to libraries in Python**.
+ **Folders containing Go source files** are called packages.
+ Its a **way of organizing code related to a specific functionality** into separate packages.
+ The following code is an example of *package declaration* and *package import*

```go
// This is a package declaration
package main 
 
// This is a package import
import "fmt" 
 
func main () {
  fmt.Println("Hello World") 
}
```

## Package Declaration
+ The first line of code in the above Go program contains a `package main`.
+ What this means is that the code in this file belongs to the `main` package.

## Package Import
+ The second line of code in the above Go program contains an `import "fmt"`.
+ `fmt` is a **built-in package** in Go which contains functions for printing to the console.
+ `import` allows us to **use code from other packages** in our program.
+ All imports should be wrapped in **double quotes**.

### Importing Multiple Packages
+ If we want to import multiple packages, we call `import` and specify the names of the packages separated by a newline.
```go
import (
  "fmt"
  "math"
)
```

### Import Aliases
+ This is similar to `import numpy as np` in Python.
+ A **short name** can be given to a package at import.
+ In the code snippet below, **`m` is an alias for the `math` package**.
+ `m.Pi` is called instead of `math.Pi`.

```go
import (
  m "math"
  "fmt"
)

func main(){
    fmt.Println(m.Pi)
}
```

# The `main` Package
+ If you are creating a binary executable, or a program that you want to execute, a `main` package with a `main()` function should be created as this is the entrypoint to the program.

# Creating Packages & Modules
+ To create a package, create a folder with the name of the package and place the Go source files in the folder. For example, if you wanted to create a `mathematics` package with functions such as `Add()`, `Subtract()`, `Divide()`, and `Multiply`, you would need to create a `mathematics` folder with the Go source files as follows:
```unix
mathematics
├── addition.go
├── division.go
├── multiplication.go
└── subtraction.go
```

### Structuring Modules
+ The files under the `mathematics` folder are called **Modules**. Modules are simply the sub-components of a package. In other words, a package is a collection of modules.

+ Also, all these modules should start with `package mathematics` as that is the only way to assign these functions to the package.

+ The philosophy of a module is that it should be able to do only one operation. For example, `addition.go` should only be able to do functions related to addition, nothing more. Modularizing code in such a way will allow developers to create a codebase that is very easy to maintain.

+ The following would be the contents of the `addition.go` and `subtraction.go` files:

```go
// addition.go
package mathematics

func Add(a, b int) int {
    return a + b
}
```

```go
// subtraction.go
package mathematics

func Subtract(a, b int) int {
    return a - b

}
```

# Exporting functions
+ There is not specific keyword to let Go know that you want to export a function in a package, you just have to **capitalize the first letter of the function** if you want to export a function. For example:
```go
func ThisWillGetExported() {
    // Stuff here
}
```
```go
func thisWillNotGetExported() {
    // Stuff here
}
```

# Importing packages locally

+ Usually, packages are meant to be imported from remote sources such as `import github.com/meeranh/math`. However, if we did not want to host our packages, and just wanted to source it locally, this are the steps we need to follow:
    1. Create a folder for your package wherever you want in your file system
    2. Run `go mod init https://github.com/username/package` in both your root folder and package folder *(Make sure your root and package directory URLs are different)*
    3. Edit the `go.mod` file in your root folder and add the following lines
        ```go
        replace github.com/username/package => ./path/to/package/folder
        ```
    4. Run `go mod tidy`
    5. Finally, `import github.com/username/package` to your main Go file and run it.

+ What this process basically does is it remaps a remote location URL to a local file path. By doing this, Go will reroute itself to this local path whenever the remote URL is called.

# Sub-packages

+ This is when you have a package inside another package. The best example of a sub-package is `math/rand`.
+ The directory structure of the `math/rand` sub-package would look like this:
```unix
math/
    rand/
        rand.go
    math.go
```
