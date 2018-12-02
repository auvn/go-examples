package handler

import (
	"context"
	"io"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
	"github.com/auvn/go-examples/example1/frwk-core/encoding"
	"github.com/auvn/go-examples/example1/s-tracking/driver"
)

type TrackRequest struct {
	DriverID id.ID
	Busy     bool
}

func (h *Handlers) Track(ctx context.Context, body io.Reader, _ io.Writer) error {
	var req TrackRequest
	if err := encoding.UnmarshalReader(body, &req); err != nil {
		return err
	}

	updatedDriver := driver.Driver{
		ID:   req.DriverID,
		Busy: req.Busy,
	}
	if err := h.Drivers.Update(ctx, updatedDriver); err != nil {
		return err
	}

	return h.DriverHeartbeats.Heartbeat(ctx, req.DriverID)
}
