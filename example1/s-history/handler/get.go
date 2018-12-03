package handler

import (
	"context"
	"io"
	"log"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
	"github.com/auvn/go-examples/example1/frwk-core/encoding"
	"github.com/auvn/go-examples/example1/frwk-core/transport/event/eventutil"
	"github.com/auvn/go-examples/example1/s-history/history"
)

type Handlers struct {
	History *history.History
	Events  *eventutil.Publisher
}

type GetRequest struct {
	RiderID id.ID
}

func (h *Handlers) Get(ctx context.Context, body io.Reader, resp io.Writer) error {
	var req GetRequest

	if err := encoding.UnmarshalReader(body, &req); err != nil {
		return err
	}

	log.Printf("%+v\n", req)

	record, err := h.History.LastByRider(ctx, req.RiderID)
	if err != nil {
		return err
	}

	return encoding.MarshalToWriter(record, resp)
}
