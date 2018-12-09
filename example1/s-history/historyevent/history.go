package historyevent

import (
	"time"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
)

const (
	TypeTripAdded = "history_trip_added"
)

type TripAdded struct {
	TripID   id.ID
	Distance int
	Duration time.Duration
}
