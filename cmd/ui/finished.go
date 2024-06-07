package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

func (ui *UI) showFinishedUI() {
	if ui.game.IsLocal == false {
		finished := tview.NewFlex()
		textView := tview.NewTextView().SetText(fmt.Sprintf("Word per minute %d", ui.game.Wpm))

		finished.AddItem(textView, 0, 1, false)
		finished.SetTitle("Total time taken" + " " + ui.game.TotalTime).SetBorder(true)

		ui.app.SetRoot(finished, true)

	} else {
		finished := tview.NewFlex()
		textView := tview.NewTextView().SetText(fmt.Sprintf("Word per minute %d", ui.game.Wpm))

		finished.AddItem(textView, 0, 1, false)
		finished.SetTitle("Total time taken" + " " + ui.game.TotalTime).SetBorder(true)

		ui.app.SetRoot(finished, true)

	}
}
