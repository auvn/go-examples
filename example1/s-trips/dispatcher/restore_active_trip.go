package dispatcher

import (
	"context"
	"fmt"
	"io"

	"github.com/auvn/go-examples/example1/s-framework/encoding"
	"github.com/auvn/go-examples/example1/s-users/usersevent"
)

func (d *Dispatchers) RestoreActiveTrip(ctx context.Context, body io.Reader) error {
	var event usersevent.DriverRestoredState
	if err := encoding.UnmarshalReader(body, &event); err != nil {
		return err
	}
	fmt.Println(event)
	return nil
}
