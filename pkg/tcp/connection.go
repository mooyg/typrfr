package tcp

import (
	"bufio"
	"encoding/json"
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

func (c *Connection) Write(s any) (n int, err error) {
	input := c.Encode(s)
	return c.conn.Write(append(input, '\n'))
}

func (c *Connection) Encode(val any) []byte {
	bytes, err := json.Marshal(val)

	if err != nil {
		slog.Error("some error occured while marshaling")
	}

	return bytes

}
func (c *Connection) Decode(val any) any {
	return val
}
