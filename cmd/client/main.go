package main

import (
	"github.com/rivo/tview"
	"typrfr/cmd/processor"
	"typrfr/cmd/ui"
)

func main() {
	game := processor.NewGame()
	app := tview.NewApplication()

	ui.Render(game, app)
	// Pass the not started state initially
	game.State <- processor.NOT_STARTED

	for {
		msg := <-game.State
		switch msg {
		case processor.IN_PROGRESS:
			ui.Render(game, app)
			break
		default:
			break
		}
	}
}
