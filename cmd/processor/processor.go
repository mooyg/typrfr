package processor

import (
	"github.com/gdamore/tcell/v2"
	"log/slog"
	"strings"
)

const demoText = "Hello world"

type State int
type Game struct {
	Sentence    string
	State       State
	StringTyped string
	Index       int
	Chars       []string
}

const (
	NOT_STARTED State = iota
	IN_PROGRESS
	FINISHED
)

func NewGame() *Game {
	slog.Info("new game...")

	return &Game{
		Sentence: demoText,
		State:    NOT_STARTED,
		Index:    0,
		Chars:    strings.Split(demoText, ""),
	}
}

func (g *Game) HasFinished() State {
	return g.State
}

func (g *Game) StartGame() *Game {
	g.State = IN_PROGRESS
	return g
}

func (g *Game) EndGame() *Game {
	g.State = FINISHED
	return g
}

func (g *Game) ProcessTyping(event *tcell.EventKey) {

	if string(event.Rune()) == g.Chars[g.Index] {
		g.Index = g.Index + 1
	} else {
		// TODO: highlight error
		slog.Info("highlight error here.")
	}
	if g.Index == len(g.Chars) {
		g.State = FINISHED
		return
	}

}
