package natsss

import (
	"bytes"
	"context"
	"log"

	"github.com/auvn/go-examples/example1/frwk-core/transport/event"
	"github.com/nats-io/go-nats-streaming"
	"github.com/pkg/errors"
)

type ServerConfig struct {
	ClusterName string
	Name        string
	URL         string
}

type Server struct {
	cfg           ServerConfig
	subscriptions event.Dispatcher
	conn          stan.Conn
}

func (s *Server) Subscribe(eventType string, handler event.HandlerFunc) *Server {
	s.subscriptions.Handle(eventType, handler)
	return s
}

func (s *Server) Serve(ctx context.Context) error {
	for eventType, dispatcher := range s.subscriptions {
		subscr := subscription{
			eventType:  eventType,
			dispatcher: dispatcher,
			ctx:        ctx,
		}

		_, err := s.conn.QueueSubscribe(
			eventType,
			eventType,
			subscr.HandleMessage,
			stan.DurableName(s.cfg.Name),
			stan.SetManualAckMode())
		if err != nil {
			return errors.Wrapf(err, "failed to subscribe to %q", eventType)
		}
	}

	errCh := make(chan error)
	go func() {
		<-ctx.Done()
		errCh <- s.conn.Close()
	}()

	log.Println("natsss: serving")
	return <-errCh
}

func NewServer(cfg ServerConfig) *Server {
	cfg.ClusterName = "test-cluster"
	return &Server{
		cfg:           cfg,
		conn:          connect(cfg.ClusterName, cfg.Name, "server"),
		subscriptions: event.Dispatcher{},
	}
}

type subscription struct {
	eventType  string
	dispatcher event.HandlerFunc
	ctx        context.Context
}

func (s *subscription) HandleMessage(msg *stan.Msg) {
	if err := s.dispatcher(s.ctx, bytes.NewBuffer(msg.Data)); err != nil {
		log.Printf("natsss: failed to handle message %q: %v", s.eventType, err)
		return
	}

	if err := msg.Ack(); err != nil {
		log.Printf("natsss: failed to ack message %q: %v", s.eventType, err)
		return
	}
}
