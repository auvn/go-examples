package trackingevent

import "github.com/auvn/go-examples/example1/s-framework/builtin/id"

const (
	TypeDriverFound = "tracking_driver_found"
)

type DriverFound struct {
	DriverID id.ID
	TripID   id.ID
}
