package view

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"github.com/allthingsexperimental/gametime/state"
	"github.com/allthingsexperimental/gametime/view/actionBar"
	"github.com/allthingsexperimental/gametime/view/chat"
	"github.com/allthingsexperimental/gametime/view/dice"
)

func Render() fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewGridLayoutWithColumns(2),
		details(),
		actions(),
	)
}

func details() fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewGridWrapLayout(fyne.NewSize(150, 300)),
		dice.Render(state.Values),
	)
}

func actions() fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		chat.Render(state.Values),
		actionBar.Render(state.Values),
	)
}
