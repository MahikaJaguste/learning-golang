package main

import (
	"fmt"
	"strconv"
)

// ---

type Book struct {
    Title  string
    Author string
}

func (b Book) String() string {
	ogTitle := b.Title
	b.Title = "Hello"
    return fmt.Sprintf("Book: %s - %s", ogTitle, b.Author)
}

type Count int

func (c *Count) String() string {
	ogC := *c
	*c = Count(100)
	return strconv.Itoa(int(ogC))
}

func WriteLog(s fmt.Stringer) {
	fmt.Println(s.String())
}

// ---

func Example1() {
	b1 := Book{Title:"Atomic Habits", Author:"Dan V."} 
	c1 := Count(5)

	WriteLog(&b1)
	WriteLog(&c1)

	fmt.Println(b1) // won't change because the methods will deference b1 and create a copy
	fmt.Println(c1)
}