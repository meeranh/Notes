---
id: "The Go Compiler"
aliases:
  - "Compiling Go To Binaries"
  - "Compiling Go Into Binaries"
tags: []
---

## Compiling Go Into Binaries

+ Assume that we have a file containing Go source named `main.go`.
+ The following command will build `main.go` to an executable named `main`:

```bash
go build main.go
```

## Running Go Without Compiling

+ We **do not have to always build binaries** just to run them.
+ Go has a `go run` command that will run **JIT (Just-In-Time)** compilation without compiling (like Python).

```bash
go run main.go
```
