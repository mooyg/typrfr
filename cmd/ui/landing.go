package ui

import (
	"typrfr/cmd/game"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (v *View) showLandingUI() {
	v.idx.Clear()

	dropdown := tview.NewDropDown().SetLabel("Select an option (hit Enter)").SetOptions([]string{"Create room", "Join room", "Offline mode"}, func(text string, index int) {
		switch index {
		case 0:
			v.Game = game.CreateRoom()
			v.initWaitingRoom()
			v.ShowScreen(v.Game.State)
		case 1:
			v.ShowScreen(game.JOIN_ROOM)
		case 2:
			v.Game = game.NewLocalGame()
			v.ShowScreen(v.Game.State)
		}
	})

	frame := tview.NewFrame(tview.NewFlex().AddItem(dropdown, 0, 1, true)).AddText("Welcome to typrfr", true, tview.AlignCenter, tcell.ColorBlue)

	v.idx.AddItem(frame, 0, 1, true)
}
