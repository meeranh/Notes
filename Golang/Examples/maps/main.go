package main

import "fmt"

type Human struct {
	name string
	age int
}

func main() {
	// This is a static map (not intended to add/remove elements later)
	myLanguages := map[string]string {
		"Java": "OOP",
		"Python": "OOP",
		"Go": "OOP",
	}

	// This is a dynamic map (intends to add/remove values later)
	myFriends := make(map[string]string)
	myFriends["John"] = "Java"
	myFriends["Ben"] = "Python"

	fmt.Println(myLanguages)
	fmt.Println(myFriends)

	// Deleting map values
	delete(myFriends, "John")
	fmt.Println(myFriends)

	// Test if a key is present
	value, ok := myFriends["John"]

	// If the key exists, ok will be true, else false.
	fmt.Println(ok) // false
	fmt.Println(value) // ""

	// Maps with a type
	favouritePeopleByOrder := map[int]Human {
		1: Human{"Mom", 50},
		2: Human{"Jithmi", 22},
		3: Human{"Gayindu", 21},
		4: Human{"Shamlan", 21},
	}

	fmt.Println(favouritePeopleByOrder)
}
