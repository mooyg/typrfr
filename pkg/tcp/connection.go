package tcp

import (
	"log/slog"
	"net"
	"typrfr/pkg/commands"
	"typrfr/pkg/tcp"
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

func (c *Connection) ReadCommand(cmd byte) (tcmd *tcp.TCPCommand, err error) {
	slog.Info("rcvd", "val", string(cmd))
	return commands.GetCommandByByte(cmd)
}
func (c *Connection) Read(tmp []byte) (n int, err error) {
	return c.conn.Read(tmp)
}
func (c *Connection) Write(s string) (n int, err error) {
	return c.conn.Write([]byte(s))
}
