package tripsevent

import "github.com/auvn/go-examples/example1/s-framework/builtin/id"

type TripReserved struct {
	ID      id.ID
	RiderID id.ID
}

type TripDriverFound struct {
	TripID   id.ID
	DriverID id.ID
}
