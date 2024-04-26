package main

import "fmt"

func main() {
	var favouritePeople [] string = [] string {
		"Gayindu",
		"Shamlan",
		"Jithmi",
	}

	// Appending a value
	favouritePeople = append(favouritePeople, "ThisWillBePoppedLater")
	fmt.Println(favouritePeople)

	// Slices are like references to arrays
	yetAnotherSlice := favouritePeople[0:3]
	fmt.Println(yetAnotherSlice)

	// Popping a value
	favouritePeople = favouritePeople[0:len(favouritePeople)-1]
	fmt.Println(favouritePeople)
}
