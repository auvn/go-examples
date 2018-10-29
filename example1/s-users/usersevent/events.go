package usersevent

import "github.com/auvn/go-examples/example1/s-framework/builtin/id"

const (
	TypeDriverRestoredState = "users_driver_restored_state"
)

type DriverRestoredState struct {
	DriverID id.ID
}
