# Serializing A Protobuf To Wire Format
+ Basically, before sending a protobuf message over the wire, it needs to be serialized to a wire format. This is done by encoding the message to a byte array.
+ Lets say that we have the following Protobuf definition:
```protobuf
message Person {
  required string name = 1;
  required int32 id = 2;
  optional string email = 3;
}
```

+ And we are trying to save it to a file. The following code will serialize the payload:
```go
func main() {
    p := &Person{
        Name: "Meeran Hassan",
        Id: 1,
        Email: "meeran@surge.global"
    }

    // Convert the message to a byte array
    out, err := proto.Marshal(p)
    checkErr(err)

    // Write the byte array to a file
    err = os.WriteFile("person.bin", out, 0644)
    checkErr(err)
}
```

# De-Serializing A Protobuf To Wire Format
+ Assume that we wanted to read the `bin` file we created earlier.
+ To do this, we need to:
    1. Create an empty Protobuf with `emptypb.Empty{}`.
    2. Read the bin file with `os.ReadFile()`.
    3. Deserialize the byte array with `proto.Unmarshal()`.
```go
func main() {
    // Read the file into a byte array
    in, err := os.ReadFile("person.bin")
    checkErr(err)

    // Create a new Person object
    p := &emptypb.Empty{}

    // De-serialize the bytes and store it in p
    err = proto.Unmarshal(in, p)
    checkErr(err)

    fmt.Println(p)
}
```
