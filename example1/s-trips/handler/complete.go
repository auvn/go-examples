package handler

import (
	"context"
	"io"
	"time"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
	"github.com/auvn/go-examples/example1/s-framework/encoding"
	"github.com/auvn/go-examples/example1/s-trips/tripsevent"
)

func (h *Handlers) Complete(ctx context.Context, body io.Reader, _ io.Writer) error {
	var req struct {
		TripID   id.ID
		DriverID id.ID
	}

	if err := encoding.UnmarshalReader(body, &req); err != nil {
		return err
	}

	return h.Events.PublishEvent(ctx, tripsevent.TypeCompleted, tripsevent.Completed{
		DriverID: req.DriverID,
		TripID:   req.TripID,
		RiderID:  id.New(),
		Distance: 100,
		Duration: time.Hour,
	})
}
