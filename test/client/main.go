package main

import (
	"typrfr/pkg/commands"
	"typrfr/pkg/server"
)

func main() {

	conn := server.JoinServer()

	s := []byte("mooybot@gmail.com\n")

	input := append([]byte{commands.WELCOME}, s...)

	conn.Write(input)

	conn.Close()
}
