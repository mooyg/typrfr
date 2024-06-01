package tcp

import (
	"fmt"
	"io"
	"log/slog"
	"net"
)

type TCPCommand struct {
	Command byte
	Data    []byte
}

type TCP struct {
	listener net.Listener
	sockets  []Connection
}

func NewTCPServer(port uint16) (*TCP, error) {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		return nil, err
	}

	return &TCP{
		listener: listener,
		sockets:  make([]Connection, 0, 4),
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
		}

		newConn := NewConnection(conn, id)

		id++

		t.sockets = append(t.sockets, newConn)

		newConn.Write("Hello world!")

		slog.Info("New connection", "id", newConn.Id)
		slog.Info("total sockets", "len", len(t.sockets))

		go handleConnection(&newConn)
	}
}

func handleConnection(c *Connection) {
	packet := make([]byte, 4096)

	tmp := make([]byte, 4096)

	defer c.Close()

	for {
		n, err := c.Read(tmp)

		_, err := c.ReadCommand(tmp[0])

		if err != nil {
			slog.Error("no valid command sent")
			break
		}

		c.previous = append(c.previous, tmp...)

		slog.Info("num of bytes", "n", n)

		if err != nil {

			if err != io.EOF {
				slog.Error("read error", "err", err)
			}

			slog.Info("End of file")

			break
		}
		packet = append(packet, tmp...)

	}

	c.Write(string(packet))

	fmt.Println("Received message:", string(packet))
}

func (t *TCP) Close() {
	t.listener.Close()
}
