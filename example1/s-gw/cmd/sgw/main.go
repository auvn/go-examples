package main

import (
	"context"

	"github.com/auvn/go-examples/example1/frwk-core/servegroup"
	"github.com/auvn/go-examples/example1/frwk-transport/natsss"
	"github.com/auvn/go-examples/example1/s-gw/gwevent"
	"github.com/auvn/go-examples/example1/s-gw/stream"
	"github.com/auvn/go-examples/example1/s-gw/web"
)

func main() {
	streams := stream.NewStreams(":8081")
	server := web.NewServer(":8080",
		web.EndpointConfig{
			Path:          "/trips/reserve",
			TargetService: "strips",
			MessageType:   "Reserve",
			Method:        "POST",
		},
		web.EndpointConfig{
			Path:          "/trips/complete",
			TargetService: "strips",
			MessageType:   "Complete",
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
			Path:          "/driver/auth",
			TargetService: "susers",
			MessageType:   "AuthenticateDriver",
			Method:        "POST",
		},
	)

	natsssServer := natsss.NewStreams(natsss.EnvStreamConfig())
	natsssServer.Subscribe(gwevent.TypeUserEvent, streams.SendUserEvent)

	servegroup.Serve(context.Background(),
		server,
		streams,
		natsssServer)
}
