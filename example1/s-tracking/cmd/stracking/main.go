package main

import (
	"context"

	"github.com/auvn/go-examples/example1/frwk-core/service"
	"github.com/auvn/go-examples/example1/frwk-core/transport/event/eventutil"
	"github.com/auvn/go-examples/example1/frwk-storage/nosql/monga"
	"github.com/auvn/go-examples/example1/frwk-transport/hottabych"
	"github.com/auvn/go-examples/example1/frwk-transport/natsss"
	"github.com/auvn/go-examples/example1/s-tracking/driver"
	"github.com/auvn/go-examples/example1/s-tracking/event/driversevent"
	"github.com/auvn/go-examples/example1/s-tracking/handler"
	"github.com/auvn/go-examples/example1/s-trips/event/tripsevent"
)

func main() {
	mongaClient := monga.MustNew(monga.EnvConfig())
	drivers := driver.NewDrivers(mongaClient)

	events := eventutil.NewPublisher(natsss.NewClient(natsss.EnvClientConfig()))

	natsssServer := natsss.NewStreams(natsss.EnvStreamConfig())

	handlers := handler.Handlers{
		Events:           events,
		Drivers:          drivers,
		DriverHeartbeats: driversevent.NewHeartbeats(events),
	}

	httpServer := hottabych.
		Handle("Track", handlers.Track)

	natsssServer.
		Subscribe(tripsevent.TypeReserved, handlers.SuggestTripDriver)

	service.Serve(context.Background(),
		httpServer,
		natsssServer)
}
