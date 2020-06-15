package actionBar

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
)

func Render(props map[string]interface{}) fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
	)
}
