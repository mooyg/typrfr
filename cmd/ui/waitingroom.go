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

	layout := ui.view.AddItem(userList, 0, 2, false)

	startButton := tview.NewButton("Start game").SetSelectedFunc(func() {
		ui.game.SendStartGameCommand(fmt.Sprintf("%d", ui.game.Room.Id))
	})

	for i, user := range ui.game.Room.Users {
		if ui.game.Room.Leader == user.Id {
			userList.AddItem("Leader ID", fmt.Sprintf("%d", user.Id), rune(97+i), nil)
		} else {
			userList.AddItem("User ID", fmt.Sprintf("%d", user.Id), rune(97+i), nil)
		}
	}

	if ui.game.Room.MyId == ui.game.Room.Leader {
		layout.AddItem(startButton, 0, 2, true)
	}

	layout.SetBorder(true).SetTitle(fmt.Sprintf("Room code %d", ui.game.Room.Id))

	ui.app.SetRoot(layout, true)

}
