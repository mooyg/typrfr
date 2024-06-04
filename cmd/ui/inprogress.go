package ui

import (
	"typrfr/cmd/processor"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (ui *UI) showInprogressUI() {
	status := tview.NewTextView().SetText("TYPRFR")
	input := tview.NewTextArea().SetPlaceholder("Start typing...")
	nav := tview.NewFlex().AddItem(tview.NewBox(), 0, 1, false).AddItem(status, 0, 1, false)

	text := tview.NewTextView().SetText(ui.game.Sentence)

	content := tview.NewFlex().AddItem(tview.NewFlex().SetDirection(tview.FlexRow).AddItem(text, 0, 1, false).AddItem(input, 0, 3, true), 0, 1, true)
	layout := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(nav, 0, 1, false).AddItem(content, 0, 1, true)

	ui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if ui.game.State == processor.IN_PROGRESS {
			ui.game.ProcessTyping(event)
			if ui.game.HasFinished() == processor.FINISHED {
				ui.showFinishedUI()
			}

		}
		return event
	})

	layout.SetBorder(true)
	ui.app.SetRoot(layout, true)
}
