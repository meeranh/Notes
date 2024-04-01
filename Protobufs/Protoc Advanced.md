# Decoding Raw Protobuf Binary Data
+ The `protoc` compiler has the `--decode_raw` CLI flag.
+ It will take in binary data from STDIN and convert it to a somewhat readable format.
+ Imagine you were doing a reconnaissance on a gRPC application and you wanted to see what the raw binary data looked like.
+ You can use the `--decode_raw` flag to decode the binary data.
+ Since Protobufs are not like JSON, the output will only have the tag number and the value.
```bash
# Input
cat raw_data.bin | protoc --decode_raw
```
```bash
# Output
1: 1
2: "Hello, World!"
3: 123
```

# Decoding Binary Data With `.proto` File
+ If you have both the binary and the `.proto` file for the message, you can use the `--decode` flag.
+ It will use the binary data and generate a human-readable output with the keys and values.
+ There are three CLI arguments for `--decode`:
    1. The message type/name to decode.
    2. The `.proto` file.
    3. The binary data.
```bash
# Input
cat raw_data.bin | protoc --decode=MyMessage my_message.proto
```
```bash
# Output
name: "Hello, World!"
phone: 123
```

# Encoding Data With `.proto` File
+ If you have the `.proto` file and you want to encode data, you can use the `--encode` flag.
+ It will take in human-readable data and convert it to binary data.
+ There are three CLI arguments for `--encode`:
    1. The message type/name to encode.
    2. The `.proto` file.
    3. The human-readable data.
```bash
# Input
echo "name: 'Hello, World!' phone: 123" | protoc --encode=MyMessage my_message.proto
```
```bash
# Output
<binary data>
```
