package tripsevent

import "github.com/auvn/go-examples/example1/s-framework/builtin/id"

const (
	TypeReserved    = "trips/reserved"
	TypeDriverFound = "trips/driver_found"
	TypeCanceled    = "trips/canceled"
)

type Reserved struct {
	TripID  id.ID
	RiderID id.ID
}

type DriverFound struct {
	TripID   id.ID
	DriverID id.ID
}

type Canceled struct {
	TripID id.ID
}
