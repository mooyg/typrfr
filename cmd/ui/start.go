package ui

import (
	"typrfr/cmd/processor"

	"github.com/rivo/tview"
)

func (ui *UI) showGameStartUI() {

	button := tview.NewButton("Hit enter to start typing").SetSelectedFunc(func() {
		ui.showScreen(processor.IN_PROGRESS)
	})

	ui.app.SetRoot(button, true)
}
