package ui

import (
	"github.com/rivo/tview"
	"typrfr/cmd/processor"
)

type UI struct {
	app  *tview.Application
	game *processor.Game
}

func (ui *UI) showScreen(state processor.State) {
	switch state {
	case processor.NOT_STARTED:
		ui.showGameStartUI()
	case processor.JOIN_ROOM:
		ui.showJoinRoomUI()
	case processor.WAITING_ROOM:
		ui.showWaitingRoomUI()
	case processor.IN_PROGRESS:
		ui.showInprogressUI()
	}
}

func Init() *UI {
	return &UI{
		app: tview.NewApplication(),
	}
}

func (ui *UI) Run() {
	ui.showScreen(processor.NOT_STARTED)

	if err := ui.app.Run(); err != nil {
		panic(err)
	}
}
