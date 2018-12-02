package handler

import (
	"github.com/auvn/go-examples/example1/frwk-core/transport/event/eventutil"
	"github.com/auvn/go-examples/example1/s-trips/trip"
)

type Handlers struct {
	Events *eventutil.Publisher
	Trips  *trip.Trips
}
