package telnet

import "io"

// NVT Network Virtual Terminal interface
type NVT interface {
	io.Writer
	io.Reader
	Close()
}
