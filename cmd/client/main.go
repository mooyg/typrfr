package main

import (
	"typrfr/cmd/processor"
	"typrfr/cmd/ui"
)

func main() {
	game := processor.CreateRoom()
	ui := ui.Init(game)
	ui.Run()
}
