# What Are Tags?
+ These are used to define the order of fields in the serialized message.
+ For example, in the following message, the tag for the field `name` is `1` and the tag for the field `id` is `2`.
```protobuf
message Person {
  string name = 1;
  int32 id = 2;
}
```

# Tag Rules
+ The smallest a tag can be is `1`.
+ The largest a tag can be is `536,870,911`.
+ The tags `19000 - 19999` are reserved tags you cannot use.

# Best Practices
+ It is better to use tags in the range `1 - 15` as they take 1 byte to encode.
+ The smaller the tag, the smaller the encoded message, and the faster the encoding and decoding.
+ Assign the most frequently occurring fields the smallest tags.
