package ui

import (
	"typrfr/cmd/game"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (v *View) showJoinRoomUI() {
	v.idx.Clear()
	input := tview.NewInputField().SetLabel("Enter the room code")

	input.SetDoneFunc(func(key tcell.Key) {
		v.Game = game.JoinRoom(input.GetText())
		if v.Game != nil {

			v.initWaitingRoom()
			v.ShowScreen(v.Game.State)
		} else {
		}
	})
	v.idx.AddItem(input, 0, 2, true)

	v.App.SetFocus(input)
}
