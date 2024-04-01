# Oneof
+ Lets say you are defining a Protobuf and you need a special kind of field:
    + This field can take in one of the following types: int, string, bool, etc.
+ You can create such a field using the `oneof` keyword:

```proto
syntax = "proto3";

message MyMessage {
    oneof gender {
        bool is_male = 1;
        string custom = 2;
    }
}
```

+ What the above code will do is create a field called `gender` which can be either a `bool` or a `string`.
+ When you call this in Go, you can pass in a struct of either type and it will work.

```go
package main

func main() {
    myGender := &MyMessage{
        Gender: &MyMessage_IsMale{IsMale: true},
    }

    myCustomGender := &MyMessage{
        Gender: &MyMessage_Custom{Custom: "Apache Attack Helicoptor"}
    }

    fmt.Println(myGender)       // Gender: {IsMale:true}
    fmt.Println(myCustomGender) // Gender: {Custom:Apache Attack Helicoptor}
}
```
