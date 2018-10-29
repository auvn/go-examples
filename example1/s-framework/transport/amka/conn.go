package amka

import (
	"fmt"
	"log"

	"github.com/go-stomp/stomp"
)

func dial(addr string) *stomp.Conn {
	conn, err := stomp.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("amka: failed to dial %q: %v", addr, err)
	}
	return conn
}

func destinationTopic(msgType string) string {
	return fmt.Sprintf("/topic/%s", msgType)
}
