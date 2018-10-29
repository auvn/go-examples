package gwevent

import (
	"context"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
	"github.com/auvn/go-examples/example1/s-framework/transport/transportutil"
)

const (
	TypeUserEvent = "gw_user_event"
)

type UserEvent struct {
	UserID id.ID
	Type   string
	Body   interface{}
}

func PublishUserEvent(ctx context.Context, p *transportutil.Publisher, event UserEvent) error {
	return p.PublishEvent(ctx, TypeUserEvent, event)
}
