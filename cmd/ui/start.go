package ui

import (
	"typrfr/cmd/processor"

	"github.com/rivo/tview"
)

func (ui *UI) showGameStartUI() {
	dropdown := tview.NewDropDown().SetLabel("Select an option (hit Enter)").SetOptions([]string{"Create room", "Join room", "Offline mode"}, func(text string, index int) {
		switch index {
		case 0:
			game := processor.CreateRoom()
			ui.game = game
			ui.showScreen(ui.game.State)
		case 1:
			ui.showScreen(processor.JOIN_ROOM)
		case 2:
			game := processor.NewLocalGame()
			ui.game = game
			ui.game.StartGame()
			ui.showScreen(ui.game.State)

		}
	})

	ui.app.SetRoot(dropdown, true)
}
