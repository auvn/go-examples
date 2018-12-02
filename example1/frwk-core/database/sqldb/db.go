package sqldb

import (
	"context"
	"log"

	"github.com/gojuno/sqlx"
)

type namedPreparer interface {
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
}

func namedGet(ctx context.Context, np namedPreparer, dest interface{}, query string, arg interface{}) error {
	stmt, err := np.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.GetContext(ctx, dest, arg)
}

type Tx struct {
	*sqlx.Tx
}

func (tx *Tx) NamedGet(ctx context.Context, dest interface{}, query string, arg interface{}) error {
	return namedGet(ctx, tx, dest, query, arg)
}

type DB struct {
	*sqlx.DB
}

func (db *DB) Transaction(ctx context.Context, fn func(ctx context.Context, tx *Tx) error) (err error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	if err := fn(ctx, &Tx{tx}); err != nil {
		return err
	}
	return nil
}

func New(driver string, connstring string) *DB {
	db, err := sqlx.Connect(driver, connstring)
	if err != nil {
		log.Fatalf("db: failed to connect(%s, %s): %v", driver, connstring, err)
	}

	return &DB{DB: db}
}
