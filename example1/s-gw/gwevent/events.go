package gwevent

import (
	"context"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
	"github.com/auvn/go-examples/example1/frwk-core/transport/event/eventutil"
)

const (
	TypeUserEvent = "gw_user_event"
)

type UserEvent struct {
	UserID id.ID
	Type   string
	Body   interface{}
}

func PublishUserEvent(ctx context.Context, p *eventutil.Publisher, event UserEvent) error {
	return p.PublishEvent(ctx, TypeUserEvent, event)
}
