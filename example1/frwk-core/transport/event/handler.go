package event

import (
	"context"
	"fmt"
	"io"

	"github.com/auvn/go-examples/example1/frwk-core/transport"
)

type Type interface {
	String() string
}

type Publisher interface {
	Publish(ctx context.Context, e transport.Event) error
}

type HandlerFunc func(ctx context.Context, req io.Reader) error

type Dispatcher map[string]HandlerFunc

func (d Dispatcher) Handle(eventType string, fn HandlerFunc) Dispatcher {
	if _, ok := d[eventType]; ok {
		panic(fmt.Sprintf("%q: handler already registered", eventType))
	}
	d[eventType] = fn
	return d
}

func Handle(eventType string, fn HandlerFunc) Dispatcher {
	d := Dispatcher{}
	return d.Handle(eventType, fn)
}
