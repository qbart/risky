package main

import (
	"fmt"

	"github.com/SoftKiwiGames/risky/risky"
)

type Animal struct {
	Name string `json:"name"`
}

func ExampleRisks() {
	animal := risky.JSON[Animal]([]byte(`{"name": "kiwi"}`))
	fmt.Println("animal =", animal.Name)

	animal = risky.JSON[Animal]([]byte(`abc`))
	fmt.Println("empty =", animal.Name == "")

	fmt.Println("int =", risky.Int("123"))
	fmt.Println("zero =", risky.Int("abc"))

	// Output:
	// animal = kiwi
	// empty = true
	// int = 123
	// zero = 0
}
