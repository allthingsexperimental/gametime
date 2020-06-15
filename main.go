package main

import (
	"fyne.io/fyne/app"

	"github.com/allthingsexperimental/gametime/view"
)

func main() {
	a := app.New()
	w := a.NewWindow("Game Time!")
	w.SetContent(view.Render())
	w.ShowAndRun()
}
