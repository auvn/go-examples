package contextutil

import (
	"context"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
)

type key int

const (
	keyActivityID key = iota
)

func WithActivityID(ctx context.Context, messageID id.ID) context.Context {
	return context.WithValue(ctx, keyActivityID, messageID)
}

func ActivityID(ctx context.Context) id.ID {
	val, _ := ctx.Value(keyActivityID).(id.ID)
	return val
}
