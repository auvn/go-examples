package handler

import (
	"github.com/auvn/go-examples/example1/frwk-core/transport/event/eventutil"
	"github.com/auvn/go-examples/example1/s-tracking/driver"
	"github.com/auvn/go-examples/example1/s-tracking/event/driversevent"
)

type Handlers struct {
	Drivers          *driver.Drivers
	Events           *eventutil.Publisher
	DriverHeartbeats *driversevent.Heartbeats
}
