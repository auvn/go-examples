package handler

import (
	"context"
	"io"
	"time"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
	"github.com/auvn/go-examples/example1/frwk-core/encoding"
	"github.com/auvn/go-examples/example1/s-history/history"
	"github.com/auvn/go-examples/example1/s-history/historyevent"
)

func (h *Handlers) Save(ctx context.Context, r io.Reader) error {
	var trip struct {
		TripID   id.ID
		RiderID  id.ID
		DriverID id.ID

		Distance int
		Duration time.Duration
	}

	if err := encoding.UnmarshalReader(r, &trip); err != nil {
		return err
	}

	historyRecord := history.Record{
		ID:        trip.TripID,
		TripID:    trip.TripID,
		RiderID:   trip.RiderID,
		DriverID:  trip.DriverID,
		Distance:  trip.Distance,
		Duration:  trip.Duration,
		CreatedAt: time.Now().UTC(),
	}
	if err := h.History.Save(ctx, historyRecord); err != nil {
		return err
	}

	return h.Events.PublishEvent(ctx,
		historyevent.TypeTripAdded,
		historyevent.TripAdded{
			TripID:   trip.TripID,
			Distance: trip.Distance,
			Duration: trip.Duration,
		})
}
