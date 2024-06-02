package tcp

import (
	"bufio"
	"log/slog"
	"net"
	"typrfr/pkg/commands"
)

type Connection struct {
	conn     net.Conn
	Id       int
	previous []byte
}

func NewConnection(conn net.Conn, id int) Connection {
	return Connection{
		conn:     conn,
		Id:       id,
		previous: []byte{},
	}
}

func (c *Connection) Close() {
	c.conn.Close()
}

func (c *Connection) ParseMessage(tmp []byte) {

	if len(tmp) == 0 {
		slog.Error("rcvd invalid data")
		return
	}

	// Ignore \n from the data

	cmd, err := commands.Command(tmp[0], tmp[1:len(tmp)-1])

	if err != nil {
		slog.Error("no valid command found")
	}
	slog.Info("data rcvd with command", "data", string(cmd.Data))
}

func (c *Connection) Read() (data string, err error) {
	return bufio.NewReader(c.conn).ReadString('\n')
}
func (c *Connection) Write(s string) (n int, err error) {
	return c.conn.Write([]byte(s))
}
