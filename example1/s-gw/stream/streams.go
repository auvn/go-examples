package stream

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
	"github.com/auvn/go-examples/example1/s-framework/encoding"
	"github.com/auvn/go-examples/example1/s-framework/httputil"
)

type StreamConfig struct {
	Path string
}

type Streams struct {
	addr  string
	m     sync.RWMutex
	users map[id.ID]*stream
}

func (ss *Streams) addUserStream(id id.ID, s *stream) bool {
	ss.m.Lock()
	defer ss.m.Unlock()
	if _, ok := ss.users[id]; ok {
		return false
	}
	ss.users[id] = s
	return true
}

func (ss *Streams) removeUserStream(id id.ID) {
	ss.m.Lock()
	defer ss.m.Unlock()

	delete(ss.users, id)
}

func (ss *Streams) lookupStream(id id.ID) (*stream, bool) {
	ss.m.RLock()
	defer ss.m.RUnlock()

	s, ok := ss.users[id]
	return s, ok
}

type subscribeRequest struct {
	UserID id.ID
}

func (s *Streams) handle(rw http.ResponseWriter, req *http.Request) {
	var subscr subscribeRequest
	if err := encoding.UnmarshalReader(req.Body, &subscr); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	flusher, ok := rw.(http.Flusher)
	if !ok {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	userStream := stream{
		w:       rw,
		flusher: flusher,
	}

	if ok := s.addUserStream(subscr.UserID, &userStream); !ok {
		rw.WriteHeader(http.StatusConflict)
		return
	}

	defer s.removeUserStream(subscr.UserID)

	if err := userStream.Serve(req.Context()); err != nil {
		log.Printf("http streams: %q: %v", subscr.UserID, err)
		return
	}
}

type userEvent struct {
	UserID id.ID
	Type   string
	Body   encoding.RawMessage
}

func (s *Streams) SendUserEvent(ctx context.Context, body io.Reader) error {
	var event userEvent
	if err := encoding.UnmarshalReader(body, &event); err != nil {
		return err
	}

	stream, ok := s.lookupStream(event.UserID)
	if !ok {
		log.Printf("stream %q does not exist", event.UserID)
		return nil
	}

	if err := stream.Send(ctx, streamEvent{
		Type: event.Type,
		Body: event.Body,
	}); err != nil {
		return err
	}

	return nil
}

func (s *Streams) Serve(ctx context.Context) error {
	return httputil.Serve(ctx, http.HandlerFunc(s.handle),
		httputil.ServeConfig{Addr: s.addr, ShutdownTimeout: time.Minute})
}

func NewStreams(addr string) *Streams {
	return &Streams{
		addr:  addr,
		users: make(map[id.ID]*stream),
	}
}

type stream struct {
	w       io.Writer
	flusher http.Flusher
}
type streamEvent struct {
	Type string
	Body encoding.RawMessage
}

func (s *stream) Send(ctx context.Context, event streamEvent) error {
	if err := encoding.MarshalToWriter(event, s.w); err != nil {
		return err
	}

	if _, err := s.flush(); err != nil {
		return err
	}
	return nil
}

func (s *stream) flush() (n int, err error) {
	defer s.flusher.Flush()
	return fmt.Fprint(s.w, "\r\n")
}

func (s *stream) Serve(ctx context.Context) error {
	for {
		if _, err := s.flush(); err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(3 * time.Second):
			continue
		}
	}
	return nil
}
