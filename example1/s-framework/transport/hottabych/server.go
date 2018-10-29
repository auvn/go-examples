package hottabych

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/auvn/go-examples/example1/s-framework/apierror"
	"github.com/auvn/go-examples/example1/s-framework/httputil"
	"github.com/auvn/go-examples/example1/s-framework/transport"
)

type Server struct {
	addr     string
	handlers transport.HandlerMap
}

func (s Server) handler(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	ctx := req.Context()

	messageType := req.Header.Get(headerMessageType)
	h, ok := s.handlers[messageType]
	if !ok {
		log.Printf("unregistered mesage type: %s", messageType)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var handlerResp bytes.Buffer
	if err := h(ctx, req.Body, &handlerResp); err != nil {
		handleError(err, rw)
		log.Printf("hottabych: %q: handle error: %v", messageType, err)
	}

	if _, err := io.Copy(rw, &handlerResp); err != nil {
		log.Printf("cannot copy handler response: %v", err)
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
	if addr == "" {
		addr = options.Addr
	}
	return &Server{
		addr:     addr,
		handlers: transport.HandlerMap{},
	}
}

func Handle(msgType string, h transport.HandlerFunc) *Server {
	return DefaultServer.Handle(msgType, h)
}

func httpCodeOf(code string) int {
	switch code {
	case apierror.CodeInternal:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
func errorCodeOf(err error) string {
	switch err := err.(type) {
	case apierror.Error:
		return err.Code
	}
	return apierror.CodeInternal
}

func handleError(err error, rw http.ResponseWriter) {
	errCode := errorCodeOf(err)
	httpCode := httpCodeOf(errCode)

	rw.Header().Set(headerReplyError, errCode)
	rw.WriteHeader(httpCode)
}
