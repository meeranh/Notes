# Defining Maps
+ In Protobufs, you can specify a map using the following syntax:
```protobuf
map<key_type, value_type> map_field = field_number;
```
+ This is a full example of a protobuf with a map:
```protobuf
syntax = "proto3";

message Person {
  string name = 1;
  int32 id = 2;
  string email = 3;

  map<string, string> phone_numbers = 4;
}
```

+ This is how you use it in Golang:
```go
func main() {
    p := &Person{
        Name: "Meeran Hassan",
        Id: 1,
        Email: "mydumbmail@dumbmail.com",
        PhoneNumbers: map[string]string{
            "home": "1234567890",
            "work": "0987654321",
        },
    }

    fmt.Println(p)
}
```
