package kafkaz

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/Shopify/sarama"
	"github.com/auvn/go-examples/example1/frwk-core/transport"
	"github.com/pkg/errors"
)

type ClientConfig struct {
	Addrs []string
}

type Client struct {
	producer sarama.SyncProducer
}

func (c *Client) Publish(ctx context.Context, event transport.Event) error {
	bb, err := ioutil.ReadAll(event.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read event body")
	}
	message := sarama.ProducerMessage{
		Key:   sarama.StringEncoder(string(event.ID)),
		Value: sarama.ByteEncoder(bb),
		Topic: event.Type,
	}
	if _, _, err := c.producer.SendMessage(&message); err != nil {
		return errors.Wrap(err, "failed to send message")
	}
	return nil
}

func (c *Client) Shutdown(ctx context.Context) error {
	return c.producer.Close()
}

func NewClient(cfg ClientConfig) *Client {
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(cfg.Addrs, producerConfig)
	if err != nil {
		log.Fatalf("kafka producer: %v\n", err)
	}

	return &Client{
		producer: producer,
	}
}
