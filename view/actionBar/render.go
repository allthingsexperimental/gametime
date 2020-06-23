package actionBar

import (
	"path"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func Render(props map[string]interface{}) fyne.CanvasObject {
	gear, _ := fyne.LoadResourceFromPath(path.Join("icons", "settings.jpg"))
	printer, _ := fyne.LoadResourceFromPath(path.Join("icons", "print.png"))
	save, _ := fyne.LoadResourceFromPath(path.Join("icons", "save.png"))
	saveaspdf, _ := fyne.LoadResourceFromPath(path.Join("icons", "topdf.jpeg"))
	leave, _ := fyne.LoadResourceFromPath(path.Join("icons", "exit.png"))
	randomize, _ := fyne.LoadResourceFromPath(path.Join("icons", "randomizer.jpeg"))
	loading, _ := fyne.LoadResourceFromPath(path.Join("icons", "load"))
	t := widget.NewToolbar(
		widget.NewToolbarAction(gear, func() {}),
		widget.NewToolbarAction(printer, func() {}),
		widget.NewToolbarAction(save, func() {}),
		widget.NewToolbarAction(saveaspdf, func() {}),
		widget.NewToolbarAction(leave, func() {}),
		widget.NewToolbarAction(randomize, func() {}),
		widget.NewToolbarAction(loading, func() {}),
	)
	return t
}
