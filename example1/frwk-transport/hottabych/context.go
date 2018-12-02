package hottabych

import (
	"context"
	"net/http"
	"time"
)

const (
	headerMessageID       = "X-Message-ID"
	headerMessageType     = "X-Message-Type"
	headerMessageDeadline = "X-Message-Deadline"
	headerReplyError      = "X-Reply-Error"
)

func contextFromRequest(req *http.Request) context.Context {
	return nil
}

func contextToRequest(ctx context.Context, req *http.Request) {
	if deadline, ok := ctx.Deadline(); ok {
		req.Header.Set(headerMessageDeadline, deadline.Format(time.RFC3339Nano))
	}

}

func contextToResponse(ctx context.Context, rw http.ResponseWriter) {

}
