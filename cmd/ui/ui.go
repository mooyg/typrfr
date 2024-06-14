package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"log/slog"
	"os"
	"typrfr/cmd/game"
	"typrfr/pkg/logger"
	"typrfr/pkg/shared"
	"typrfr/pkg/utils"
)

type View struct {
	App  *tview.Application
	idx  *tview.Flex
	Game *game.Game
}

func Init() View {
	app := tview.NewApplication()
	frame := tview.NewFlex()

	frame.SetBorder(true).SetTitle("typrfr").SetBackgroundColor(tcell.ColorBlue)

	return View{
		App: app,
		idx: frame,
	}
}

func (v *View) Run() {
	v.ShowScreen(game.NOT_STARTED)
	if err := v.App.SetRoot(v.idx, true).Run(); err != nil {
		slog.Info("some error occured")
		os.Exit(2)
	}
}

func (v *View) ShowScreen(state game.GameState) {
	switch state {
	case game.NOT_STARTED:
		v.showLandingUI()
	case game.WAITING_ROOM:
		v.showWaitingRoomUI()
	case game.JOIN_ROOM:
		v.showJoinRoomUI()
	case game.IN_PROGRESS:
		v.showInProgressUI()
	case game.FINISHED:
		v.showFinishedUI()
	}
}

// Listen to changes like `NEW_USER_JOINED` etc
func (v *View) ListenChanges() {

	for {
		data, err := v.Game.ClientConn.Read()

		if err != nil {
			slog.Error("error occured while processing message from the server (ListenChanges)", "err", err)
			os.Exit(2)
		}
		cmd := utils.Unmarshal[shared.TCPCommand[any]](data)

		switch cmd.Command {
		case shared.NEW_USER_JOINED:
			newRoomData := utils.Unmarshal[shared.TCPCommand[shared.MultiplayerRoom]](data).Data

			v.Game.Room = &newRoomData

			logger.Log.Println("room updated as new user joined", v.Game.Room)

			v.ShowScreen(v.Game.State)

			v.App.ForceDraw()
		case shared.START_GAME:
			v.Game.StartGame()
			v.ShowScreen(v.Game.State)
			v.App.ForceDraw()
			logger.Log.Println("room started")
		}

	}

}
