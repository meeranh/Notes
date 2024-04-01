package main

import (
	"fmt"

	enumPb "github.com/meeranh/protobuf/proto"
)

func getEnum() *enumPb.Enumeration {
	return &enumPb.Enumeration{
		EyeColor: enumPb.EyeColor_EYE_COLOR_RED,
	}
}

func main() {
	myEye := getEnum()
	fmt.Println(myEye)
}
