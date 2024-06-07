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
	END_GAME
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

type User struct {
	Conn *Connection
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

type Rooms []*Room

var rooms Rooms = make(Rooms, 0, 10)

func CreateRoom(conn *Connection) *Room {
	// Initialise the leader
	room := &Room{
		Users: []User{
			User{
				Conn: conn,
				Id:   conn.Id,
				Wpm:  0,
			},
		},
		Id:      utils.GenCode(),
		Text:    utils.GenText(),
		Leader:  conn.Id,
		Started: false,
		MyId:    conn.Id,
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

	user := User{
		Id:   conn.Id,
		Conn: conn,
		Wpm:  0,
	}
	rooms[idx].Users = append(rooms[idx].Users, user)

	UserJoined(rooms[idx])

	conn.Write(&TCPCommand[Room]{
		Command: JOIN_ROOM,
		Data: Room{
			Id:      rooms[idx].Id,
			Users:   rooms[idx].Users,
			Text:    rooms[idx].Text,
			Leader:  rooms[idx].Leader,
			Started: rooms[idx].Started,
			MyId:    conn.Id,
		},
	})
	return rooms[idx]
}

func UserJoined(room *Room) {
	idx := sort.Search(len(rooms), func(i int) bool {
		return rooms[i].Id == room.Id
	})

	for _, user := range rooms[idx].Users {
		slog.Info("writing to clients", "id", user.Id)

		user.Conn.Write(&TCPCommand[Room]{
			Command: NEW_USER_JOINED,
			Data: Room{
				Id:      rooms[idx].Id,
				Users:   rooms[idx].Users,
				Text:    rooms[idx].Text,
				Leader:  rooms[idx].Leader,
				Started: rooms[idx].Started,
				MyId:    user.Conn.Id,
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
	for _, user := range rooms[idx].Users {
		slog.Info("starting game for clients in room", "val", rooms[idx].Id)

		user.Conn.Write(&TCPCommand[Room]{
			Command: START_GAME,
			Data:    *rooms[idx],
		})
	}

	return rooms[idx]
}

func EndGame(c *Connection, id int) *Room {
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

	for _, user := range rooms[idx].Users {
		slog.Info("ending game", "val", rooms[idx].Id)

		user.Conn.Write(&TCPCommand[Room]{
			Command: START_GAME,
			Data:    *rooms[idx],
		})
	}

	return rooms[idx]
}
