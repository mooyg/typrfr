package processor

import (
	"log/slog"
	"strings"

	"github.com/gdamore/tcell/v2"
)

const demoText = "Hello world"

type State int
type Game struct {
	Sentence    string
	State       chan State
	CurrentWord Word
	Words       []Word
}

const (
	NOT_STARTED State = iota
	IN_PROGRESS
	FINISHED
)

func NewGame() *Game {
	stateChannel := make(chan State)

	slog.Info("new game...")

	return &Game{
		Sentence: demoText,
		State:    stateChannel,
	}
}

func (g *Game) HasFinished() chan<- State {
	return g.State
}

func (g *Game) StartGame() *Game {
	g.State <- IN_PROGRESS
	return g
}

func (g *Game) EndGame() *Game {
	g.State <- FINISHED
	return g
}

func (g *Game) ConstructWords(s string) []Word {
	words := strings.Split(s, " ")
	constructedWords := make([]Word, 0, len(words))

	for _, v := range words {
		constructedWords = append(constructedWords, Word{
			ExpectedText: v,
			ProvidedText: []byte(""),
			Letters:      strings.Split(v, ""),
		})
	}

	g.Words = constructedWords

	return g.Words
}

// TODO: Use a better Data Structure maybe
type Word struct {
	ExpectedText string
	ProvidedText []byte
	Letters      []string
}

func ProcessTyping(event *tcell.EventKey, words []Word, currIndex int) {
	words[currIndex].ProvidedText = append(words[currIndex].ProvidedText, []byte(event.Name())...)
}
