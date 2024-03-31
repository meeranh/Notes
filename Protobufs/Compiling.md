# Compiling Protobufs
+ Since Protobufs are language agnostic, you can compile them into code for a variety of languages.
+ The `protoc` binary must be available on your system.
+ If you want to compile a `.proto` file to Go, you can use the following command:
```bash
protoc --go_out=. path/to/your/file.proto
```

+ Also, if you want to compile a `.proto` file to Python and Go, you can use the following command:
```bash
protoc --python_out=. --go_out=. path/to/your/file.proto
```
