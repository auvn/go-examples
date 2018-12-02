package message

import (
	"context"
	"fmt"
	"io"

	"github.com/auvn/go-examples/example1/frwk-core/transport"
)

type Requester interface {
	Request(ctx context.Context, msg transport.Message) (*transport.Reply, error)
}

type HandlerFunc func(ctx context.Context, req io.Reader, resp io.Writer) error

type Dispatcher map[string]HandlerFunc

func (d Dispatcher) Handle(s string, fn HandlerFunc) Dispatcher {
	if _, ok := d[s]; ok {
		panic(fmt.Sprintf("handler already registered for %s", s))
	}
	d[s] = fn
	return d
}

func Handle(messageType string, fn HandlerFunc) Dispatcher {
	d := Dispatcher{}
	return d.Handle(messageType, fn)
}
