package dispatcher

import (
	"context"
	"io"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
	"github.com/auvn/go-examples/example1/s-framework/encoding"
	"github.com/auvn/go-examples/example1/s-trips/tripsevent"
)

func (d *Dispatchers) FindProvider(ctx context.Context, body io.Reader) error {
	var event tripsevent.Reserved
	if err := encoding.UnmarshalReader(body, &event); err != nil {
		return err
	}

	return d.Events.PublishDriverFound(ctx, tripsevent.DriverFound{
		DriverID: id.New(),
		TripID:   event.TripID,
	})
}
