package commands

import (
	"fmt"
	"typrfr/pkg/tcp"
)

const (
	WELCOME = iota
	CREATE_USER
	CREATE_ROOM
	JOIN_ROOM
)

var commandMap = map[byte]string{
	WELCOME:     "welcome",
	CREATE_USER: "createUser",
	CREATE_ROOM: "createRoom",
	JOIN_ROOM:   "joinRoom",
}

func GetCommandByByte(b byte) (cmd *tcp.TCPCommand, err error) {
	if _, ok := commandMap[b]; ok {
		return &tcp.TCPCommand{
			Command: b,
		}, nil
	}
	return nil, fmt.Errorf("missing command")
}
