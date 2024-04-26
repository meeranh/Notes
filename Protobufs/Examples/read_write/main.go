package main

import (
	"fmt"
	"os"

	chatPb "github.com/meeranh/protobuf/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func writeToFile(file string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	checkErr(err)

	os.WriteFile(file, out, 0644)
}

func readFromFile(file string, pb proto.Message) {
	bytes, readErr := os.ReadFile(file)
	checkErr(readErr)

	decodeErr := proto.Unmarshal(bytes, pb)
	checkErr(decodeErr)
}

func main() {
	myChat := &chatPb.Chat{
		Id: 1,
		Msg: "Hello There!",
	}

	writeToFile("protoOutput.txt", myChat)

	newPb := &emptypb.Empty{}
	readFromFile("protoOutput.txt", newPb)

	fmt.Println(newPb)
}
