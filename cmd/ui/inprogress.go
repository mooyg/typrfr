package ui

import (
	"github.com/rivo/tview"
)

func (ui *UI) showInprogressUI() {
	status := tview.NewTextView().SetText("TYPRFR")
	input := tview.NewTextArea().SetPlaceholder("Start typing...")
	// sentence := tview.NewTextView().SetText(ui.game.Sentence)
	nav := tview.NewFlex().AddItem(tview.NewBox(), 0, 1, false).AddItem(status, 0, 1, false)

	text := tview.NewTextView().SetText(ui.game.Sentence)

	content := tview.NewFlex().AddItem(tview.NewFlex().SetDirection(tview.FlexRow).AddItem(text, 0, 1, false).AddItem(input, 0, 3, true), 0, 1, true)

	layout := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(nav, 0, 1, false).AddItem(content, 0, 1, true)
	layout.SetBorder(true)
	ui.app.SetRoot(layout, true)
}
