package transport

import (
	"github.com/micro/go-micro/transport"
	"github.com/micro/go-micro/tunnel"
)

type tunListener struct {
	l tunnel.Listener
}

func (t *tunListener) Addr() string {
	return t.l.Addr()
}

func (t *tunListener) Close() error {
	return t.l.Close()
}

func (t *tunListener) Accept(fn func(socket transport.Socket)) error {
	for {
		// accept connection
		c, err := t.l.Accept()
		if err != nil {
			return err
		}
		// execute the function
		go fn(c)
	}
}
