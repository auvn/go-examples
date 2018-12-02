package handler

import (
	"context"
	"io"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
	"github.com/auvn/go-examples/example1/frwk-core/encoding"
	"github.com/auvn/go-examples/example1/s-tracking/event/trackingevent"
)

func (d *Handlers) SuggestTripDriver(ctx context.Context, body io.Reader) error {
	var tripReserved struct {
		TripID id.ID
	}

	if err := encoding.UnmarshalReader(body, &tripReserved); err != nil {
		return err
	}

	driver, err := d.Drivers.Lookup(ctx)
	if err != nil {
		return err
	}

	return d.Events.PublishEvent(ctx,
		trackingevent.TypeTripDriverSuggested,
		trackingevent.TripDriverSuggested{
			DriverID: driver,
			TripID:   tripReserved.TripID,
		})

}
