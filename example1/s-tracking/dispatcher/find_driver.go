package dispatcher

import (
	"context"
	"io"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
	"github.com/auvn/go-examples/example1/s-framework/encoding"
	"github.com/auvn/go-examples/example1/s-gw/gwevent"
	"github.com/auvn/go-examples/example1/s-tracking/driversevent"
	"github.com/auvn/go-examples/example1/s-tracking/trackingevent"
)

func (d *Dispatchers) FindDriver(ctx context.Context, body io.Reader) error {
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

	if err := d.Events.PublishEvent(ctx, gwevent.TypeUserEvent, gwevent.UserEvent{
		Type:   driversevent.TypeAssignedTrip,
		UserID: driver,
		Body: driversevent.AssignedTrip{
			TripID: tripReserved.TripID,
		},
	}); err != nil {
		return err
	}

	if err := d.Drivers.Lock(ctx, driver); err != nil {
		return err
	}

	return d.Events.PublishEvent(ctx, trackingevent.TypeDriverFound, trackingevent.DriverFound{
		DriverID: driver,
		TripID:   tripReserved.TripID,
	})

}
