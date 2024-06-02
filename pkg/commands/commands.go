package commands

import (
	"fmt"
)

type TCPCommand struct {
	Command byte
	Data    []byte
}

const (
	WELCOME = iota
	START
)

func Command(cmd byte, data []byte) (tcmd *TCPCommand, err error) {
	switch cmd {
	case WELCOME:
		return &TCPCommand{
			Command: WELCOME,
			Data:    data,
		}, nil
	case START:
		return &TCPCommand{
			Command: START,
			Data:    data,
		}, nil
	default:
		return nil, fmt.Errorf("missing command")
	}
}
