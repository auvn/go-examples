package history

import (
	"context"
	"time"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
)

type Record struct {
	RideID   id.ID
	RiderID  id.ID
	DriverID id.ID
	Duration time.Duration
	Distance int
}
type History struct {
}

func (h *History) Save(ctx context.Context, r Record) error {
	return nil
}

func (h *History) ByRiderID(ctx context.Context, id id.ID) (int, error) {
	return 0, nil
}
