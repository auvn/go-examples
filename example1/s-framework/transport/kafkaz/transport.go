package kafkaz

import (
	"context"

	"github.com/auvn/go-examples/example1/s-framework/transport"
)

type Client struct {
}

func (c Client) Deliver(ctx context.Context, msg transport.Message) error {
	return nil
}

type Server struct {
	m transport.DispatcherMap
}

func (s *Server) Subscribe(msgType string, dispatcher transport.DispatcherFunc) *Server {
	s.m.Register(msgType, dispatcher)
	return s
}

func (s Server) Serve(ctx context.Context) error {
	<-ctx.Done()
	return nil
}

func NewServer() *Server {
	return &Server{
		m: transport.DispatcherMap{},
	}
}
