# Numbers
### Integers
+ `int32` - signed 32-bit integer
+ `int64` - signed 64-bit integer
+ `sint32` - signed 32-bit integer with variable-length encoding
+ `sint64` - signed 64-bit integer with variable-length encoding
+ `uint32` - unsigned 32-bit integer with variable-length encoding
+ `uint64` - unsigned 64-bit integer with variable-length encoding
+ `fixed32` - unsigned 32-bit integer with fixed-length encoding
+ `fixed64` - unsigned 64-bit integer with fixed-length encoding
+ `sfixed32` - signed 32-bit integer with fixed-length encoding
+ `sfixed64` - signed 64-bit integer with fixed-length encoding

### Floating Point Numbers
+ `float` - 32-bit floating point number
+ `double` - 64-bit floating point number

# Other
+ `bool` - boolean value
+ `string` - UTF-8 encoded or 7-bit ASCII text
+ `bytes` - sequence of bytes

# Example
+ We want to create a message with the following properties:
    1. *Age*
    2. *Account name*
    3. *Profile picture*
    4. *Profile verified*
    5. *Height*

```protobuf
syntax = "proto3";

message Account {
    uint32 id = 1;
    string name = 2;
    bytes thumbnail = 3;
    bool is_verified = 4;
    float height = 5;
}
```
