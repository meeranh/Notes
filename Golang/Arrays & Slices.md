# Arrays
## Array Syntax
+ An array is a **fixed-size** sequence of elements of the same type.
+ The syntax for an array is `var name [size]type`. `size`.
```go
func main(){
    // Long syntax
    var x [5]int = [5]int{1, 2, 3, 4, 5}

    // Shorthand syntax
    y := [5]int{1, 2, 3, 4, 5}
}
```

## Slices / Slice Literals
+ This is an array without a fixed size. In other words, you do not need to specify the size of a slice.
+ You can create a slice with the same syntax as an array, but without the size.
```go
func main(){
    // Long syntax
    var x []int = []int{1, 2, 3, 4, 5}

    // Shorthand syntax
    y := []int{1, 2, 3, 4, 5}
}
```

## Appending Slices

+ You can append elements to a slice using the `append` function.
```go
func main(){
	var favouritePeople [] string = [] string {"Gayindu", "Shamlan", "Jithmi"}
	favouritePeople = append(favouritePeople, "ThisWillBePoppedLater")
	fmt.Println(favouritePeople)
}
```
+ The `append` function will take in a slice, and return a new appended slice. It does not modify an existing slice.
+ So, you need to reassign the new slice to the variable through `x = append(x, 6)` as just running `append(x, 6)` will not work.

## Popping Slices via Ranges

+ You will need to filter out the last element of a slice using a range like the following example
```go
func main(){
    var favouritePeople [] string = [] string {"Gayindu", "Shamlan", "Jithmi", "ThisWillBePoppedLater"}
    favouritePeople = append(favouritePeople[0:len(favouritePeople)-1])
    fmt.Println(favouritePeople)
}
```

## Slice Capacity
+ The `cap()` function is used to find a capacity of a slice.
+ The capacity of a slice is basically the length of the underlying array.
```go
func main() {
	mySliceLiteral := [] int {1, 2, 3, 4, 6, 7}
	firstSlice := mySliceLiteral[0:5]

	fmt.Println(cap(firstSlice)) // 6
	fmt.Println(len(firstSlice)) // 5
}
```

## Creating a Slice With `make`
+ Similar to a slice literal, you can create a slice with the `make` function, which will basically act as a dynamic array.
+ The `make` function takes in three parameters:
    1. The type of the slice: `[]int`
    2. The length of the slice: `5`
    3. The capacity of the slice: `10`
```go
func main() {
    // Capacity is optional
    mySlice := make([]int, 5)
    fmt.Println(mySlice)        // [0 0 0 0 0]
    fmt.Println(len(mySlice))   // 5
    fmt.Println(cap(mySlice))   // 5

    // With capacity
    capSlice := make([]int, 5, 10)
    fmt.Println(mySlice)        // [0 0 0 0 0]
    fmt.Println(len(mySlice))   // 5
    fmt.Println(cap(mySlice))   // 10

    // A nill slice
    nilSlice := make([]int, 0)
    append(nilSlice, 10)
    fmt.Println(nilSlice)       // [10]
}
```
