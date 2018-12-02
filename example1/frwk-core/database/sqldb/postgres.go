package sqldb

import _ "github.com/jackc/pgx/stdlib"

type Postgres struct{}

func NewPostgres(name string) *DB {
	return New("pgx", "postgres://user:userpwd@postgres:5432/user?connect_timeout=5&sslmode=disable")
}
