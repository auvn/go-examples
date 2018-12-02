package redis

import (
	"context"
	"encoding/json"

	"github.com/auvn/go-examples/example1/frwk-core/builtin/id"
	"github.com/auvn/go-examples/example1/frwk-core/storage"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type Client interface {
	redis.Cmdable
}

type Config struct {
	Addrs []string
}

func New(cfg Config) Client {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		ClusterSlots: func() ([]redis.ClusterSlot, error) {
			nodes := make([]redis.ClusterNode, len(cfg.Addrs))
			for i := range cfg.Addrs {
				nodes[i].Addr = cfg.Addrs[i]
			}
			return []redis.ClusterSlot{
				{
					Start: 0,
					End:   16386,
					Nodes: nodes,
				},
			}, nil
		},
	})
	return client
}

type Value struct {
	V interface{}
}

func (v Value) MarshalBinary() ([]byte, error)  { return json.Marshal(v.V) }
func (v Value) UnmarshalBinary(bb []byte) error { return json.Unmarshal(bb, v.V) }

type PrimitiveStorage struct {
	name  string
	redis redis.Cmdable
}

func (s *PrimitiveStorage) Set(ctx context.Context, id id.ID, v interface{}) error {
	return s.redis.HSet(s.name, string(id), Value{v}).Err()
}

func (s *PrimitiveStorage) Get(ctx context.Context, id id.ID, dest interface{}) error {
	ok, err := s.redis.HExists(s.name, string(id)).Result()
	if err != nil {
		return err
	}

	if !ok {
		return errors.Wrapf(storage.ErrNotFound, "%s: %q", s.name, id)
	}

	return s.redis.HGet(s.name, string(id)).Scan(Value{dest})
}

func (s *PrimitiveStorage) Delete(ctx context.Context, id id.ID) error {
	return s.redis.HDel(s.name, string(id)).Err()
}

func NewPrimitiveStorage(name string, redis Client) *PrimitiveStorage {
	return &PrimitiveStorage{
		name:  name,
		redis: redis,
	}
}
