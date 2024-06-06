package tcp

// Turn this into it's own module so it can be used on the client
import (
	"fmt"
	"log/slog"
	"sort"
	"strconv"
	"typrfr/pkg/utils"
)

const (
	WELCOME = iota
	CREATE_ROOM
	JOIN_ROOM
	ERROR
	NEW_USER_JOINED
)

type TCPCommand[T any] struct {
	Command byte
	Data    T
}

func RunCommand(cmd byte, data []byte, conn *Connection) {
	switch cmd {
	case WELCOME:
	case CREATE_ROOM:
		CreateRoom(conn)
	case JOIN_ROOM:
		i, err := strconv.ParseInt(string(data), 10, 32)
		if err != nil {
			fmt.Errorf("error parsing")
		}
		JoinRoom(conn, int(i))
	default:
		fmt.Errorf("missing command")
	}
}

type Room struct {
	Id          int
	Connections []*Connection
	Text        string
}

type Rooms []*Room

var rooms Rooms = make(Rooms, 0, 10)

func CreateRoom(conn *Connection) *Room {
	room := &Room{
		Connections: []*Connection{conn},
		Id:          utils.GenCode(),
		Text:        utils.GenText(),
	}

	rooms = append(rooms, room)

	conn.Write(&TCPCommand[Room]{
		Command: CREATE_ROOM,
		Data: Room{
			Id:          room.Id,
			Text:        room.Text,
			Connections: room.Connections,
		},
	})

	return room
}

func JoinRoom(conn *Connection, id int) *Room {
	if len(rooms) == 0 {
		slog.Info("no rooms found")
		conn.Write(&TCPCommand[*Room]{
			Command: ERROR,
			Data:    nil,
		})
		return nil
	}

	idx := sort.Search(len(rooms), func(i int) bool {
		return rooms[i].Id == id
	})

	if idx == -1 {
		conn.Write(&TCPCommand[*Room]{
			Command: ERROR,
			Data:    nil,
		})

		return nil
	}

	rooms[idx].Connections = append(rooms[idx].Connections, conn)

	UserJoined(rooms[idx])

	conn.Write(&TCPCommand[Room]{
		Command: JOIN_ROOM,
		Data: Room{
			Id:          rooms[idx].Id,
			Text:        rooms[idx].Text,
			Connections: rooms[idx].Connections,
		},
	})
	return rooms[idx]
}

func UserJoined(room *Room) {
	idx := sort.Search(len(rooms), func(i int) bool {
		return rooms[i].Id == room.Id
	})

	for _, conn := range rooms[idx].Connections {
		slog.Info("writing to clients", "id", conn.Id)
		conn.Write(&TCPCommand[Room]{
			Command: NEW_USER_JOINED,
			Data: Room{
				Id:          rooms[idx].Id,
				Text:        rooms[idx].Text,
				Connections: rooms[idx].Connections,
			},
		})
	}
}
