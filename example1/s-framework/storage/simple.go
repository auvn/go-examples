package storage

import (
	"context"
	"errors"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
)

var (
	ErrNotFound = errors.New("not found")
)

type Simple interface {
	Set(ctx context.Context, id id.ID, v interface{}) error
	Get(ctx context.Context, id id.ID, dest interface{}) error
	Delete(ctx context.Context, id id.ID) error
}
