package main

import (
	"fmt"

	foodPb "github.com/meeranh/protobuf/proto"
)

func addFood() *foodPb.Food {
	return &foodPb.Food{
		Id:     1,
		Name:   "Kottu",
		Rating: 10,
		Types:  []string{"Cheese Kottu", "Dolphin Kottu", "Chicken Kottu"},
	}
}

func main() {
	foodProto := addFood()

	fmt.Println(foodProto)
}
