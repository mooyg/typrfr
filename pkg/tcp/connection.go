package tcp

import "net"

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

func (c *Connection) Write(b []byte) (n int, err error) {
	return c.conn.Write(b)
}
func (c *Connection) Read(tmp []byte) (n int, err error) {
	return c.conn.Read(tmp)
}
