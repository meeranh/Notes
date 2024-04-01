package main

import (
	"fmt"

	catPb "github.com/meeranh/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func toJSON(pb proto.Message) string {
	json, err := protojson.Marshal(pb)
	checkErr(err)

	return string(json)
}

func fromJSON(json string, pb proto.Message) {
	err := protojson.Unmarshal([]byte(json), pb)
	checkErr(err)
}

func main() {
	myCat := &catPb.Cats{
		Id:   1,
		Name: "Poos",
	}

	myJson := toJSON(myCat)
	fmt.Println(myJson)

	myEmptyCat := &catPb.Cats{}
	fromJSON(myJson, myEmptyCat)

	fmt.Println(myEmptyCat)
}
