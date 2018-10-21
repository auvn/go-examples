package trip

import (
	"context"
	"errors"
	"fmt"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
	"github.com/auvn/go-examples/example1/s-framework/database/sqldb"
)

var (
	ErrActiveExists = errors.New("active exists")
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

func (tt Trips) ActiveByRider(ctx context.Context, rider id.ID) (*Trip, error) {
	const query = `
		select id, driver, rider 
		from trips where active and rider = :rider`

	return nil, nil
}

func (rr Trips) Create(ctx context.Context, r Trip) error {
	const (
		query = `
			insert into trips (id, rider, status) 
			values (:id, :rider, :status)`
	)
	fmt.Printf("%+v\n", r)
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
