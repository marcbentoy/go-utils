package main

import (
	"fmt"
	"github.com/marcbentoy/go-utils/person"
)

func main() {
	fmt.Println("Implementing git subtree")

	p1 := sdkperson.NewPerson("vince")
	fmt.Printf("p1: %v\n", p1)
}
