package trip

import (
	"context"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
	"github.com/auvn/go-examples/example1/s-framework/database/sqldb"
)

type Trip struct {
	ID     id.ID
	Driver id.ID
	Rider  id.ID
	Status string
}

type Trips struct {
	*sqldb.DB
}

func (rr Trips) Create(ctx context.Context, r Trip) error {
	const query = `
		insert into trips (id, driver, rider, status) 
		values (:id, :driver, :rider, :status)`

	if _, err := rr.NamedExecContext(ctx, query, r); err != nil {
		return err
	}
	return nil
}

func (rr Trips) Update(ctx context.Context, r Trip) error {
	const (
		selectQuery = `select status from trips id = :id for update`
		updateQuery = `update trips set status=:status where id = :id`
	)

	return rr.Transaction(ctx, func(ctx context.Context, tx *sqldb.Tx) error {
		var (
			status string
		)
		if err := tx.NamedGet(ctx, &status, selectQuery, r); err != nil {
			return err
		}

		if _, err := tx.NamedExecContext(ctx, updateQuery, r); err != nil {
			return err
		}
		return nil
	})
}

func NewTrips(db *sqldb.DB) *Trips {
	return &Trips{db}
}
