package handler

import (
	"github.com/auvn/go-examples/example1/s-framework/transport/transportutil"
	"github.com/auvn/go-examples/example1/s-trips/trip"
)

type Handlers struct {
	Events *transportutil.Publisher
	Trips  *trip.Trips
}
