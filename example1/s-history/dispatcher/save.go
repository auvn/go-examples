package dispatcher

import (
	"context"
	"io"
	"time"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
)

type TripCompleted struct {
	TripID   id.ID
	RiderID  id.ID
	DriverID id.ID
	Distance int
	Duration time.Duration
}

func (d *Dispatchers) Save(ctx context.Context, body io.Reader) error {

	return nil
}
