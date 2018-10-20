package main

import (
	"context"

	"github.com/auvn/go-examples/example1/s-framework/service"
	"github.com/auvn/go-examples/example1/s-framework/transport/hottabych"
	"github.com/auvn/go-examples/example1/s-users/handler"
)

func main() {

	h := handler.Handlers{}
	server := hottabych.
		Handle("AuthenticateDriver", h.AuthenticateDriver).
		Handle("AuthenticateRider", h.AuthenticateRider).
		Handle("RestoreDriverState", h.RestoreDriverState).
		Handle("RestoreRiderState", h.RestoreRiderState)

	service.Serve(context.Background(), server)
}
