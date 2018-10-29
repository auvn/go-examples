package web

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
	"github.com/auvn/go-examples/example1/s-framework/httputil"
	"github.com/auvn/go-examples/example1/s-framework/transport"
	"github.com/auvn/go-examples/example1/s-framework/transport/hottabych"
	"github.com/gorilla/mux"
)

type EndpointConfig struct {
	Method string
	Path   string

	TargetService string
	MessageType   string

	Port int
}

type Server struct {
	h    http.Handler
	addr string
}

func (s *Server) Serve(ctx context.Context) error {
	return httputil.Serve(ctx,
		s.h,
		httputil.ServeConfig{Addr: s.addr, ShutdownTimeout: time.Minute})
}

func NewServer(addr string, endpoints ...EndpointConfig) *Server {
	clients := map[string]transport.Requester{}

	r := mux.NewRouter()
	for _, e := range endpoints {
		client, ok := clients[e.TargetService]
		if !ok {
			httpClient := hottabych.NewClient()
			httpClient.Port = e.Port
			client = httpClient
			clients[e.TargetService] = client
		}
		r.
			Handle(e.Path, newHandler(e, client)).
			Methods(e.Method)
	}

	return &Server{
		h:    r,
		addr: addr,
	}
}

const (
	headerMessageID = "X-Message-ID"
)

func newHandler(cfg EndpointConfig, client transport.Requester) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		messageID := id.New()
		rw.Header().Set(headerMessageID, string(messageID))

		msg := transport.Message{
			ID:        messageID,
			Body:      r.Body,
			Recipient: cfg.TargetService,
			Type:      cfg.MessageType,
		}

		log.Printf("web: handling %+v\n", msg)

		reply, err := client.Request(r.Context(), msg)
		if err != nil {
			log.Printf("failed to handle request %q -> %q: %v\n", cfg.TargetService, cfg.MessageType, err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer reply.Body.Close()

		if err := handleReply(*reply, rw); err != nil {
			log.Printf("failed to handle reply: %v", err)
		}
	}
}

type response struct {
	Body  *json.RawMessage `json:",omitempty"`
	Error string           `json:",omitempty"`
}

func handleReply(reply transport.Reply, rw http.ResponseWriter) error {
	rw.Header().Set("Content-Type", "application/json")

	if !reply.Successful {
		rw.WriteHeader(http.StatusBadRequest)
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, reply.Body); err != nil {
		return err
	}

	resp := response{
		Error: reply.Error,
	}
	if buf.Len() != 0 {
		body := json.RawMessage(buf.Bytes())
		resp.Body = &body
	}

	enc := json.NewEncoder(rw)
	return enc.Encode(resp)
}
