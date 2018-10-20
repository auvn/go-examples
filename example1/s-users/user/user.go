package user

import (
	"context"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
)

type User struct {
	ID id.ID
}

type Users interface {
	Create(ctx context.Context, u User) error
	ByID(ctx context.Context, id id.ID) (User, error)
}
