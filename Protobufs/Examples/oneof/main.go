package main

import (
	"fmt"

	myPb "github.com/meeranh/protobuf/proto"
)

func main() {
	msgWithNum := &myPb.MyMessage{
		MyData: &myPb.MyMessage_MyNum{MyNum: 3},
	}

	msgWithString := &myPb.MyMessage{
		MyData: &myPb.MyMessage_MyMsg{MyMsg: "Cool"},
	}

	fmt.Println(msgWithNum)
	fmt.Println(msgWithString)
}
