package main

import (
	"context"

	"github.com/auvn/go-examples/example1/s-framework/service"
	"github.com/auvn/go-examples/example1/s-framework/transport/hottabych"
	"github.com/auvn/go-examples/example1/s-framework/transport/natsss"
	"github.com/auvn/go-examples/example1/s-framework/transport/transportutil"
	"github.com/auvn/go-examples/example1/s-tracking/dispatcher"
	"github.com/auvn/go-examples/example1/s-tracking/driver"
	"github.com/auvn/go-examples/example1/s-tracking/handler"
	"github.com/auvn/go-examples/example1/s-trips/tripsevent"
)

func main() {
	drivers := driver.NewDrivers()
	events := transportutil.NewPublisher(
		natsss.NewClient(natsss.ClientConfig{
			Name:        "stracking",
			ClusterName: "test-cluster",
		}))
	handlers := handler.Handlers{
		Events:  events,
		Drivers: drivers,
	}
	httpServer := hottabych.
		Handle("Track", handlers.Track)

	dispatchers := dispatcher.Dispatchers{
		Events:  events,
		Drivers: drivers,
	}
	natsssServer := natsss.NewServer(
		natsss.ServerConfig{ClusterName: "test-cluster", Name: "stracking"})

	natsssServer.
		Subscribe(tripsevent.TypeReserved, dispatchers.FindDriver).
		Subscribe(tripsevent.TypeCompleted, dispatchers.UnlockDriver)

	service.Serve(context.Background(),
		httpServer,
		natsssServer)
}
