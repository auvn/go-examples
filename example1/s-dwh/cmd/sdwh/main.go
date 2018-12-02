package main

import (
	"context"

	"github.com/auvn/go-examples/example1/frwk-transport/natsss"
	"github.com/auvn/go-examples/example1/s-framework/servegroup"
)

func main() {
	server := natsss.NewServer(natsss.ServerConfig{Name: "sdwh", ClusterName: "test-cluster"})
	servegroup.Serve(context.Background(), server)
}
