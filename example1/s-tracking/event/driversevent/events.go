package driversevent

import (
	"context"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
	"github.com/auvn/go-examples/example1/frwk-core/transport/event/eventutil"
	"github.com/auvn/go-examples/example1/s-gw/gwevent"
)

const (
	TypeAssignedTrip = "assigned_trip"
	TypeHeartbeat    = "heartbeat"
)

type Heartbeat struct{}

type AssignedTrip struct {
	TripID id.ID
}

type Heartbeats struct {
	publisher *eventutil.Publisher
}

func (h *Heartbeats) Heartbeat(ctx context.Context, driverID id.ID) error {
	return gwevent.PublishUserEvent(ctx,
		h.publisher,
		gwevent.UserEvent{
			Type:   TypeHeartbeat,
			UserID: driverID,
			Body:   Heartbeat{},
		})
}

func NewHeartbeats(p *eventutil.Publisher) *Heartbeats {
	return &Heartbeats{
		publisher: p,
	}
}
