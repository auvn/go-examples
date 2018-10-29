package handler

import (
	"context"
	"io"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
	"github.com/auvn/go-examples/example1/s-framework/encoding"
	"github.com/auvn/go-examples/example1/s-gw/gwevent"
	"github.com/auvn/go-examples/example1/s-tracking/driversevent"
)

type TrackRequest struct {
	DriverID id.ID
}

func (h *Handlers) Track(ctx context.Context, body io.Reader, _ io.Writer) error {
	var req TrackRequest
	if err := encoding.UnmarshalReader(body, &req); err != nil {
		return err
	}

	if err := h.Drivers.Update(ctx, req.DriverID); err != nil {
		return err
	}

	if err := h.Events.PublishEvent(ctx, gwevent.TypeUserEvent, gwevent.UserEvent{
		UserID: req.DriverID,
		Type:   driversevent.TypeHeartbeat,
		Body:   driversevent.Heartbeat{},
	}); err != nil {
		return err
	}

	return nil
}
