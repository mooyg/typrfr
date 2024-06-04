package ui

import (
	"typrfr/cmd/processor"

	"github.com/rivo/tview"
)

func (ui *UI) showGameStartUI() {
	button := tview.NewButton("Hit enter to start typing").SetSelectedFunc(func() {
		ui.game.StartGame()
		ui.showScreen(processor.IN_PROGRESS)
	})

	ui.app.SetRoot(button, true)
}
