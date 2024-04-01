# Comments

## Single Line Comments
+ Single line comments are created using `//` and can be placed anywhere in the file.

```protobuf
// This is a single line comment
message Person {
  // This is another single line comment
  string name = 1; // This is a comment at the end of the line
}
```

## Multi Line Comments
+ Multi line comments are created using `/*` and `*/` and can be placed anywhere in the file.

```protobuf
/*
This is a multi line comment
It can span multiple lines
*/
message Person {
  /*
  This is another multi line comment
  It can span multiple lines
  */
  string name = 1; /* This is a comment at the end of the line */
}
```
