package main

import (
	"context"

	"github.com/auvn/go-examples/example1/frwk-core/servegroup"
	"github.com/auvn/go-examples/example1/frwk-core/transport/event/eventutil"
	"github.com/auvn/go-examples/example1/frwk-storage/nosql/monga"
	"github.com/auvn/go-examples/example1/frwk-transport/hottabych"
	"github.com/auvn/go-examples/example1/frwk-transport/natsss"
	"github.com/auvn/go-examples/example1/s-calculations/calculationsevent"
	"github.com/auvn/go-examples/example1/s-history/handler"
	"github.com/auvn/go-examples/example1/s-history/history"
	"github.com/auvn/go-examples/example1/s-trips/event/tripsevent"
)

func main() {
	natsssStreams := natsss.NewStreams(natsss.EnvStreamConfig())
	mongoClient := monga.MustNew(monga.EnvConfig())
	handlers := handler.Handlers{
		History: history.New(mongoClient),
		Events:  eventutil.NewPublisher(natsss.NewClient(natsss.EnvClientConfig())),
	}

	httpServer := hottabych.
		Handle("Get", handlers.Get)

	natsssStreams.
		Subscribe(tripsevent.TypeCompleted, handlers.Save).
		Subscribe(calculationsevent.TypeBreakdownCalculated, handlers.SaveBreakdown)

	servegroup.Serve(context.Background(), natsssStreams, httpServer)
}
