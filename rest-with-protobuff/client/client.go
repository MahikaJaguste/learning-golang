package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	"restwithprotobuff.com/proto/echo"
)

func makeRequest(request *echo.EchoRequest) *echo.EchoResponse {

	req, err := proto.Marshal(request)
	if err != nil {
		log.Fatalf("Unable to marshal request : %v", err)
	}

	resp, err := http.Post("http://0.0.0.0:8080/echo", "application/x-binary", bytes.NewReader(req))
	if err != nil {
		log.Fatalf("Unable to read from the server : %v", err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Unable to read bytes from request : %v", err)
	}

	respObj := &echo.EchoResponse{}
	proto.Unmarshal(respBytes, respObj)
	return respObj

}

func main() {

	request := &echo.EchoRequest{Name: "Sushil"}
	resp := makeRequest(request)
	fmt.Printf("Response from API is : %v\n", resp.GetMessage())
}