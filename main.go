package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"

	"github.com/allthingsexperimental/gametime/ability"
	"github.com/allthingsexperimental/gametime/view/layout"
)

type Character struct {
	Name      string
	Abilities []*ability.Ability
}

func main() {
	a := app.New()
	w := a.NewWindow("Game Time!")

	text1 := widget.NewLabel("topleft")
	text2 := widget.NewLabel("Middle Label")
	text3 := widget.NewLabel("bottomright")

	w.SetContent(fyne.NewContainerWithLayout(layout.NewHomeLayout(), text1, text2, text3))
	w.ShowAndRun()
	/*

		char := Character{}
		fmt.Println("Hello, stranger!")
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
	*/
}
