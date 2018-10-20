package transport

import (
	"context"
	"fmt"
	"io"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
)

type Message struct {
	ID        id.ID
	Recipient string
	Type      string
	Body      io.Reader
}

type Event struct {
	ID   id.ID
	Type string
	Body io.Reader
}

type Reply struct {
	Body  io.ReadCloser
	Error []byte
}

type ResponseWriter interface {
	io.Writer
}

type RequestReader interface {
	io.Reader
}

type Requester interface {
	Request(ctx context.Context, msg Message) (*Reply, error)
}

type Enqueuer interface {
	Enqueue(ctx context.Context, msg Message) error
}

type Publisher interface {
	Publish(ctx context.Context, event Event) error
}

type HandlerFunc func(ctx context.Context, req io.Reader, resp io.Writer) error
type DispatcherFunc func(ctx context.Context, body io.Reader) error

type HandlerMap map[string]HandlerFunc

func (m HandlerMap) Register(s string, fn HandlerFunc) {
	if _, ok := m[s]; ok {
		panic(fmt.Sprintf("handler already registered for %s", s))
	}
	m[s] = fn
}

type DispatcherMap map[string]DispatcherFunc

func (m DispatcherMap) Register(s string, fn DispatcherFunc) {
	if _, ok := m[s]; ok {
		panic(fmt.Sprintf("dispatcher already registered for %s", s))
	}
	m[s] = fn
}
