package event

import "github.com/auvn/go-examples/example1/frwk-core/builtin/id"

const (
	TypeTripAdded = "history_trip_added"
)

type TripAdded struct {
	TripID id.ID
}
