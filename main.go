package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"

	"github.com/allthingsexperimental/gametime/view"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("Game Time!")
	w.SetContent(view.Render())
	w.ShowAndRun()
}
