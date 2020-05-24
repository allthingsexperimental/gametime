package main

import (
	"fmt"

	"github.com/allthingsexperimental/gametime/ability"
)

type Character struct {
	Name      string
	Abilities []*ability.Ability
}

func main() {
	char := Character{}
	fmt.Println("Hello, world!")
	fmt.Print("What should I call you? ")
	fmt.Scanln(&char.Name)
	fmt.Printf("Nice to meet you, %s!\n", char.Name)
	for _, v := range ability.DefaultAbilities {
		a := ability.New(v)
		char.Abilities = append(char.Abilities, a)
	}
	fmt.Printf("Here is your sheet so far %s\n", char.Name)
	for _, v := range char.Abilities {
		fmt.Printf("%s: %d\n", v.Name, v.Base)
	}
}
