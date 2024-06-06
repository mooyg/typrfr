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
	NEW_USER_JOINED
	START_GAME
	ERROR
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
	case START_GAME:
		i, err := strconv.ParseInt(string(data), 10, 32)
		if err != nil {
			fmt.Errorf("error parsing")
		}
		StartGame(conn, int(i))
	default:
		fmt.Errorf("missing command")
	}
}

type Room struct {
	Id          int
	Connections []*Connection
	Text        string
	Leader      int
	Connection  Connection
	Started     bool
}

type Rooms []*Room

var rooms Rooms = make(Rooms, 0, 10)

func CreateRoom(conn *Connection) *Room {
	room := &Room{
		Connections: []*Connection{conn},
		Id:          utils.GenCode(),
		Text:        utils.GenText(),
		Leader:      conn.Id,
		Connection:  *conn,
		Started:     false,
	}

	rooms = append(rooms, room)

	conn.Write(&TCPCommand[Room]{
		Command: CREATE_ROOM,
		Data:    *room,
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
			Connection:  *conn,
			Connections: rooms[idx].Connections,
			Text:        rooms[idx].Text,
			Leader:      rooms[idx].Leader,
			Started:     rooms[idx].Started,
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
				Connection:  *conn,
				Connections: rooms[idx].Connections,
				Text:        rooms[idx].Text,
				Leader:      rooms[idx].Leader,
				Started:     rooms[idx].Started,
			},
		})
	}
}

func StartGame(c *Connection, id int) *Room {
	if len(rooms) == 0 {
		slog.Info("no rooms found")
		c.Write(&TCPCommand[*Room]{
			Command: ERROR,
			Data:    nil,
		})
		return nil
	}

	idx := sort.Search(len(rooms), func(i int) bool {
		return rooms[i].Id == id
	})

	if idx == -1 {
		c.Write(&TCPCommand[*Room]{
			Command: ERROR,
			Data:    nil,
		})

		return nil
	}
	for _, conn := range rooms[idx].Connections {
		slog.Info("starting game for clients in room", "val", rooms[idx].Id)

		conn.Write(&TCPCommand[Room]{
			Command: START_GAME,
			Data:    *rooms[idx],
		})
	}

	return rooms[idx]
}
