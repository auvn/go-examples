package storage

import (
	"context"
	"errors"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
)

var (
	ErrNotFound = errors.New("not found")
)

type Primitive interface {
	Set(ctx context.Context, id id.ID, v interface{}) error
	Get(ctx context.Context, id id.ID, dest interface{}) error
	Delete(ctx context.Context, id id.ID) error
}
