package handler

import (
	"context"
	"io"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
	"github.com/auvn/go-examples/example1/s-history/history"
)

type Handlers struct {
	history *history.History
}

type GetRequest struct {
	RiderID id.ID
}

type GetResponse struct {
	Count int
}

func (h *Handlers) Get(ctx context.Context, body io.Reader, resp io.Writer) error {
	return nil
}
