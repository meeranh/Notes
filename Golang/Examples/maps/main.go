package main

import "fmt"

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
}
