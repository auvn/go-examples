package handler

import (
	"context"
	"io"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
	"github.com/auvn/go-examples/example1/frwk-core/encoding"
	"github.com/auvn/go-examples/example1/s-gw/gwevent"
	"github.com/auvn/go-examples/example1/s-trips/event/ridersevent"
	"github.com/auvn/go-examples/example1/s-trips/event/tripsevent"
	"github.com/auvn/go-examples/example1/s-trips/trip"
)

type ReserveRequest struct {
	RiderID id.ID
}

func (h *Handlers) Reserve(ctx context.Context, body io.Reader, w io.Writer) error {
	var req ReserveRequest
	if err := encoding.UnmarshalReader(body, &req); err != nil {
		return err
	}

	newTrip := trip.Trip{
		ID:      id.New(),
		RiderID: req.RiderID,
	}

	if err := h.Trips.Create(ctx, newTrip); err != nil {
		return err
	}

	err := h.Events.PublishEvent(ctx,
		tripsevent.TypeReserved,
		tripsevent.Reserved{
			TripID:  newTrip.ID,
			RiderID: newTrip.RiderID,
		})
	if err != nil {
		return err
	}

	return gwevent.PublishUserEvent(ctx, h.Events,
		gwevent.UserEvent{
			Type:   ridersevent.TypeTripReserved,
			UserID: req.RiderID,
			Body: ridersevent.TripReserved{
				TripID: newTrip.ID,
			}},
	)
}
