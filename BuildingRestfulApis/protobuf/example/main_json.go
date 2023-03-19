package main

import (
	"encoding/json"
	"fmt"
)
import pb "protobuf/protofiles"

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "RicardoJ",
		Email: "rr@mail.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "5550555", Type: pb.Person_HOME},
		},
	}

	body, _ := json.Marshal(p)
	fmt.Println(string(body))
}
