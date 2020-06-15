package chat

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/allthingsexperimental/gametime/dice"
)

var chat *chatGroup

func Render(props map[string]interface{}) fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewGridLayoutWithRows(2),
		renderChatGroup(),
		renderCommandBar(),
	)
}

func renderChatGroup() fyne.CanvasObject {
	if chat == nil {
		chat = &chatGroup{Group: *widget.NewGroupWithScroller("chat")}
	}
	return chat
}

type chatGroup struct {
	widget.Group
	items []fyne.CanvasObject
}

func (c *chatGroup) Append(i fyne.CanvasObject) {
	c.items = append(c.items, i)
	c.Group.Append(i)
}

func (c *chatGroup) MinSize() fyne.Size {
	s := c.Group.MinSize()
	if len(c.items) > 0 {
		ch := 0
		for _, i := range c.items {
			ch += i.MinSize().Height
		}
		s.Height = s.Height + ch
	}
	return s
}

func renderCommandBar() fyne.CanvasObject {
	e := &cmdEntry{}
	e.ExtendBaseWidget(e)
	e.SetPlaceHolder("I await your command...")
	return e
}

type cmdEntry struct {
	widget.Entry
	value string
}

func (e *cmdEntry) onEnter() {
	roll := dice.Roller(20)
	l := widget.NewLabel(fmt.Sprintf("You rolled a: %v", roll))
	chat.Append(l)
	chat.Resize(chat.MinSize())
}

func (e *cmdEntry) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyReturn, fyne.KeyEnter:
		e.onEnter()
	default:
		e.Entry.KeyDown(key)
	}
}
