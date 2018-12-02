package servegroup

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func onSigkill(fn func()) {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		<-ch
		fn()
	}()
}

type Server interface {
	Serve(ctx context.Context) error
}

func Serve(ctx context.Context, ss ...Server) {

	ctx, cancel := context.WithCancel(ctx)
	onSigkill(cancel)

	group, ctx := errgroup.WithContext(ctx)
	for _, s := range ss {
		s := s
		group.Go(func() error {
			return s.Serve(ctx)
		})
	}

	log.Printf("serve group: %v\n", group.Wait())
}
