package main

import (
	"fmt"

	enumPb "github.com/meeranh/protobuf/proto"
)

func getEnum() *enumPb.Enumeration {
	return &enumPb.Enumeration{
		EyeColor: 1,
	}
}

func main() {
	myEye := getEnum()
	fmt.Println(myEye)
}
