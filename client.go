package telnet

import (
	"fmt"
	"io"
	"net"
	"time"
)

type tclient struct {
	remoteAddr string
	out        io.Writer
	conn       net.Conn
}

// NewClient return a new Network Virtual Terminal
func NewClient(host string, w io.Writer) (NVT, error) {
	conn, err := net.DialTimeout("tcp", host, 30*time.Second)
	if err != nil {
		return nil, err
	}

	return &tclient{
		remoteAddr: host,
		out:        w,
		conn:       conn,
	}, nil
}

func (c *tclient) Write(p []byte) (n int, err error) {
	n = len(p)
	fmt.Fprint(c.out, p)
	return
}

func (c *tclient) Read(p []byte) (n int, err error) {

	return
}

func (c *tclient) Close() {
	c.conn.Close()
}
