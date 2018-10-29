package dispatcher

import (
	"github.com/auvn/go-examples/example1/s-framework/transport/transportutil"
	"github.com/auvn/go-examples/example1/s-tracking/driver"
)

type Dispatchers struct {
	Events  *transportutil.Publisher
	Drivers *driver.Drivers
}
