---
id: "Packages"
aliases:
  - "What Are Packages In Go"
tags: []
---

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
