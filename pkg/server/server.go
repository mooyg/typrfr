package server

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"typrfr/pkg/tcp"
)

func CreateServer(port uint16) (*tcp.TCP, error) {
	server, err := tcp.NewTCPServer(port)

	slog.Info("Starting server on ", "port", port)

	if err != nil {
		return nil, fmt.Errorf("Error creating server %w", err)
	}

	return server, nil
}

func JoinServer() net.Conn {
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		Port: 5555,
	})

	if err != nil {
		slog.Error("error occured while establishing connection")
		os.Exit(1)
	}
	return conn
}
