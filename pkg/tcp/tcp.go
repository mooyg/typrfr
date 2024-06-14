package tcp

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
)

type TCP struct {
	listener net.Listener
	sockets  map[int]chan Connection
}

func NewTCPServer(port uint16) (*TCP, error) {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		return nil, err
	}

	return &TCP{
		listener: listener,
		sockets:  make(map[int]chan Connection),
	}, nil

}

func (t *TCP) Start() {

	id := 0

	slog.Info("Accepting connections", "addr", t.listener.Addr())

	defer t.Close()

	for {
		conn, err := t.listener.Accept()

		if err != nil {
			slog.Error("server error: ", "err", err)
			os.Exit(2)
		}

		newConn := NewConnection(conn, id)

		slog.Info("new conn", "val", newConn)

		id++

		t.sockets[id] = make(chan Connection, 1)
		t.sockets[id] <- newConn

		slog.Info("total sockets", "len", len(t.sockets))

		go handleConnection(&newConn, t.sockets)
	}
}

func handleConnection(c *Connection, sockets map[int]chan Connection) {
	packet := make([]byte, 4096)

	tmp := make([]byte, 4096)

	defer c.Close()

	for {
		data, err := c.Read()

		c.ParseMessage([]byte(data), sockets)

		slog.Info("num of bytes", "n", len(data))

		if err != nil {

			if err != io.EOF {
				slog.Error("read error", "err", err)

			}

			slog.Info("End of file")

			break
		}

		packet = append(packet, tmp...)

	}

	fmt.Println("Received message:", string(packet))
}

func (t *TCP) Close() {
	t.listener.Close()
}
