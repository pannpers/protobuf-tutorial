package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	pb "github.com/pannpers/protobuf-tutorial/protofiles"
)

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "Jonh Doe",
		Email: "jonh.doe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "0120-0000", Type: pb.Person_MOBILE},
		},
	}

	p1 := &pb.Person{}
	body, _ := proto.Marshal(p)
	_ = proto.Unmarshal(body, p1)

	fmt.Printf("Original struct loaded from proto file: %v\n", p)
	fmt.Printf("Marshaled proto data: %v\n", body)
	fmt.Printf("Unmarshaled struct: %v\n", p1)
}
