package tcp

// Turn this into it's own module so it can be used on the client
import (
	"log/slog"
	"sort"
	"strconv"
	"typrfr/pkg/shared"
	"typrfr/pkg/utils"
)

func RunCommand(cmd byte, data []byte, conn *Connection, sockets map[int]chan Connection) {
	switch cmd {
	case shared.REQUEST_USER_ID:
		requestUserId(conn)
	case shared.CREATE_ROOM:
		createRoom(conn)
	case shared.JOIN_ROOM:
		roomId, err := strconv.ParseInt(string(data), 10, 32)
		if err != nil {
			slog.Error("error parsing room id")
		}
		joinRoom(conn, int(roomId), sockets)
	default:
		slog.Error("no valid cmd found")
	}
}

// Writes the user id to the client
func requestUserId(c *Connection) {
	user := shared.User{
		Id: c.Id,
	}

	c.Write(shared.TCPCommand[shared.User]{
		Command: shared.REQUEST_USER_ID,
		Data:    user,
	})
}

var Rooms []shared.MultiplayerRoom

func createRoom(c *Connection) {
	room := shared.MultiplayerRoom{
		Id: utils.GenCode(),
		Users: []shared.User{
			{
				Id: c.Id,
			},
		},
		Leader: c.Id,
		Text:   utils.GenText(),
	}

	slog.Info("created a room with", "id", room.Id)

	Rooms = append(Rooms, room)

	c.Write(shared.TCPCommand[shared.MultiplayerRoom]{
		Data:    room,
		Command: shared.CREATE_ROOM,
	})
}

func joinRoom(c *Connection, roomId int, sockets map[int]chan Connection) *shared.MultiplayerRoom {
	if len(Rooms) == 0 {
		slog.Info("no rooms found")
		return nil
	}
	idx := sort.Search(len(Rooms), func(i int) bool {
		return Rooms[i].Id == roomId
	})
	if idx == -1 {
		slog.Info("no room found with", "id", roomId)
		return nil
	}
	newUser := shared.User{
		Id: c.Id,
	}

	Rooms[idx].Users = append(Rooms[idx].Users, newUser)

	slog.Info("Total users in room", "val", Rooms[idx].Users)

	c.Write(shared.TCPCommand[shared.MultiplayerRoom]{
		Command: shared.JOIN_ROOM,
		Data:    Rooms[idx],
	})

	// Send a message on the client so a re-render can happen
	newUserJoined(&Rooms[idx], sockets)

	return &Rooms[idx]
}

func newUserJoined(room *shared.MultiplayerRoom, sockets map[int]chan Connection) {
	// Finding the user from connections for now due to import cycle.
	slog.Info("sockets", "len", sockets)
	for _, user := range room.Users {
		slog.Info("starting the process to send join notification")

		conn := <-sockets[user.Id]

		slog.Info("sending new user join to user", "id", user.Id)

		slog.Info("conn info", "val", conn)

		conn.Write(&shared.TCPCommand[shared.MultiplayerRoom]{
			Command: shared.NEW_USER_JOINED,
			Data:    *room,
		})

	}
}
