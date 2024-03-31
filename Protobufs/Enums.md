# Enums
+ Enums are used to define a set of named integer constants.
+ In Protobufs, enum tags should always start with a `0`.
+ They can be defined using the `enum` keyword like the following:
```proto
enum EyeColor {
  GREEN = 0;
  BLUE = 1;
  RED = 2;
}

message Account {
    string name = 1;
    EyeColor type = 2;
}
```
