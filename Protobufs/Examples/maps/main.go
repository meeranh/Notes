package main

import (
	"fmt"

	friendPb "github.com/meeranh/protobuf/proto"
)

func main() {

	myFriend := &friendPb.Friends{
		Friends: map[int32]*friendPb.NameWrapper{
			1: {Name: "Jithmi"},
			2: {Name: "Gayindu"},
			3: {Name: "Shamlan"},
		},
	}

	fmt.Println(myFriend)
}
