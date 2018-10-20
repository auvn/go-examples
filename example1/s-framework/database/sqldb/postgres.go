package sqldb

import (
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
)

type Postgres struct{}

func NewPostgres(name string) *DB {
	return New("pgx",
		fmt.Sprintf("postgres://user:userpwd@postgres:5432/%s?connect_timeout=5&sslmode=disable", name))
}
