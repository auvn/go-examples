package main

import (
	"context"

	"github.com/auvn/go-examples/example1/s-framework/service"
	"github.com/auvn/go-examples/example1/s-framework/transport/kafkaz"
	"github.com/auvn/go-examples/example1/s-history/dispatcher"
)

func main() {
	s := service.Service{}

	dispatchers := dispatcher.Dispatchers{}

	kafkaSubscriber := kafkaz.NewServer().
		Subscribe("ride", dispatchers.DispatchRide).
		Subscribe("ridev2", dispatchers.DispatchRide)

	s.Serve(context.Background(), kafkaSubscriber)
}
