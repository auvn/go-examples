package web

import (
	"context"
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
			client = hottabych.NewClient()
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
		defer r.Body.Close()

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

		if reply.Error != nil {
			rw.WriteHeader(http.StatusBadRequest)
			if _, err := rw.Write(reply.Error); err != nil {
				log.Printf("cannot write error: %v\n", err)
			}
			return
		}

		if _, err := io.Copy(rw, reply.Body); err != nil {
			log.Printf("cannot copy reply body to response: %v\n", err)
		}
	}
}
