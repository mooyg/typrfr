package main

import (
	"fmt"
	"typrfr/cmd/processor"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const demoText = `Hello world`

func main() {
	app := tview.NewApplication()
	words := processor.ConstructWords(demoText)
	textArea := tview.NewTextArea().SetPlaceholder("Start Typing...")

	// Divide each letter into regions.
	textView := tview.NewTextView().SetText(demoText)

	currIndex := 0
	textArea.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		//TODO: Get the event here and highlight the text view
		processor.ProcessTyping(event, words, currIndex)
		if event.Rune() == 32 {
			fmt.Fprintf(textView, "Recieved  %s", string(words[currIndex].ProvidedText))
			currIndex++
		}

		return event
	})

	container := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(textView, 0, 1, false).AddItem(textArea, 0, 1, true)

	if err := app.SetRoot(container, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
