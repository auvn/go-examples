package kafkaz

import (
	"bytes"
	"context"
	"log"

	"github.com/Shopify/sarama"
	"github.com/auvn/go-examples/example1/s-framework/transport/event"
)

type ServerConfig struct {
	GroupID string
	Addrs   []string
}
type Server struct {
	m        event.Dispatcher
	consumer sarama.ConsumerGroup
}

func (s *Server) Subscribe(msgType string, handler event.HandlerFunc) *Server {
	s.m.Handle(msgType, handler)
	return s
}

func (s Server) Serve(ctx context.Context) error {
	topics := make([]string, 0, len(s.m))
	for msgType := range s.m {
		topics = append(topics, msgType)
	}
	return s.consumer.Consume(ctx, topics, consumerGroupHandler{s.m})
}

func NewServer(cfg ServerConfig) *Server {
	consumerCfg := sarama.NewConfig()
	consumerCfg.Version = sarama.V0_10_2_0
	consumer, err := sarama.NewConsumerGroup(cfg.Addrs, cfg.GroupID, consumerCfg)
	if err != nil {
		log.Fatalf("kafka consumer: %v\n", err)
	}
	return &Server{
		consumer: consumer,
		m:        event.Dispatcher{},
	}
}

type consumerGroupHandler struct {
	m event.Dispatcher
}

func (consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (c consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		dispatcher, ok := c.m[msg.Topic]
		if !ok {
			log.Printf("consumer: unknown dispatcher for topic %q\n", msg.Topic)
			continue
		}
		if err := dispatcher(session.Context(), bytes.NewBuffer(msg.Value)); err != nil {
			log.Printf("%s consumer: failed to dispatch message: %v\n", msg.Topic, err)
			continue
		}
		session.MarkMessage(msg, "")
	}
	return nil
}
