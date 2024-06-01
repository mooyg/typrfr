package server

import (
	"fmt"
	"log/slog"
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
