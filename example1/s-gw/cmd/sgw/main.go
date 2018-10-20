package main

import (
	"context"

	"github.com/auvn/go-examples/example1/s-framework/service"
	"github.com/auvn/go-examples/example1/s-gw/web"
)

func main() {
	server := web.NewServer(":8080",
		web.EndpointConfig{
			Path:          "/trips/request",
			TargetService: "strips",
			MessageType:   "Request",
			Method:        "POST",
		},
		web.EndpointConfig{
			Path:          "/trips/update",
			TargetService: "strips",
			MessageType:   "Update",
			Method:        "POST",
		},

		web.EndpointConfig{
			Path:          "/history/get",
			TargetService: "shistory",
			MessageType:   "Get",
			Method:        "GET",
		},

		web.EndpointConfig{
			Path:          "/tracking/track",
			TargetService: "stracking",
			MessageType:   "Track",
			Method:        "POST",
		},

		web.EndpointConfig{
			Path:          "/driver/login",
			TargetService: "susers",
			MessageType:   "AuthenticateDriver",
			Method:        "POST",
		},
		web.EndpointConfig{
			Path:          "/rider/login",
			TargetService: "susers",
			MessageType:   "AuthenticateRider",
			Method:        "POST",
		},

		web.EndpointConfig{
			Path:          "/service/health",
			TargetService: "shealth",
			MessageType:   "Get",
			Method:        "GET",
		},
	)

	service.Serve(context.Background(), server)
}
