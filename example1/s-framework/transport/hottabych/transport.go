package hottabych

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/auvn/go-examples/example1/s-framework/transport"
)

const (
	headerMessageID   = "X-Message-ID"
	headerMessageType = "X-Message-Type"
	headerReplyError  = "X-Reply-Error"
)

type Client struct {
	Port       int
	httpClient http.Client
}

func (c Client) Request(ctx context.Context, msg transport.Message) (*transport.Reply, error) {
	if c.Port == 0 {
		c.Port = 1200
	}
	recipientURL := fmt.Sprintf("http://%s:%d", msg.Recipient, c.Port)
	req, err := http.NewRequest(http.MethodPost, recipientURL, msg.Body)
	if err != nil {
		return nil, err
	}

	req.Header.Set(headerMessageID, string(msg.ID))
	req.Header.Set(headerMessageType, msg.Type)

	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return &transport.Reply{
		Body:       resp.Body,
		Error:      resp.Header.Get(headerReplyError),
		Successful: resp.StatusCode == http.StatusOK,
	}, nil
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: 2 * time.Second,
		},
	}
}
