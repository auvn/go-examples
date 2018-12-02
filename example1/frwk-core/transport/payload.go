package transport

import (
	"io"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
)

type Message struct {
	ID        id.ID
	Recipient string
	Type      string
	Body      io.Reader
	//Deadline  time.Time
}

type Event struct {
	ID   id.ID
	Type string
	Body io.Reader
}

type Reply struct {
	Body       io.ReadCloser
	Error      string
	Successful bool
}

// TODO
type Payload struct {
	ID   id.ID
	Data interface{}
}
