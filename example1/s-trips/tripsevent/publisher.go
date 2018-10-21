package tripsevent

import (
	"bytes"
	"context"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
	"github.com/auvn/go-examples/example1/s-framework/encoding"
	"github.com/auvn/go-examples/example1/s-framework/transport"
)

type Publisher struct {
	publisher transport.Publisher
}

func (p *Publisher) PublishReserved(ctx context.Context, e Reserved) error {
	return p.publishEvent(ctx, TypeReserved, e)
}

func (p *Publisher) PublishDriverFound(ctx context.Context, e DriverFound) error {
	return p.publishEvent(ctx, TypeDriverFound, e)
}

func (p *Publisher) PublishCanceled(ctx context.Context, e Canceled) error {
	return p.publishEvent(ctx, TypeCanceled, e)
}

func (p *Publisher) publishEvent(ctx context.Context, typ string, e interface{}) error {
	var buf bytes.Buffer
	if err := encoding.MarshalToWriter(e, &buf); err != nil {
		return err
	}
	_ = transport.Event{
		ID:   id.New(),
		Body: &buf,
		Type: typ,
	}
	//return p.publisher.Publish(ctx, event)
	return nil
}
