package main

import (
	"context"

	"github.com/auvn/go-examples/example1/s-framework/database/sqldb"
	"github.com/auvn/go-examples/example1/s-framework/service"
	"github.com/auvn/go-examples/example1/s-framework/transport/hottabych"
	"github.com/auvn/go-examples/example1/s-trips/handler"
	"github.com/auvn/go-examples/example1/s-trips/trip"
)

func main() {
	db := sqldb.NewPostgres("strips")

	handlers := handler.Handlers{
		Trips: trip.NewTrips(db),
	}

	httpTransport := hottabych.
		Handle("Reserve", handlers.Reserve)
	//Handle("Update", handlers.Request)

	service.Serve(context.Background(), httpTransport)
}
