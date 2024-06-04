package tcp

import (
	"bufio"
	"log/slog"
	"net"
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

	slog.Info("rcvd message", "msg", string(tmp))
	// Ignore EOL from the data
	RunCommand(tmp[0], tmp[1:len(tmp)-1], c)
}

func (c *Connection) Read() (data string, err error) {
	return bufio.NewReader(c.conn).ReadString('\n')
}
func (c *Connection) Write(s string) (n int, err error) {
	return c.conn.Write([]byte(s))
}
