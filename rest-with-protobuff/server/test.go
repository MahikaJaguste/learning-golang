package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"restwithprotobuff.com/proto/echo"
)

func test() {
	req := &echo.EchoRequest{Name: "Sushil"}
	fmt.Println(req)
	data, err := proto.Marshal(req)
	if err != nil {
		log.Fatalf("Error while marshalling the object : %v", err)
	}
	fmt.Println(data)

	res := &echo.EchoRequest{}
	err = proto.Unmarshal(data, res)
	fmt.Println(res)
	if err != nil {
		log.Fatalf("Error while un-marshalling the object : %v", err)
	}
	fmt.Printf("Value from un-marshalled data is %v", res.GetName())

}