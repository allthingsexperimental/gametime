package dice

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func Render(props map[string]interface{}) fyne.CanvasObject {
	label := widget.NewLabel("")
	if roll, ok := props["roll"]; ok {
		label.Text = string(roll.(uint8))
	}
	return label
}
