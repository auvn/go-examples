package handler

import (
	"context"
	"io"
	"log"

	"github.com/auvn/go-examples/example1/frwk-core/encoding"
)

type Handlers struct {
}

func (h *Handlers) Consume(ctx context.Context, r io.Reader) error {
	var req encoding.RawMessage
	if err := encoding.UnmarshalReader(r, &req); err != nil {
		return err
	}

	log.Printf("consumed: %s\n", req)
	return nil
}
