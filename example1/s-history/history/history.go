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
	ID        id.ID         `bson:"_id,omitempty"`
	TripID    id.ID         `bson:"trip_id"`
	RiderID   id.ID         `bson:"rider_id"`
	DriverID  id.ID         `bson:"driver_id"`
	Distance  int           `bson:"distance"`
	Duration  time.Duration `bson:"duration"`
	Breakdown *breakdown    `bson:"breakdown,omitempty"`
	CreatedAt time.Time     `bson:"created_at,omitempty"`
}

type breakdown struct {
	Total string `bson:"total"`
}

type Breakdown struct {
	TripID id.ID
	Total  string
}

type History struct {
	m *mgo.Collection
}

func (h *History) Save(ctx context.Context, r Record) error {
	return h.m.Insert(r)
}

func (h *History) LastByRider(ctx context.Context, riderID id.ID) (Record, error) {
	var r Record
	if err := h.m.Find(bson.M{"rider_id": riderID}).Sort("-created_at").One(&r); err != nil {
		return Record{}, err
	}
	return r, nil
}

func (h *History) SaveBreakdown(ctx context.Context, b Breakdown) error {
	err := h.m.UpdateId(
		b.TripID,
		bson.M{"$set": bson.M{"breakdown": bson.M{"total": b.Total}}})
	return err
}

func New(m *monga.Client) *History {
	return &History{
		m: m.C("history"),
	}
}
