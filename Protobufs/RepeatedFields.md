# What Are Repeated Fields?
+ Imagine you want to send an array/list of values in that field.
+ For example, you want to send a list of phone numbers of a person.
+ In this case, you can use repeated fields like the following:
```proto
message Person {
  string name = 1;
  int32 id = 2;
  repeated string phone_numbers = 3;
}
```

+ So when sending, we can send multiple phone numbers like the following:
```proto
{
  "name": "John Doe",
  "id": 1234,
  "phone_numbers": ["123-456-7890", "123-456-7891"]
}
```
