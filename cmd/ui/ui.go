package ui

import (
	"typrfr/cmd/processor"

	"github.com/rivo/tview"
)

type UI struct {
	app  *tview.Application
	game *processor.Game
}

func (ui *UI) showScreen(state processor.State) {
	switch state {
	case processor.NOT_STARTED:
		ui.showGameStartUI()
	case processor.IN_PROGRESS:
		ui.showInprogressUI()
	}
}

func Init(game *processor.Game) *UI {
	return &UI{
		app:  tview.NewApplication(),
		game: game,
	}
}

func (ui *UI) Run() {
	ui.showScreen(processor.NOT_STARTED)
	if err := ui.app.Run(); err != nil {
		panic(err)
	}
}
