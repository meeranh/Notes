# Maps
+ These are similar to dictionaries in Python. They are key-value pairs.
+ There are two ways to create a map, you can either create a fixed size map or a dynamic map.

## A Fixed Sized Map via `new`
+ When creating a map like this, you will not be able to add or remove elements from the map. This is because the `new` keyword is used behind the scenes.
+ This is because the compiler will allocate a fixed amount of memory for this specific variable.
```go
func main() {
    varName := map[keyType]valueType{
        key1: value1,
        key2: value2,
        key3: value3,
    }
}
```

## A Dynamic Sized Map via `make`
+ One important thing with maps in Go is that when you create a map, it should be created with the `make` function as follows:
```
func main() {
    varName := make(map[keyType]valueType)
    varName[key1] = value1
    varName[key2] = value2
    varName[key3] = value3
}
```

## Mutating Maps
+ Changing, accessing, and deleting a value of a map's key can be done using the following syntax:
```go
varName[key1] = "NewValue"  // Changing
fmt. Println(varName[key1]) // Accessing
delete(varName, key1)       // Deleting
```

## Testing If A Key Is Present
+ We can use the following syntax to check if a key is present in a map:
```go
value, ok := varName[key]
```
+ If the key is present, `ok` will be `true`, otherwise, it will be `false`.
+ The `value` will be the value of the key if it is present, otherwise, it will be the zero value of the value type. For example:
```go
mymap := make(map[string]string {
    "key1": "value1",
    "key2": "value2",
})

func main() {
    value, ok := mymap["key3"]
    fmt.Println(ok)     // false
    fmt.Println(value)  // ""
}
```
