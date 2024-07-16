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
	message, err := greetings.Hello("Max")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Max", "John", "Dev"}
	m, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range names {
		fmt.Printf("%v\n", m[v])
	}

	fmt.Println(m)
	
}
