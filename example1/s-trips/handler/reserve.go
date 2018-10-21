package handler

import (
	"context"
	"io"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
	"github.com/auvn/go-examples/example1/s-framework/encoding"
	"github.com/auvn/go-examples/example1/s-trips/trip"
	"github.com/auvn/go-examples/example1/s-trips/tripsevent"
	"github.com/pkg/errors"
)

type ReserveRequest struct {
	RiderID id.ID
}

type ReserveResponse struct {
	TripID id.ID
}

func (h *Handlers) Reserve(ctx context.Context, body io.Reader, w io.Writer) error {
	var req ReserveRequest
	if err := encoding.UnmarshalReader(body, &req); err != nil {
		return err
	}
	newTrip := trip.Trip{
		ID:    id.New(),
		Rider: req.RiderID,
	}

	if err := h.Trips.Create(ctx, newTrip); err != nil {
		switch errors.Cause(err) {
		case trip.ErrActiveExists:
			return err
		default:
			return err
		}
	}

	err := h.Events.PublishReserved(ctx, tripsevent.Reserved{
		TripID:  newTrip.ID,
		RiderID: newTrip.Rider,
	})
	if err != nil {
		return err
	}

	resp := ReserveResponse{
		TripID: newTrip.ID,
	}
	return encoding.MarshalToWriter(resp, w)
}
