package main

import (
	"context"

	"github.com/auvn/go-examples/example1/frwk-core/servegroup"
	"github.com/auvn/go-examples/example1/frwk-core/transport/event/eventutil"
	"github.com/auvn/go-examples/example1/frwk-storage/nosql/monga"
	"github.com/auvn/go-examples/example1/frwk-transport/hottabych"
	"github.com/auvn/go-examples/example1/frwk-transport/natsss"
	"github.com/auvn/go-examples/example1/s-history/handler"
	"github.com/auvn/go-examples/example1/s-history/history"
	"github.com/auvn/go-examples/example1/s-trips/event/tripsevent"
)

func main() {
	natsssStreams := natsss.
		NewStreams(natsss.StreamConfig{ClusterName: "test-cluster", Name: "shistory"})

	mongoClient := monga.MustNew(monga.Config{
		Name:  "shistory",
		Hosts: []string{"localhost:27017"},
	})
	handlers := handler.Handlers{
		History: history.New(mongoClient),
		Events:  eventutil.NewPublisher(natsss.NewClient(natsss.ClientConfig{Name: "shistory"})),
	}

	httpServer := hottabych.
		Handle("Get", handlers.Get)

	natsssStreams.Subscribe(tripsevent.TypeCompleted, handlers.Save)

	servegroup.Serve(context.Background(), natsssStreams, httpServer)
}
