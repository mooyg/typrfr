package ui

import (
	"typrfr/cmd/processor"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (ui *UI) showJoinRoomUI() {

	input := tview.NewInputField().SetLabel("Enter the room code")

	layout := tview.NewFlex().AddItem(input, 0, 2, true)

	input.SetDoneFunc(func(key tcell.Key) {
		game := processor.JoinRoom(input.GetText())
		if game != nil {
			ui.game = game
			ui.showScreen(game.State)
		} else {
			ui.showErrorUI()
		}

	})

	ui.app.SetRoot(layout, true)
}
