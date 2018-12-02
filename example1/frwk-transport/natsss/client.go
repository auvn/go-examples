package natsss

import (
	"context"
	"io/ioutil"

	"github.com/auvn/go-examples/example1/frwk-core/transport"
	"github.com/nats-io/go-nats-streaming"
	"github.com/pkg/errors"
)

type ClientConfig struct {
	Name        string
	ClusterName string
	URL         string
}

type Client struct {
	cfg  ClientConfig
	conn stan.Conn
}

func (c *Client) Publish(ctx context.Context, event transport.Event) error {
	bb, err := ioutil.ReadAll(event.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read body")
	}

	if err := c.conn.Publish(event.Type, bb); err != nil {
		return errors.Wrap(err, "failed to publish body")
	}
	return nil
}

func NewClient(cfg ClientConfig) *Client {
	return &Client{
		cfg:  cfg,
		conn: connect("test-cluster", cfg.Name, "client"),
	}
}
