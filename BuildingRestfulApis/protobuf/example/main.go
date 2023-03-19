package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
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

	pp := &pb.Person{}
	body, _ := proto.Marshal(p)
	_ = proto.Unmarshal(body, pp)
	fmt.Println("Original struct loaded from proto file: ", p, "\n")
	fmt.Println("Marsahlled proto data: ", body, "\n")
	fmt.Println("Unmarshaled struct: ", pp)
}
