package ui

import (
	"fmt"
	"github.com/rivo/tview"
)

func (v *View) initWaitingRoom() {
	go v.ListenChanges()
}
func (v *View) showWaitingRoomUI() {
	v.idx.Clear()

	banner := tview.NewTextView().SetText(`
 ___       __   ________  ___  _________  ___  ________   ________     
|\  \     |\  \|\   __  \|\  \|\___   ___\\  \|\   ___  \|\   ____\    
\ \  \    \ \  \ \  \|\  \ \  \|___ \  \_\ \  \ \  \\ \  \ \  \___|    
 \ \  \  __\ \  \ \   __  \ \  \   \ \  \ \ \  \ \  \\ \  \ \  \  ___  
  \ \  \|\__\_\  \ \  \ \  \ \  \   \ \  \ \ \  \ \  \\ \  \ \  \|\  \ 
   \ \____________\ \__\ \__\ \__\   \ \__\ \ \__\ \__\\ \__\ \_______\
    \|____________|\|__|\|__|\|__|    \|__|  \|__|\|__| \|__|\|_______|
	`)
	userList := tview.NewList()

	for i, user := range v.Game.Room.Users {
		if user.Id == v.Game.Room.Leader {
			userList.AddItem("Leader Id", fmt.Sprintf("%d", user.Id), rune(97+i), func() {})
		} else {
			userList.AddItem("User Id", fmt.Sprintf("%d", user.Id), rune(97+i), func() {})
		}
	}

	v.idx.SetTitle(fmt.Sprintf("Room code %d", v.Game.Room.Id))

	v.idx.SetDirection(tview.FlexRow).AddItem(banner, 0, 1, false).AddItem(userList, 0, 2, false)

}
