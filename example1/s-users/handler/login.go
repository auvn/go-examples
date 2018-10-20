package handler

import (
	"context"
	"io"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
)

type AuthenticateRiderRequest struct {
	UserID id.ID
}

func (h *Handlers) AuthenticateRider(ctx context.Context, in io.Reader, _ io.Writer) error {
	return nil
}

type AuthenticateDriverRequest struct {
	UserID id.ID
}

func (h *Handlers) AuthenticateDriver(ctx context.Context, in io.Reader, _ io.Writer) error {
	return nil
}
