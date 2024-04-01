# Enums
+ Enums are used to define a set of named integer constants.
+ In Protobufs, enum tags should always start with a `0`.
+ Names of the enums should also be in uppercase.
+ The first enum should be prefixed with `ENUM_NAME_UNSPECIFIED` and should have a value of `0`.
+ They can be defined using the `enum` keyword like the following:
```protobuf
enum EyeColor {
  EYE_COLOR_UNSPECIFIED = 0;
  EYE_COLOR_GREEN = 1;
  EYE_COLOR_BLUE = 2;
  EYE_COLOR_RED = 3;
}

message Account {
    string name = 1;
    EyeColor type = 2;
}
```
