package main

import (
	"context"

	"github.com/auvn/go-examples/example1/frwk-core/servegroup"
	"github.com/auvn/go-examples/example1/frwk-core/transport/event/eventutil"
	"github.com/auvn/go-examples/example1/frwk-transport/natsss"
	"github.com/auvn/go-examples/example1/s-calculations/handler"
	"github.com/auvn/go-examples/example1/s-history/historyevent"
)

func main() {
	server := natsss.NewStreams(natsss.StreamConfig{Name: "scalculations"})
	client := natsss.NewClient(natsss.ClientConfig{Name: "scalculations"})

	handlers := handler.Handlers{
		Events: eventutil.NewPublisher(client),
	}

	server.Subscribe(historyevent.TypeTripAdded, handlers.CalculateTrip)

	servegroup.Serve(context.Background(), server)
}
