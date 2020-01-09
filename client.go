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

// NewNVT return a new Network Virtual Terminal
func NewClient(server string, w io.Writer) (NVT, error) {
	conn, err := net.DialTimeout("tcp", server, 30*time.Second)
	if err != nil {
		return nil, err
	}

	return &tclient{
		remoteAddr: server,
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
