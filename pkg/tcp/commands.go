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
)

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
	Sentence    string
}

type Rooms []*Room

var rooms Rooms = make(Rooms, 0, 10)

func CreateRoom(conn *Connection) *Room {
	room := &Room{
		Connections: []*Connection{conn},
		Id:          utils.GenCode(),
		Sentence:    utils.GenText(),
	}

	rooms = append(rooms, room)

	slog.Info("created room", "id", room.Id)

	conn.Write(fmt.Sprintf("%d\n", room.Id))

	return room
}
func JoinRoom(conn *Connection, id int) *Room {
	idx := sort.Search(len(rooms), func(i int) bool {
		return rooms[i].Id == id
	})
	rooms[idx].Connections = append(rooms[idx].Connections, conn)

	UserJoined(rooms[idx])

	return rooms[idx]
}

func UserJoined(room *Room) {
	idx := sort.Search(len(rooms), func(i int) bool {
		return rooms[i].Id == room.Id
	})
	for _, conn := range rooms[idx].Connections {
		conn.Write("New user joined room")
	}
}
