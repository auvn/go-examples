package trip

import (
	"context"
	"errors"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
	"github.com/auvn/go-examples/example1/s-framework/storage"
	"github.com/auvn/go-examples/example1/s-framework/storage/redis"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrRiderNotFound = errors.New("rider not found")
	ErrActiveExists  = errors.New("active exists")
)

type Trip struct {
	ID        id.ID
	DriverID  *id.ID
	RiderID   id.ID
	Completed bool
}

type Driver struct {
	ID     id.ID
	TripID id.ID
}

type Rider struct {
	ID     id.ID
	TripID id.ID
}

type Trips struct {
	riders  storage.Simple
	trips   storage.Simple
	drivers storage.Simple
}

func (tt *Trips) Create(ctx context.Context, t Trip) error {
	if err := tt.trips.Set(ctx, t.ID, t); err != nil {
		return err
	}

	if err := tt.riders.Set(ctx, t.RiderID, Rider{ID: t.RiderID, TripID: t.ID}); err != nil {
		return err
	}

	return nil
}

func (tt *Trips) ByID(ctx context.Context, id id.ID) (*Trip, error) {
	var trip Trip
	if err := tt.trips.Get(ctx, id, &trip); err != nil {
		return nil, err
	}
	return &trip, nil
}

func (tt *Trips) ByRider(ctx context.Context, riderID id.ID) (*Trip, error) {
	var rider Rider
	if err := tt.riders.Get(ctx, riderID, &rider); err != nil {
		return nil, err
	}

	var trip Trip
	if err := tt.trips.Get(ctx, rider.TripID, &trip); err != nil {
		return nil, err
	}

	return &trip, nil
}
func (tt *Trips) ByDriver(ctx context.Context, driverID id.ID) (*Trip, error) {
	var driver Driver
	if err := tt.drivers.Get(ctx, driverID, &driver); err != nil {
		return nil, err
	}

	var trip Trip
	if err := tt.trips.Get(ctx, driver.TripID, &trip); err != nil {
		return nil, err
	}

	return &trip, nil
}

func (tt Trips) AssignDriver(ctx context.Context, d Driver) error {
	trip, err := tt.ByID(ctx, d.TripID)
	if err != nil {
		return err
	}

	trip.DriverID = &d.ID

	if err := tt.trips.Set(ctx, trip.ID, trip); err != nil {
		return err
	}

	return tt.drivers.Set(ctx, d.ID, d)
}

func NewTrips(r redis.Client) *Trips {
	return &Trips{
		riders:  redis.NewSimpleStorage("riders", r),
		trips:   redis.NewSimpleStorage("trips", r),
		drivers: redis.NewSimpleStorage("drivers", r),
	}
}
