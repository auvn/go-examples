package natsss

import (
	"fmt"
	"log"

	"github.com/nats-io/go-nats-streaming"
	"github.com/nats-io/nuid"
)

var (
	clientSuffix = nuid.Next()
)

func clientID(name, suffix string) string {
	return fmt.Sprintf("%s-%s-%s", name, suffix, clientSuffix)
}

func connect(cluster, name, suffix string) stan.Conn {
	conn, err := stan.Connect(cluster, clientID(name, suffix))
	if err != nil {
		log.Fatalf("natsss: failed to connect: %v", err)
	}
	return conn
}
