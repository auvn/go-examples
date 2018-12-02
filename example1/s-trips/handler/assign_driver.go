package handler

import (
	"context"
	"io"
	"log"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
	"github.com/auvn/go-examples/example1/frwk-core/encoding"
	"github.com/auvn/go-examples/example1/s-gw/gwevent"
	"github.com/auvn/go-examples/example1/s-trips/event/ridersevent"
	"github.com/auvn/go-examples/example1/s-trips/trip"
)

func (d *Handlers) AssignDriver(ctx context.Context, body io.Reader) error {
	var foundDriver struct {
		TripID   id.ID
		DriverID id.ID
	}

	if err := encoding.UnmarshalReader(body, &foundDriver); err != nil {
		return err
	}

	activeTrip, err := d.Trips.ByID(ctx, foundDriver.TripID)
	if err != nil {
		return err
	}

	log.Printf("assigning driver %q to trip %q\n", foundDriver.DriverID, foundDriver.TripID)

	if err := d.Trips.AssignDriver(ctx,
		trip.Driver{ID: foundDriver.DriverID, TripID: foundDriver.TripID}); err != nil {
		return err
	}

	return gwevent.PublishUserEvent(ctx, d.Events,
		gwevent.UserEvent{
			UserID: activeTrip.RiderID,
			Type:   ridersevent.TypeDriverFound,
			Body: ridersevent.DriverFound{
				TripID:   activeTrip.ID,
				DriverID: foundDriver.DriverID,
			},
		})
}
