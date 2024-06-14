package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

func (v *View) showFinishedUI() {
	v.idx.Clear()

	textView := tview.NewTextView().SetText(fmt.Sprintf("Word per minute %d", v.Game.Speed))

	v.idx.SetTitle("Total time taken" + " " + v.Game.TotalTime).SetBorder(true)

	v.idx.AddItem(textView, 0, 1, false)
}
