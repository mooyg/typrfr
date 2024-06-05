package tcpclient

import (
	"bufio"
	"log/slog"
	"net"
	"os"
)

type TCPClient struct {
	conn *net.TCPConn
}

func New() *TCPClient {
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		Port: 5555,
	})

	if err != nil {
		slog.Error("error occured while establishing connection")
		os.Exit(1)
	}

	return &TCPClient{
		conn: conn,
	}

}

type Room struct {
	Id   int
	Text string
}

func (c *TCPClient) Read() (data []byte, err error) {
	d, err := bufio.NewReader(c.conn).ReadString('\n')
	tmp := []byte(d)
	dataWithoutEOL := tmp[0 : len(tmp)-1]

	return dataWithoutEOL, err
}

func (c *TCPClient) Write(s []byte) (n int, err error) {
	return c.conn.Write(s)
}
