package hottabych

import "os"

var (
	options = struct {
		Addr string
	}{}

	DefaultServer *Server
)

func init() {
	options.Addr = os.Getenv("HOTTABYCH_ADDR")

	DefaultServer = NewServer("")
}
