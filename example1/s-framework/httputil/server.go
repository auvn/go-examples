package httputil

import (
	"context"
	"log"
	"net/http"
	"time"
)

type ServeConfig struct {
	Addr            string
	ShutdownTimeout time.Duration
}

func Serve(ctx context.Context, handler http.Handler, cfg ServeConfig) error {
	errCh := make(chan error, 1)

	s := http.Server{
		Addr:    cfg.Addr,
		Handler: handler,
	}

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			errCh <- err
		}
	}()

	go func() {
		errCh <- s.ListenAndServe()
	}()

	log.Println("http: serving")
	return <-errCh
}
