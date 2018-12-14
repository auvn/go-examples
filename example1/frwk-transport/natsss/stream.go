package natsss

import (
	"bytes"
	"context"
	"log"
	"os"

	"github.com/auvn/go-examples/example1/frwk-core/service"
	"github.com/auvn/go-examples/example1/frwk-core/transport/event"
	"github.com/nats-io/go-nats-streaming"
	"github.com/pkg/errors"
)

type StreamConfig struct {
	ClusterName string
	Name        string
	URL         string
}

type Streams struct {
	cfg           StreamConfig
	subscriptions event.Dispatcher
	conn          stan.Conn
}

func (s *Streams) Subscribe(eventType string, handler event.HandlerFunc) *Streams {
	s.subscriptions.Handle(eventType, handler)
	return s
}

func (s *Streams) Serve(ctx context.Context) error {
	for eventType, dispatcher := range s.subscriptions {
		subscr := eventsSubscription{
			eventType:  eventType,
			dispatcher: dispatcher,
			// TODO pass something via payload and build context here
			ctx: ctx,
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

func NewStreams(cfg StreamConfig) *Streams {
	cfg.ClusterName = "test-cluster"
	return &Streams{
		cfg:           cfg,
		conn:          connect(cfg.URL, cfg.ClusterName, cfg.Name, "server"),
		subscriptions: event.Dispatcher{},
	}
}

func EnvURL() string {
	return os.Getenv("NATSSS_URL")
}
func EnvStreamConfig() StreamConfig {
	return StreamConfig{
		Name: service.EnvName(),
		URL:  EnvURL(),
	}

}

type eventsSubscription struct {
	eventType  string
	dispatcher event.HandlerFunc
	ctx        context.Context
}

func (s *eventsSubscription) HandleMessage(msg *stan.Msg) {
	if err := s.dispatcher(s.ctx, bytes.NewBuffer(msg.Data)); err != nil {
		log.Printf("natsss: failed to handle message %q: %v", s.eventType, err)
		return
	}

	if err := msg.Ack(); err != nil {
		log.Printf("natsss: failed to ack message %q: %v", s.eventType, err)
		return
	}
}
