package dispatcher

import (
	"github.com/auvn/go-examples/example1/s-trips/trip"
	"github.com/auvn/go-examples/example1/s-trips/tripsevent"
)

type Dispatchers struct {
	Trips  *trip.Trips
	Events *tripsevent.Publisher
}
