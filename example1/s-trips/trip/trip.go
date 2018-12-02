package trip

import (
	"context"
	"errors"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
	"github.com/auvn/go-examples/example1/frwk-storage/nosql/monga"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrNotFound = errors.New("not found")
)

type Trip struct {
	ID        id.ID  `bson:"_id,omitempty"`
	DriverID  *id.ID `bson:"driver_id,omitempty"`
	RiderID   id.ID  `bson:"rider_id"`
	Completed bool   `bson:"completed"`
}

type Driver struct {
	ID     id.ID
	TripID id.ID
}

type Trips struct {
	trips *mgo.Collection
}

func (tt *Trips) Create(ctx context.Context, t Trip) error {
	return tt.trips.Insert(t)
}

func (tt *Trips) ByID(ctx context.Context, id id.ID) (*Trip, error) {
	var t Trip
	if err := tt.trips.FindId(id).One(&t); err != nil {
		return nil, err
	}
	return &t, nil
}

func (tt *Trips) ByRider(ctx context.Context, riderID id.ID) (*Trip, error) {
	var t Trip

	if err := tt.trips.Find(bson.M{"rider_id": riderID}).One(&t); err != nil {
		return nil, err
	}

	return &t, nil
}
func (tt *Trips) ByDriver(ctx context.Context, driverID id.ID) (*Trip, error) {
	var t Trip

	if err := tt.trips.Find(bson.M{"driver_id": driverID}).One(&t); err != nil {
		return nil, err
	}

	return &t, nil
}

func (tt Trips) AssignDriver(ctx context.Context, d Driver) error {
	ch := mgo.Change{
		Update: bson.M{"$set": bson.M{"driver_id": d.ID}},
	}
	_, err := tt.trips.Find(bson.M{"_id": d.TripID, "driver_id": nil}).Apply(ch, nil)
	return err
}

func (tt Trips) Complete(ctx context.Context, driverID id.ID) (*Trip, error) {
	ch := mgo.Change{
		Update: bson.M{"$set": bson.M{"completed": true}},
	}

	var t Trip
	_, err := tt.trips.Find(bson.M{"driver_id": driverID, "completed": false}).Apply(ch, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func NewTrips(s *monga.Client) *Trips {
	return &Trips{
		trips: s.C("trips"),
	}
}
