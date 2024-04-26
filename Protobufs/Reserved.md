# The `reserved` Keyword
+ Lets say that you have deployed a web application.
+ It has gained popularity, and a lot of people depend on this web application for daily tasks.
+ You want to make changes which include removing a tag from the Protobuf message.
+ Lets say we have the following message:
```protobuf
message MyMessage {
    string name = 1;
    string email = 2;
    string phone = 3;
}
```

+ If we wanted to remove the `email` field and just deleted it, and changed the `phone` tag to `3`, there will be problems.
+ This is because Protobufs use the tag number to deserialize the data, and if we change the tag number, the deserialization will fail.
+ To prevent this, we can use the `reserved` keyword:
```protobuf
message MyMessage {
    string name = 1;
    reserved 2;
    string phone = 3;
}
```
+ This will prevent the `email` field from being used, and the tag `2` will be reserved.
