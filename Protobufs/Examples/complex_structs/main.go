package main

import (
	"fmt"

	childPb "github.com/meeranh/protobuf/proto"
)

func createChild() *childPb.Infant {
	return &childPb.Infant{
		Id:       1,
		FullName: "Namal Rajapakshe",
		Metadata: &childPb.Metadata{
			IsMale: false,
			Dob:    "2002-08-12",
		},
		Parents: []*childPb.Parent{
			{
				Id:       1,
				FullName: "Mahinda Rajapakshe",
			},
			{
				Id:       2,
				FullName: "Shiranthi Rajapakshe",
			},
		},
	}
}

func main() {
	myChild := createChild()
	fmt.Println(myChild)
}
