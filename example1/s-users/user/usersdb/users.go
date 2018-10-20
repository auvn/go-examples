package usersdb

import (
	"context"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
	"github.com/auvn/go-examples/example1/s-users/user"
)

type DB struct {
}

func (*DB) Create(ctx context.Context, u user.User) error {
	panic("implement me")
}

func (*DB) ByID(ctx context.Context, id id.ID) (user.User, error) {
	panic("implement me")
}

func New() *DB {
	return &DB{}
}
