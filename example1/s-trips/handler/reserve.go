package handler

import (
	"context"
	"io"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
)

type ReserveRequest struct {
	RiderID id.ID
}

type ReserveResponse struct {
	RideID id.ID
}

func (h *Handlers) Reserve(ctx context.Context, body io.Reader, _ io.Writer) error {
	return nil
}
