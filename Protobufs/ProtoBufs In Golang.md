# Protocol Buffers In Golang
+ Assume that you define a `.proto` file and you build it using `protoc`.
+ You now have to use that compiled Protobuf in your Go code.
+ There are some key things to do when importing a Protobuf to your code:
    1. Check if your Protobuf definition is in the same package as your Go code.
    2. Create structs out of the Protobuf definition.
    3. The created structs should be pointers and references.

# Protobuf Package
+ When defining your Protobuf package, you will have to use the `option go_package` directive.
+ The value of `option go_package` should be the same as the package name in your Go code *(go.mod)*.
```protobuf
syntax = "proto3";
option go_package = "github.com/username/repo/package";
package your_package;
```
# Creating Structs Out Of Defined Protobufs
+ Lets assume that you have created your Protobufs in a `/proto` folder.
+ You will have to import that folder with the folder name:
```go
package main
import pb "github.com/username/repo/package/proto"
```

# Creating Pointers And References
+ When creating structs out of Protobufs, you should create pointers and references.
+ You can create a pointer to a struct by using the `&` operator.
+ So lets say you used `user := &pb.User{Id: 1, Name: "John Doe"}`.
+ What is happening is you are creating a new `pb.User` instance and storing it in memory.
+ The memory address of that struct is only being saved into `user`.
```go
package main
import pb "github.com/username/repo/package/proto"

func main() {
    // Create a pointer to a struct
    user := &pb.User{
        Id: 1,
        Name: "John Doe",
    }
}
```
