package main

import (
	"context"

	"github.com/auvn/go-examples/example1/frwk-core/servegroup"
	"github.com/auvn/go-examples/example1/frwk-transport/natsss"
	"github.com/auvn/go-examples/example1/s-calculations/calculationsevent"
	"github.com/auvn/go-examples/example1/s-dwh/handler"
	"github.com/auvn/go-examples/example1/s-tracking/event/trackingevent"
	"github.com/auvn/go-examples/example1/s-trips/event/tripsevent"
)

func main() {
	server := natsss.NewStreams(natsss.EnvStreamConfig())

	handlers := handler.Handlers{}
	server.
		Subscribe(trackingevent.TypeTripDriverSuggested, handlers.Consume).
		Subscribe(calculationsevent.TypeBreakdownCalculated, handlers.Consume).
		Subscribe(tripsevent.TypeCompleted, handlers.Consume).
		Subscribe(tripsevent.TypeReserved, handlers.Consume)

	servegroup.Serve(context.Background(), server)
}
