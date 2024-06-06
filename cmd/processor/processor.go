package processor

import (
	"encoding/json"
	"github.com/gdamore/tcell/v2"
	"log/slog"
	"math/rand"
	"os"
	"strings"
	"time"
	"typrfr/cmd/tcpclient"
	"typrfr/pkg/tcp"
	"typrfr/pkg/utils"
)

type State int

type Room struct {
	Id          int
	Text        string
	Connections []*tcp.Connection
	Leader      int
	Connection  tcp.Connection
	Started     bool
}

type Game struct {
	Sentence    string
	State       State
	Index       int
	Chars       []string
	timeStarted time.Time
	timeEnded   time.Time
	TotalTime   string
	Room        Room
	Conn        *tcpclient.TCPClient
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
		Sentence: text,
		State:    NOT_STARTED,
		Index:    0,
		Chars:    strings.Split(text, ""),
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
		Sentence: room.Data.Text,
		State:    WAITING_ROOM,
		Index:    0,
		Chars:    strings.Split(room.Data.Text, ""),
		Room:     room.Data,
		Conn:     conn,
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
		Sentence: room.Data.Text,
		State:    WAITING_ROOM,
		Index:    0,
		Chars:    strings.Split(room.Data.Text, ""),
		Room:     room.Data,
		Conn:     conn,
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
