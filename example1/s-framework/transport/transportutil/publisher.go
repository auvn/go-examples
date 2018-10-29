package transportutil

import (
	"bytes"
	"context"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
	"github.com/auvn/go-examples/example1/s-framework/encoding"
	"github.com/auvn/go-examples/example1/s-framework/transport"
)

type Publisher struct {
	Publisher transport.Publisher
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

func NewPublisher(p transport.Publisher) *Publisher {
	return &Publisher{
		Publisher: p,
	}
}
