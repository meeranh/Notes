package main

import "fmt"

type User struct {
	Name string
	Age  int
	Woke bool
}

// A normal method
func (u User) IsWoke() {
	if u.Woke {
		fmt.Println(u.Name, "is woke")
	} else {
		fmt.Println(u.Name, "is not woke")
	}
}

// A value receiver method
func (u User) setValueName(newName string) {
	u.Name = newName
}

// Method with a pointer receiver
func (u *User) SetPointerName(newName string) {
	u.Name = newName
}

// You could do this with just a pointer function also
func changeName(oldName *string, newName string) {
	*oldName = newName
}

func main() {
	mySelf := User{
		"Meeran",
		21,
		false,
	}
	mySelf.IsWoke()

	// This will not change the name of the user
	mySelf.setValueName("NotMeeran")
	fmt.Println(mySelf.Name) // Meeran

	// This will change the name of the user
	mySelf.SetPointerName("NotMeeran")
	fmt.Println(mySelf.Name) // NotMeeran

	// You could also do this with a pointer function
	changeName(&mySelf.Name, "KottuNaana")
	fmt.Println(mySelf.Name) // KottuNaana
}
