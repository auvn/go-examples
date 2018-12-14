package natsss

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/go-nats"
	"github.com/nats-io/go-nats-streaming"
	"github.com/nats-io/nuid"
	"github.com/pkg/errors"
)

var (
	clientSuffix = nuid.Next()
)

func clientID(name, suffix string) string {
	return fmt.Sprintf("%s-%s-%s", name, suffix, clientSuffix)
}

func connectNats(url, clientID string) *nats.Conn {
	options := func(o *nats.Options) error {
		cb := func(prefix string) nats.ConnHandler {
			return func(conn *nats.Conn) {
				log.Printf("natsss: %s: %s", prefix, clientID)
			}
		}
		o.DisconnectedCB = cb("disconnected")
		o.ClosedCB = cb("closed")
		o.DiscoveredServersCB = cb("discovered servers")
		o.ReconnectedCB = cb("reconnected")
		o.MaxReconnect = -1
		o.PingInterval = 5 * time.Second
		return nil
	}

	nc, err := nats.Connect(url, options, nats.Name(clientID))
	if err != nil {
		log.Fatalf("natsss: failed to connect to nats: %v", err)
	}
	return nc
}

func connect(url, cluster, name, suffix string) stan.Conn {
	if url == "" {
		url = nats.DefaultURL
	}
	clientIDstr := clientID(name, suffix)
	natsConn := connectNats(url, clientIDstr)
	conn, err := stan.Connect(
		cluster,
		natsConn.Opts.Name,
		stan.Pings(2, 3),
		stan.NatsConn(natsConn))
	if err != nil {
		log.Fatalln(errors.Wrap(err, "natsss: failed to connect to stan"))
	}
	return conn
}
