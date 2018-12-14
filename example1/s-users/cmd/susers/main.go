package main

import (
	"context"

	"github.com/auvn/go-examples/example1/frwk-core/service"
	"github.com/auvn/go-examples/example1/frwk-transport/hottabych"
	"github.com/auvn/go-examples/example1/frwk-transport/natsss"
	"github.com/auvn/go-examples/example1/s-users/handler"
	"github.com/auvn/go-examples/example1/s-users/usersevent"
)

func main() {

	h := handler.Handlers{
		Events: &usersevent.Publisher{
			Publisher: natsss.NewClient(
				natsss.ClientConfig{ClusterName: "test-cluster", Name: "susers"}),
		},
	}
	server := hottabych.
		Handle("AuthenticateDriver", h.AuthenticateDriver).
		Handle("AuthenticateRider", h.AuthenticateRider).
		Handle("RestoreDriverState", h.RestoreDriverState).
		Handle("RestoreRiderState", h.RestoreRiderState)

	service.Serve(context.Background(), server)
}
