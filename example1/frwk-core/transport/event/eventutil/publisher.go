package eventutil

import (
	"bytes"
	"context"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
	"github.com/auvn/go-examples/example1/frwk-core/encoding"
	"github.com/auvn/go-examples/example1/frwk-core/transport"
	"github.com/auvn/go-examples/example1/frwk-core/transport/event"
)

type Publisher struct {
	Publisher event.Publisher
}

func (p *Publisher) PublishEvent(ctx context.Context, eventType string, v interface{}) error {
	var buf bytes.Buffer
	if err := encoding.MarshalToWriter(v, &buf); err != nil {
		return nil
	}
	return p.Publisher.Publish(ctx, transport.Event{
		Body: &buf,
		ID:   id.New(),
		Type: eventType,
	})
}

func NewPublisher(p event.Publisher) *Publisher {
	return &Publisher{
		Publisher: p,
	}
}
