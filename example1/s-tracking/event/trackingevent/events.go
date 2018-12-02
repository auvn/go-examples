package trackingevent

import "github.com/auvn/go-examples/example1/frwk-core/builtin/id"

const (
	TypeTripDriverSuggested = "tracking_suggested_trip_driver"
)

type TripDriverSuggested struct {
	DriverID id.ID
	TripID   id.ID
}
