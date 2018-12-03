package history

import (
	"context"
	"time"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
	"github.com/auvn/go-examples/example1/frwk-storage/nosql/monga"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Record struct {
	ID       id.ID         `bson:"_id,omitempty"`
	TripID   id.ID         `bson:"trip_id"`
	RiderID  id.ID         `bson:"rider_id"`
	DriverID id.ID         `bson:"driver_id"`
	Distance int           `bson:"distance"`
	Duration time.Duration `bson:"duration"`
}
type History struct {
	m *mgo.Collection
}

func (h *History) Save(ctx context.Context, r Record) error {
	return h.m.Insert(r)
}

func (h *History) LastByRider(ctx context.Context, riderID id.ID) (Record, error) {
	var r Record
	if err := h.m.Find(bson.M{"rider_id": riderID}).One(&r); err != nil {
		return Record{}, err
	}
	return r, nil
}

func New(m *monga.Client) *History {
	return &History{
		m: m.C("history"),
	}
}
