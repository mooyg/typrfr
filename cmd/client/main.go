package main

import (
	"typrfr/cmd/processor"
	"typrfr/cmd/ui"
)

func main() {
	game := processor.NewLocalGame()
	ui := ui.Init(game)
	ui.Run()
}
