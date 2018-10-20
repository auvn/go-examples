package handler

import (
	"context"
	"io"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
)

const (
	TypeDriver = "driver"
	TypeRider  = "rider"
)

type RestoreDriverStateRequest struct {
	UserID id.ID
}

func (h *Handlers) RestoreDriverState(ctx context.Context, body io.Reader, _ io.Writer) error {
	return nil
}

type RestoreRiderStateRequest struct {
	UserID id.ID
}

func (h *Handlers) RestoreRiderState(ctx context.Context, body io.Reader, _ io.Writer) error {
	return nil
}
