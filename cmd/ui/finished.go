package ui

import (
	"github.com/rivo/tview"
)

func (ui *UI) showFinishedUI() {
	thanks := tview.NewBox().SetTitle("Total time taken" + " " + ui.game.TotalTime).SetBorder(true)
	ui.app.SetRoot(thanks, true)
}
