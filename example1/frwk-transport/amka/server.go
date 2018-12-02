package amka

import (
	"bytes"
	"context"
	"log"

	"github.com/auvn/go-examples/example1/s-framework/transport/event"
	"github.com/go-stomp/stomp"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type ServerConfig struct {
	Name string
	Addr string
}

type Server struct {
	name          string
	subscriptions event.Dispatcher
	conn          *stomp.Conn
}

func (s *Server) Subscribe(eventType string, handler event.HandlerFunc) *Server {
	s.subscriptions.Handle(eventType, handler)
	return s
}

func (s *Server) Serve(ctx context.Context) error {
	errGroup, ctx := errgroup.WithContext(ctx)

	for eventType, subscr := range s.subscriptions {
		destination := destinationTopic(eventType)
		subscription, err := s.conn.Subscribe(destination,
			stomp.AckClientIndividual,
			stomp.SubscribeOpt.Id(s.name))
		if err != nil {
			return errors.Wrap(err, "amka: failed to subscribe")
		}

		disp := subscription{
			subscription: subscription,
			handler:      subscr,
		}

		errGroup.Go(func() error {
			return disp.Serve(ctx)
		})

	}
	log.Println("amka: serving")
	return errGroup.Wait()
}

func NewServer(cfg ServerConfig) *Server {
	return &Server{
		name:          cfg.Name,
		subscriptions: event.Dispatcher{},
		conn:          dial(cfg.Addr),
	}
}

type subscription struct {
	subscription *stomp.Subscription
	handler      event.HandlerFunc
}

func (d *subscription) Serve(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		msg, err := d.subscription.Read()
		if err != nil {
			log.Printf("amka: failed to read message: %v\n", err)
			continue
		}

		body := bytes.NewBuffer(msg.Body)

		if err := d.handler(ctx, body); err != nil {
			log.Printf("amka: failed to handle message: %v\n", err)
			continue
		}

		if err := msg.Conn.Ack(msg); err != nil {
			log.Printf("amka: failed to send ask: %v\n", err)
			continue
		}
	}
}
