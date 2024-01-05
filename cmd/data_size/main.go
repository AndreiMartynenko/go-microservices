package main

import (
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"

	desc "gitlab.ozon.dev/go/classroom-5/week-2/lecture-1/user_v1"
)

func main() {
	session := &desc.UserInfo{
		Id:      1,
		Name:    "John",
		IsHuman: true,
	}

	dataJson, _ := json.Marshal(session)
	fmt.Printf("\n\ndataJson len %d byte \n%v\n", len(dataJson), dataJson)

	dataPb, _ := proto.Marshal(session)
	fmt.Printf("dataPb len %d byte \n%v\n", len(dataPb), dataPb)

}
