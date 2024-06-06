package ui

import (
	"log/slog"
	"typrfr/cmd/processor"
	"typrfr/pkg/tcp"
	"typrfr/pkg/utils"

	"github.com/rivo/tview"
)

type UI struct {
	app  *tview.Application
	game *processor.Game
	view *tview.Flex
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
		app:  tview.NewApplication(),
		view: tview.NewFlex(),
	}
}

func (ui *UI) Run() {
	ui.showScreen(processor.NOT_STARTED)
	if err := ui.app.Run(); err != nil {
		panic(err)
	}
}
func (ui *UI) ListenChanges() {
	for {
		data, err := ui.game.Conn.Read()
		if err != nil {
			slog.Error("some error occured while processing a message")
		}

		cmd := utils.Unmarshal[tcp.TCPCommand[any]](data)

		if cmd.Command == tcp.NEW_USER_JOINED {
			d := utils.Unmarshal[tcp.TCPCommand[processor.Room]](data)
			ui.game.Room = d.Data
			ui.updateWaitingRoom()
			ui.app.Draw()
		}
	}
}
