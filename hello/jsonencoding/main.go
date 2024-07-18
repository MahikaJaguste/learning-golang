// nested json
// struct tags
package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	Name string
	Age int
	isStudent bool
}

func main() {
	p := Person{Name: "Max", Age: 5, isStudent: true}

	b, _ := json.Marshal(p)
	fmt.Println(b)

	var pDecoded Person
	_ = json.Unmarshal(b, &pDecoded) // sets isStudent to false by default, as Marshal and Unmarshal do not affect unexported fields
	fmt.Println(pDecoded) 

	m := map[string]int{
		"Sam": 5,
		"Abhi": 10,
	}

	b2, _ := json.Marshal(m)
	fmt.Println(b2)

	var mDecoded map[string]int
	_ = json.Unmarshal(b2, &mDecoded)
	fmt.Println(mDecoded)

	b = []byte(`{"Name":"Bob","Food":"Pickle"}`)
	json.Unmarshal(b, &pDecoded)
	fmt.Println(pDecoded) // age still comes 5 as that field is not affected, only the fields in the encoded data present in the struct are overwritten

	var pDecodedClean Person
	json.Unmarshal(b, &pDecodedClean)
	fmt.Println(pDecodedClean) // here age takes default value 0



}