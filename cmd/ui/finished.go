package ui

import "github.com/rivo/tview"

func (ui *UI) showFinishedUI() {
	thanks := tview.NewBox().SetTitle("Thanks for participating").SetBorder(true)
	ui.app.SetRoot(thanks, true)
}
