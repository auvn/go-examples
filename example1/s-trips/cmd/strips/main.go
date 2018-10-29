package main

import (
	"context"

	"github.com/auvn/go-examples/example1/s-framework/service"
	"github.com/auvn/go-examples/example1/s-framework/storage/redis"
	"github.com/auvn/go-examples/example1/s-framework/transport/hottabych"
	"github.com/auvn/go-examples/example1/s-framework/transport/natsss"
	"github.com/auvn/go-examples/example1/s-framework/transport/transportutil"
	"github.com/auvn/go-examples/example1/s-tracking/trackingevent"
	"github.com/auvn/go-examples/example1/s-trips/dispatcher"
	"github.com/auvn/go-examples/example1/s-trips/handler"
	"github.com/auvn/go-examples/example1/s-trips/trip"
	"github.com/auvn/go-examples/example1/s-users/usersevent"
)

func main() {
	r := redis.New(redis.Config{Addrs: []string{"localhost:7001", "localhost:7000"}})

	events := transportutil.NewPublisher(
		natsss.NewClient(natsss.ClientConfig{
			Name:        "strips",
			ClusterName: "test-cluster",
		}))

	handlers := handler.Handlers{
		Trips:  trip.NewTrips(r),
		Events: events,
	}
	httpTransport := hottabych.
		Handle("Reserve", handlers.Reserve).
		Handle("Complete", handlers.Complete)

	dispatchers := dispatcher.Dispatchers{
		Events: events,
	}
	natsssServer := natsss.NewServer(natsss.ServerConfig{
		Name:        "strips",
		ClusterName: "test-cluster",
	})
	natsssServer.
		Subscribe(usersevent.TypeDriverRestoredState, dispatchers.RestoreActiveTrip).
		Subscribe(trackingevent.TypeDriverFound, dispatchers.AssignDriver)

	service.Serve(context.Background(),
		httpTransport,
		natsssServer)
}
