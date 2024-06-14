package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (v *View) showInProgressUI() {
	v.idx.Clear()
	renderedText := tview.NewTextView().SetText(v.Game.RenderedText).SetRegions(true).SetDynamicColors(true).SetToggleHighlights(true)

	input := tview.NewTextArea().SetPlaceholder("Start typing...")

	typr := tview.NewTextView().SetText(`
___________                     _____________________ 
\__    ___/__.__._____________  \_   _____/\______   \
  |    | <   |  |\____ \_  __ \  |    __)   |       _/
  |    |  \___  ||  |_> >  | \/  |     \    |    |   \
  |____|  / ____||   __/|__|     \___  /    |____|_  /
          \/     |__|                \/            \/ 
`)
	frame := tview.NewFlex().AddItem(tview.NewBox(), 0, 1, false).AddItem(typr, 0, 2, false)

	container := tview.NewFlex().SetDirection(tview.FlexRow)

	container.AddItem(renderedText, 0, 2, false).AddItem(input, 0, 2, true)

	v.idx.SetDirection(tview.FlexRow).AddItem(frame, 0, 1, false).AddItem(container, 0, 2, true)

	// Set focus manually because for some weird reason container doesn't get the focus
	v.App.SetFocus(input)

	input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return event
	})

}
