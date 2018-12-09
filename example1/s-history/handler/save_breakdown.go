package handler

import (
	"context"
	"io"

	"github.com/auvn/go-examples/example1/frwk-core/encoding"
	"github.com/auvn/go-examples/example1/s-history/history"
)

func (h *Handlers) SaveBreakdown(ctx context.Context, r io.Reader) error {
	var req history.Breakdown

	if err := encoding.UnmarshalReader(r, &req); err != nil {
		return err
	}

	return h.History.SaveBreakdown(ctx, req)
}
