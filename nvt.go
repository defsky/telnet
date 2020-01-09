package telnet

import "io"

type NVT interface {
	io.Writer
	io.Reader
}
