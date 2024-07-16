package main

import (
	"fmt"
	"log"
	"rsc.io/quote/v3"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	fmt.Println(quote.HelloV3())
	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
