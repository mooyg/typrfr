package game

import (
	"encoding/json"
	"log/slog"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
	"typrfr/cmd/tcpclient"
	"typrfr/pkg/logger"
	"typrfr/pkg/shared"
	"typrfr/pkg/utils"

	"github.com/gdamore/tcell/v2"
)

type GameState int

const (
	NOT_STARTED  GameState = iota
	WAITING_ROOM           // skipped if game is local
	JOIN_ROOM              // Enter code for joining a room
	IN_PROGRESS
	FINISHED
)

type Game struct {
	State   GameState
	Text    string
	Speed   int
	isLocal bool
	Index   int
	Chars   []string
	// String which controls the rendering behaviour of the text (region colouring)
	RenderedText string
	Room         *shared.MultiplayerRoom
	Me           shared.User
	ClientConn   *tcpclient.TCPClient
	timeStarted  time.Time
	timeEnded    time.Time
	TotalTime    string
}

type FileData struct {
	Id   int
	Para string
}

func NewLocalGame() *Game {
	content, err := os.ReadFile("para.json")

	if err != nil {
		slog.Error("error while opening the file", "err", err)
	}
	var payload []FileData

	err = json.Unmarshal(content, &payload)

	if err != nil {
		slog.Error("error during unmarshal()", "err", err)
	}

	randomText := payload[rand.Intn(len(payload))].Para
	chars := strings.Split(randomText, "")

	return &Game{
		State:        IN_PROGRESS,
		Text:         randomText,
		Index:        0,
		Chars:        chars,
		isLocal:      true,
		Room:         nil,
		Speed:        0,
		RenderedText: randomText,
		ClientConn:   nil,
	}
}

func CreateRoom() *Game {
	conn := tcpclient.New()

	conn.Write([]byte{shared.REQUEST_USER_ID})
	data, err := conn.Read()

	if err != nil {
		slog.Error("error while reading request user id data from the client", "err", err)
		os.Exit(2)
	}
	// Get the user id from the server
	user := utils.Unmarshal[shared.TCPCommand[shared.User]](data).Data

	conn.Write([]byte{shared.CREATE_ROOM})

	d, e := conn.Read()
	if e != nil {
		slog.Error("error while reading create room data from the client", e)
		os.Exit(2)
	}

	roomData := utils.Unmarshal[shared.TCPCommand[shared.MultiplayerRoom]](d)

	return &Game{
		State:        WAITING_ROOM,
		isLocal:      false,
		Text:         roomData.Data.Text,
		Speed:        0,
		Index:        0,
		Chars:        strings.Split(roomData.Data.Text, ""),
		RenderedText: roomData.Data.Text,
		Room:         &roomData.Data,
		Me:           user,
		ClientConn:   conn,
	}
}

func JoinRoom(roomId string) *Game {
	conn := tcpclient.New()

	conn.Write([]byte{shared.REQUEST_USER_ID})
	data, err := conn.Read()

	if err != nil {
		slog.Error("error while reading request user id data from the client", "err", err)
		os.Exit(2)
	}
	// Get the user id from the server
	user := utils.Unmarshal[shared.TCPCommand[shared.User]](data).Data

	input := []byte{shared.JOIN_ROOM}

	conn.Write(append(input, roomId...))

	d, e := conn.Read()

	if e != nil {
		slog.Error("error while reading create room data from the client", e)
		os.Exit(2)
	}
	roomData := utils.Unmarshal[shared.TCPCommand[shared.MultiplayerRoom]](d).Data

	return &Game{
		State:        WAITING_ROOM,
		isLocal:      false,
		Text:         roomData.Text,
		Speed:        0,
		Index:        0,
		Chars:        strings.Split(roomData.Text, ""),
		RenderedText: roomData.Text,
		Room:         &roomData,
		Me:           user,
		ClientConn:   conn,
	}
}
func (g *Game) SendStartGameCommand(roomId string) *Game {
	input := []byte{shared.START_GAME}

	g.ClientConn.Write(append(input, roomId...))

	return g
}
func (g *Game) StartGame() *Game {
	g.State = IN_PROGRESS
	g.timeStarted = time.Now()
	return g
}

func (g *Game) ProcessTyping(event *tcell.EventKey) {

	if string(event.Rune()) == g.Chars[g.Index] {
		g.Index = g.Index + 1
	} else {
		// TODO: highlight error
		logger.Log.Println("highlight error here.")
	}

	if g.Index == len(g.Chars) {
		g.timeEnded = time.Now()

		totalTime := g.timeEnded.Sub(g.timeStarted)
		g.TotalTime = totalTime.String()
		totalSeconds := totalTime.Seconds()
		totalMinutes := totalSeconds / 60

		g.State = FINISHED

		wordCount := len(strings.Split(g.Text, " "))

		g.Speed = int(math.Round(float64(wordCount) / totalMinutes))
		return
	}

}
