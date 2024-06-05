package ui

import "github.com/rivo/tview"

func (ui *UI) showErrorUI() {
	box := tview.NewBox().SetTitle("No room exists").SetBorder(true)

	ui.app.SetRoot(box, true)
}
