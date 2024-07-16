package main

import (
	"fmt"
	"lucky-number.alexedwards.net/internal/random"
	"github.com/fatih/color"
)

func main() {
	fmt.Printf("Your lucky number is %d!\n", random.Number())
	color.Cyan("DONE")
}
