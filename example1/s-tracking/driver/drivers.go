package driver

import (
	"context"
	"time"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
	"github.com/auvn/go-examples/example1/frwk-storage/nosql/monga"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Driver struct {
	ID   id.ID
	Busy bool
}

type driverRecord struct {
	ID           id.ID     `bson:"_id,omitempty"`
	Busy         bool      `bson:"busy"`
	UpdatedAt    time.Time `bson:"updated_at"`
	LockDeadline time.Time `bson:"lock_deadline"`
}

var (
	ErrNotFound = errors.New("not found")
)

type Drivers struct {
	collection *mgo.Collection
	lockTTL    time.Duration
}

func (dd *Drivers) Update(ctx context.Context, d Driver) error {
	rec := driverRecord{
		ID:        d.ID,
		Busy:      d.Busy,
		UpdatedAt: time.Now().UTC(),
	}
	_, err := dd.collection.Upsert(bson.M{"_id": d.ID}, rec)
	return err
}

func (dd *Drivers) Lookup(ctx context.Context) (id.ID, error) {
	now := time.Now().UTC()

	var d driverRecord

	ch := mgo.Change{
		Update: bson.M{"$set": bson.M{"lock_deadline": now.Add(dd.lockTTL)}},
	}

	_, err := dd.collection.
		Find(bson.M{
			"busy":          false,
			"lock_deadline": bson.M{"$lt": now}}).
		Sort("updated_at").
		Apply(ch, &d)

	if err != nil {
		switch err {
		case mgo.ErrNotFound:
			return "", ErrNotFound
		default:
			return "", err
		}
	}

	return d.ID, nil
}

func NewDrivers(s *monga.Client) *Drivers {
	return &Drivers{
		collection: s.C("drivers"),
		lockTTL:    15 * time.Second,
	}
}
