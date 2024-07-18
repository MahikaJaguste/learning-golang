package main

import (
	"fmt"
	"reflect"
	"example/hello/structs/exportedfields"
)

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Maxa", 21})
	fmt.Println(person{name: "Maxo", age:22})
	fmt.Println(person{age: 23})
	fmt.Println(person{name: "Maxi"})
	fmt.Println(&person{name: "Maxu", age:24})
	fmt.Println(*newPerson("Maxe"))

	s := person{"Max", 25}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 26
	fmt.Println(sp.age)

	t := struct {
		name string
		age int
	} {
		name: "Mahi",
		age: 30,
	}
	fmt.Println(t, reflect.TypeOf(t))
	tConv := person(t)
	fmt.Println(tConv, reflect.TypeOf(tConv))

	importedStruct := exportedfields.Person{Name:"Sam"} // unknown field age in struct literal of type exportedfields.Person 
	fmt.Println(importedStruct, reflect.TypeOf(importedStruct)) // only exported fields are accessible in this package


}