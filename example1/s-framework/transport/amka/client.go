package amka

import (
	"context"
	"io/ioutil"

	"github.com/auvn/go-examples/example1/s-framework/transport"
	"github.com/go-stomp/stomp"
	"github.com/pkg/errors"
)

var (
	headerMessageID = "x-message-id"
)

type ClientConfig struct {
	Addr string
}

type Client struct {
	client *stomp.Conn
}

func (c *Client) Publish(ctx context.Context, event transport.Event) error {
	bb, err := ioutil.ReadAll(event.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read event body")
	}

	eventIDHeader := stomp.SendOpt.Header(headerMessageID, string(event.ID))

	err = c.client.Send(destinationTopic(event.Type), "application/octet", bb,
		stomp.SendOpt.Receipt,
		eventIDHeader)
	if err != nil {
		return errors.Wrap(err, "failed to send event")
	}
	return nil
}

func (c *Client) Shutdown(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func() {
		errCh <- c.client.Disconnect()
	}()

	go func() {
		<-ctx.Done()
		errCh <- ctx.Err()
	}()

	return <-errCh
}

func NewClient(cfg ClientConfig) *Client {
	return &Client{
		client: dial(cfg.Addr),
	}
}
