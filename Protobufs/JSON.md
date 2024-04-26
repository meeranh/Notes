# Protobuf To JSON
+ Using the `protojson` package, we can use the `protojson.Marshall()` function to convert a Protobuf to JSON.
+ The syntax for this function is as follows:
```go
func main() {
    // Create a new instance of the Protobuf message
    message := &pb.Message{
        Field1: "Hello",
        Field2: 42,
    }

    // Convert the Protobuf message to JSON
    json, err := protojson.Marshal(message)
    if err != nil {
        log.Fatalf("Failed to marshal message: %v", err)
    }

    // Print the JSON
    fmt.Println(string(json))
}
```

# JSON To Protobuf
+ We can do the reverse also using the `protojson.Unmarshal()` function.
+ This is the syntax:
```go
func main() {
    // Create a JSON string
    json := `{"field1":"Hello","field2":42}`

    // Create an empty instance of the Protobuf message
    message := &pb.Message{}

    // Convert the JSON to Protobuf
    if err := protojson.Unmarshal([]byte(json), message); err != nil {
        log.Fatalf("Failed to unmarshal message: %v", err)
    }

    // Print the Protobuf message
    fmt.Println(message)
}
```
