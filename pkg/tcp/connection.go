package tcp

import (
	"bufio"
	"encoding/json"
	"log/slog"
	"net"
)

type Connection struct {
	conn net.Conn
	Id   int
}

func NewConnection(conn net.Conn, id int) Connection {
	return Connection{
		conn: conn,
		Id:   id,
	}
}

func (c *Connection) Close() {
	c.conn.Close()
}

// Seperate the command rcvd and the data
func (c *Connection) ParseMessage(tmp []byte, sockets map[int]chan Connection) {
	if len(tmp) == 0 {
		slog.Error("rcvd invalid data")
		return
	}

	slog.Info("rcvd message", "msg", string(tmp))
	// Ignore EOL from the data and run the command
	RunCommand(tmp[0], tmp[1:len(tmp)-1], c, sockets)
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
