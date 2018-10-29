package driversevent

import "github.com/auvn/go-examples/example1/s-framework/builtin/id"

const (
	TypeAssignedTrip = "assigned_trip"
	TypeHeartbeat    = "heartbeat"
)

type Heartbeat struct{}

type AssignedTrip struct {
	TripID id.ID
}
