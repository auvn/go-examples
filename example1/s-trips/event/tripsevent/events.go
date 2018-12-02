package tripsevent

import (
	"time"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
)

const (
	TypeReserved  = "trips_reserved"
	TypeCompleted = "trips_completed"
)

type Reserved struct {
	TripID  id.ID
	RiderID id.ID
}

type Completed struct {
	TripID   id.ID
	DriverID id.ID
	RiderID  id.ID

	Distance int
	Duration time.Duration
}
