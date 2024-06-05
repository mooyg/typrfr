package ui

import (
	"fmt"
	"github.com/rivo/tview"
)

func (ui *UI) showWaitingRoomUI() {
	userList := tview.NewList()

	layout := tview.NewFlex().AddItem(userList, 0, 2, false)

	for _, v := range ui.game.Room.Connections {
		userList.AddItem("User Id", fmt.Sprintf("%d", v.Id), 'a', nil)
	}

	layout.SetBorder(true).SetTitle(fmt.Sprintf("Room code %d", ui.game.Room.Id))

	ui.app.SetRoot(layout, true)
}
