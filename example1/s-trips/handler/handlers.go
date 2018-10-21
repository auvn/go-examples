package handler

import (
	"github.com/auvn/go-examples/example1/s-trips/trip"
	"github.com/auvn/go-examples/example1/s-trips/tripsevent"
)

type Handlers struct {
	Events *tripsevent.Publisher
	Trips  *trip.Trips
}
