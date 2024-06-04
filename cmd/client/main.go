package main

import (
	"typrfr/cmd/processor"
	"typrfr/cmd/ui"
)

func main() {
	game := processor.NewGame()
	ui := ui.Init(game)
	ui.Run()
}
