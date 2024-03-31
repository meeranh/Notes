# Nested Messages
+ You can define one or more messages inside a message.
+ However, it is not a good practice to define a message inside another message.
+ It is better to define the nested message outside the parent message.

```proto
message Person {
  string name = 1;
  int32 id = 2;
  bool is_student = 3;
  Address address = 4;

  message Address {
    string street = 1;
    string city = 2;
  }
}
```
