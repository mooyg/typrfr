package processor

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
	"typrfr/cmd/tcpclient"
	"typrfr/pkg/tcp"
	"typrfr/pkg/utils"

	"github.com/gdamore/tcell/v2"
)

type State int
type User struct {
	Conn *tcp.Connection
	Id   int
	Wpm  int
}

type Room struct {
	Id      int
	Users   []User
	Text    string
	Leader  int
	Started bool
	MyId    int
}

type Game struct {
	Sentence     string
	State        State
	Index        int
	Chars        []string
	timeStarted  time.Time
	timeEnded    time.Time
	TotalTime    string
	Room         Room
	Conn         *tcpclient.TCPClient
	Wpm          int
	IsLocal      bool
	RenderedText string
}

const (
	NOT_STARTED State = iota
	JOIN_ROOM         // Skipped if created a room
	WAITING_ROOM
	IN_PROGRESS
	FINISHED
)

type Data struct {
	Id   int
	Para string
}

func NewLocalGame() *Game {
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
		Sentence:     text,
		State:        NOT_STARTED,
		Index:        0,
		Chars:        strings.Split(text, ""),
		IsLocal:      true,
		RenderedText: text,
	}
}

func CreateRoom() *Game {

	conn := tcpclient.New()

	s := []byte("some\n")

	input := append([]byte{tcp.CREATE_ROOM}, s...)

	conn.Write(input)

	data, err := conn.Read()

	if err != nil {
		slog.Error("error while reading data from the client")
		os.Exit(2)
	}

	room := utils.Unmarshal[tcp.TCPCommand[Room]](data)

	return &Game{
		Sentence:     room.Data.Text,
		State:        WAITING_ROOM,
		Index:        0,
		Chars:        strings.Split(room.Data.Text, ""),
		Room:         room.Data,
		Conn:         conn,
		IsLocal:      false,
		RenderedText: room.Data.Text,
	}
}

func JoinRoom(id string) *Game {

	conn := tcpclient.New()

	s := append([]byte(id), '\n')
	input := append([]byte{tcp.JOIN_ROOM}, s...)

	conn.Write(input)

	data, err := conn.Read()

	if err != nil {
		slog.Error("error occured while joining a room")
		os.Exit(2)
	}

	room := utils.Unmarshal[tcp.TCPCommand[Room]](data)

	if room.Command == tcp.ERROR {
		return nil
	}

	game := &Game{
		Sentence:     room.Data.Text,
		State:        WAITING_ROOM,
		Index:        0,
		Chars:        strings.Split(room.Data.Text, ""),
		Room:         room.Data,
		Conn:         conn,
		IsLocal:      false,
		RenderedText: room.Data.Text,
	}

	return game
}

func (g *Game) HasFinished() State {
	return g.State
}

func (g *Game) StartGame() *Game {
	g.State = IN_PROGRESS
	g.timeStarted = time.Now()
	return g
}
func (g *Game) SendStartGameCommand(roomId string) *Game {
	s := append([]byte(roomId), '\n')
	input := append([]byte{tcp.START_GAME}, s...)
	g.Conn.Write(input)

	return g
}

func (g *Game) EndGame() *Game {
	g.State = FINISHED
	return g
}

func (g *Game) ProcessTyping(event *tcell.EventKey) {

	unwrapped := g.Unwrap(g.Chars[g.Index])

	if string(event.Rune()) == unwrapped {
		g.Chars[g.Index] = unwrapped
		g.RenderedText = strings.Join(g.Chars, "")
		g.Index = g.Index + 1
	} else {
		g.Chars[g.Index] = g.WrapColor(g.Chars[g.Index])
		g.RenderedText = strings.Join(g.Chars, "")
	}

	if g.Index == len(g.Chars) {
		g.timeEnded = time.Now()

		totalTime := g.timeEnded.Sub(g.timeStarted)
		g.TotalTime = totalTime.String()
		totalSeconds := totalTime.Seconds()
		totalMinutes := totalSeconds / 60

		g.State = FINISHED

		wordCount := len(strings.Split(g.Sentence, " "))

		g.Wpm = int(math.Round(float64(wordCount) / totalMinutes))

		return
	}
}
func (g *Game) WrapColor(c string) string {

	re := regexp.MustCompile(`\[.*?\]`)

	if len(re.FindAllString(c, -1)) > 0 {
		return c
	}

	out := fmt.Sprintf("%s%s%s", "[#ff0000]", c, "[white]")

	return out
}
func (g *Game) Unwrap(c string) string {
	r := strings.Replace(c, "[#ff0000]", "", 1)
	s := strings.Replace(r, "[white]", "", 1)
	return s
}
