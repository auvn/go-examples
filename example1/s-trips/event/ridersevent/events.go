package ridersevent

import "github.com/auvn/go-examples/example1/frwk-core/builtin/id"

const (
	TypeTripReserved = "trip_reserved"
	TypeDriverFound  = "driver_found"
)

type TripReserved struct {
	TripID id.ID
}

type DriverFound struct {
	TripID   id.ID
	DriverID id.ID
}
