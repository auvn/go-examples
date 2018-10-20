package hottabych

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/auvn/go-examples/example1/s-framework/httputil"
	"github.com/auvn/go-examples/example1/s-framework/transport"
)

const (
	headerMessageID   = "X-Message-ID"
	headerMessageType = "X-Message-Type"
)

type Client struct {
	Port       int
	httpClient http.Client
}

func (c Client) Request(ctx context.Context, msg transport.Message) (*transport.Reply, error) {
	if msg.Recipient == "" {
		return nil, transport.ErrUnknownRecipient
	}

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

	var errCode []byte
	if resp.StatusCode != http.StatusOK {
		errCode = []byte(resp.Status)
	}

	return &transport.Reply{
		Body:  resp.Body,
		Error: errCode,
	}, nil

}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: 2 * time.Second,
		},
	}
}

type Server struct {
	addr     string
	handlers transport.HandlerMap
}

func (s Server) handler(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	//messageID := req.Header.Get(headerMessageID)
	messageType := req.Header.Get(headerMessageType)

	h, ok := s.handlers[messageType]
	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("unknown handler"))
		return
	}

	if err := h(req.Context(), req.Body, rw); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
}

func (s Server) Serve(ctx context.Context) error {
	return httputil.Serve(ctx,
		http.HandlerFunc(s.handler),
		httputil.ServeConfig{Addr: s.addr, ShutdownTimeout: time.Minute})
}

func (s *Server) Handle(msgType string, h transport.HandlerFunc) *Server {
	s.handlers.Register(msgType, h)
	return s
}

func NewServer(addr string) *Server {
	return &Server{
		addr:     addr,
		handlers: transport.HandlerMap{},
	}
}

var (
	DefaultServer = NewServer(":1200")
)

func Handle(msgType string, h transport.HandlerFunc) *Server {
	return DefaultServer.Handle(msgType, h)
}
