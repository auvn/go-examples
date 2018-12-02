package main

import (
	"context"

	"github.com/auvn/go-examples/example1/frwk-transport/hottabych"
	"github.com/auvn/go-examples/example1/frwk-transport/natsss"
	"github.com/auvn/go-examples/example1/s-framework/servegroup"
	"github.com/auvn/go-examples/example1/s-history/dispatcher"
	"github.com/auvn/go-examples/example1/s-history/handler"
	"github.com/auvn/go-examples/example1/s-trips/event/tripsevent"
)

func main() {

	handlers := handler.Handlers{}

	httpServer := hottabych.Handle("Get", handlers.Get)

	dispatchers := dispatcher.Dispatchers{}

	server := natsss.NewServer(natsss.ServerConfig{ClusterName: "test-cluster", Name: "shistory"})
	server.Subscribe(tripsevent.TypeCompleted, dispatchers.Save)

	servegroup.Serve(context.Background(), server, httpServer)
}
