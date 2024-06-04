package processor

import (
	"encoding/json"
	"github.com/gdamore/tcell/v2"
	"log/slog"
	"math/rand"
	"os"
	"strings"
	"time"
)

type State int
type Game struct {
	Sentence    string
	State       State
	Index       int
	Chars       []string
	timeStarted time.Time
	timeEnded   time.Time
	TotalTime   string
}

const (
	NOT_STARTED State = iota
	IN_PROGRESS
	FINISHED
)

type Data struct {
	Id   int
	Para string
}

func NewLocalGame() *Game {
	slog.Info("new game...")
	content, err := os.ReadFile("para.json")

	if err != nil {
		slog.Error("error while opening the file", "err", err)
	}
	var payload []Data
	err = json.Unmarshal(content, &payload)

	if err != nil {
		slog.Error("error during unmarshal()", "err", err)
	}

	text := payload[rand.Intn(len(payload))].Para

	return &Game{
		Sentence: text,
		State:    NOT_STARTED,
		Index:    0,
		Chars:    strings.Split(text, ""),
	}
}

func (g *Game) HasFinished() State {
	return g.State
}

func (g *Game) StartGame() *Game {
	g.State = IN_PROGRESS
	g.timeStarted = time.Now()
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
		g.timeEnded = time.Now()

		g.TotalTime = g.timeEnded.Sub(g.timeStarted).String()
		g.State = FINISHED
		return
	}

}
