package dispatcher

import (
	"context"
	"io"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
	"github.com/auvn/go-examples/example1/s-framework/encoding"
)

func (d *Dispatchers) UnlockDriver(ctx context.Context, body io.Reader) error {
	var completedTrip struct {
		DriverID id.ID
	}
	if err := encoding.UnmarshalReader(body, &completedTrip); err != nil {
		return err
	}

	if err := d.Drivers.Unlock(ctx, completedTrip.DriverID); err != nil {
		return err
	}
	return nil
}
