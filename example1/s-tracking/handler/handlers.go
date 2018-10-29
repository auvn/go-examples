package handler

import (
	"github.com/auvn/go-examples/example1/s-framework/transport/transportutil"
	"github.com/auvn/go-examples/example1/s-tracking/driver"
)

type Handlers struct {
	Drivers *driver.Drivers
	Events  *transportutil.Publisher
}
