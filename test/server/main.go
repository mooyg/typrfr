package main

import (
	"log/slog"
	"os"
	"typrfr/pkg/server"
)

func main() {
	server, err := server.CreateServer(5555)

	if err != nil {
		slog.Error("could not start the server", "err", err)
		os.Exit(1)
	}

	defer server.Close()

	server.Start()
}
