package ui

import (
	"fmt"
	"github.com/rivo/tview"
)

func (ui *UI) showWaitingRoomUI() {
	ui.updateWaitingRoom()
	go ui.ListenChanges()
}
func (ui *UI) updateWaitingRoom() {
	ui.view.Clear()

	userList := tview.NewList()

	textView := tview.NewTextView()

	layout := ui.view.AddItem(userList, 0, 2, false).AddItem(textView, 0, 1, false)

	for _, v := range ui.game.Room.Connections {
		userList.AddItem("User Id", fmt.Sprintf("%d", v.Id), 'a', nil)
	}

	layout.SetBorder(true).SetTitle(fmt.Sprintf("Room code %d", ui.game.Room.Id))

	ui.app.SetRoot(layout, true)

}
