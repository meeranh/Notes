# Value Receiver Methods
+ These are similar to class methods in OOP.
+ Basically if you have a `struct`, you can assign it a method.
+ The following is the syntax to do so:
```go
type User struct {
	Name	string
	Age		int
	Woke	bool
}

func (u User) IsWoke () {
	if u.Woke {
		fmt.Println(u.Name, "is woke")
	} else {
		fmt.Println(u.Name, "is not woke")
	}
}
```
+ This will now be a property of the `User` struct. You can access it as follows:
```go
func main() {
	mySelf := User{
		"Meeran",
		21,
		false,
	}
	mySelf.IsWoke()
}
```

# Pointer Receiver Methods
+ The problem with normal value receiver methods are that their changes are not reflected outside the method.
+ If for example we create a setter as such, it will not work:
```go
// A value receiver method
func (u User) setValueName(newName string){
	u.Name = newName

func main() {
    mySelf := User{
        "Meeran
    }
    mySelf.setValueName("Not Meeran")
    fmt.Println(mySelf.Name) // Meeran
}
```
+ This is because the method creates a copy of the struct, and then makes changes to the copy.
+ If we want to make changes to the actual struct, we need to use pointers.
+ It will refer to the actual memory location of the struct:
```go
// A pointer receiver method
func (u *User) setPointerName(newName string){
    u.Name = newName
}

func main() {
    mySelf := User{
        "Meeran
    }
    mySelf.setPointerName("Not Meeran")
    fmt.Println(mySelf.Name) // Not Meeran
}
```
