package driversevent

import "github.com/auvn/go-examples/example1/frwk-core/builtin/id"

const (
	TypeRestoredTrip = "restored_trip"
)

type RestoredTrip struct {
	TripID *id.ID
}
