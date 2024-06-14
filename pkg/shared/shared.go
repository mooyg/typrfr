package shared

type TCPCommand[T any] struct {
	Command byte
	Data    T
}

const (
	REQUEST_USER_ID = iota
	CREATE_ROOM
	JOIN_ROOM
	NEW_USER_JOINED
	START_GAME
	END_GAME
	ERROR
)

type User struct {
	Id int
}

type MultiplayerRoom struct {
	// Room id given to the room
	Id    int
	Users []User
	// User id of the leader in the room
	Leader int
	// Sentence assigned for the room
	Text       string
	InProgress bool
}
