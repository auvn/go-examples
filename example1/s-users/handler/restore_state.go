package handler

import (
	"context"
	"io"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
	"github.com/auvn/go-examples/example1/s-framework/encoding"
	"github.com/auvn/go-examples/example1/s-users/usersevent"
)

const (
	TypeDriver = "driver"
	TypeRider  = "rider"
)

type RestoreDriverStateRequest struct {
	DriverID id.ID
}

func (h *Handlers) RestoreDriverState(ctx context.Context, body io.Reader, _ io.Writer) error {
	var req RestoreDriverStateRequest
	if err := encoding.UnmarshalReader(body, &req); err != nil {
		return err
	}

	return h.Events.PublishDriverRestoredState(ctx, usersevent.DriverRestoredState{
		DriverID: req.DriverID,
	})
}

type RestoreRiderStateRequest struct {
	UserID id.ID
}

func (h *Handlers) RestoreRiderState(ctx context.Context, body io.Reader, _ io.Writer) error {
	return nil
}
