package handler

import (
	"context"
	"io"
	"time"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
	"github.com/auvn/go-examples/example1/frwk-core/encoding"
	"github.com/auvn/go-examples/example1/s-trips/event/tripsevent"
)

func (h *Handlers) Complete(ctx context.Context, body io.Reader, _ io.Writer) error {
	var req struct {
		DriverID id.ID
	}

	if err := encoding.UnmarshalReader(body, &req); err != nil {
		return err
	}

	completedTrip, err := h.Trips.Complete(ctx, req.DriverID)
	if err != nil {
		return err
	}

	return h.Events.PublishEvent(ctx,
		tripsevent.TypeCompleted,
		tripsevent.Completed{
			DriverID: req.DriverID,
			TripID:   completedTrip.ID,
			RiderID:  completedTrip.RiderID,
			Distance: 100,
			Duration: time.Hour,
		})
}
