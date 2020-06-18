package actionBar

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func Render(props map[string]interface{}) fyne.CanvasObject {
	gear, _ := fyne.LoadResourceFromURLString("https://banner2.cleanpng.com/20180410/xkq/kisspng-computer-icons-download-gears-5accce884910e3.2300818315233716562993.jpg")

	t := widget.NewToolbar(
		widget.NewToolbarAction(gear, func() {}),
	)
	return t
}
