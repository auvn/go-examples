package main

import (
	"context"

	"github.com/auvn/go-examples/example1/frwk-core/service"
	"github.com/auvn/go-examples/example1/frwk-core/transport/event/eventutil"
	"github.com/auvn/go-examples/example1/frwk-storage/nosql/monga"
	"github.com/auvn/go-examples/example1/frwk-transport/hottabych"
	"github.com/auvn/go-examples/example1/frwk-transport/natsss"
	"github.com/auvn/go-examples/example1/s-tracking/event/trackingevent"
	"github.com/auvn/go-examples/example1/s-trips/handler"
	"github.com/auvn/go-examples/example1/s-trips/trip"
)

func main() {
	mongaClient := monga.MustNew(
		monga.Config{
			Hosts: []string{"localhost:27017"},
			Name:  "strips",
		})
	natsssServer := natsss.NewServer(natsss.ServerConfig{
		Name: "strips",
	})

	events := eventutil.NewPublisher(
		natsss.NewClient(natsss.ClientConfig{
			Name: "strips",
		}))

	handlers := handler.Handlers{
		Trips:  trip.NewTrips(mongaClient),
		Events: events,
	}
	httpTransport := hottabych.
		Handle("Reserve", handlers.Reserve).
		Handle("Complete", handlers.Complete)

	natsssServer.
		Subscribe(trackingevent.TypeTripDriverSuggested, handlers.AssignDriver)

	service.Serve(context.Background(),
		httpTransport,
		natsssServer)
}
