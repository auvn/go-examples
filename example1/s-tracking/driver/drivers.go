package driver

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/auvn/go-examples/example1/s-framework/builtin/id"
)

type Driver struct {
	DriverID  id.ID
	UpdatedAt time.Time
}

var (
	ErrNotFound = errors.New("not found")
)

type Drivers struct {
	m sync.Mutex

	drivers  map[id.ID]Driver
	tempLock map[id.ID]time.Time
	locked   map[id.ID]struct{}

	lookupTTL time.Duration
}

func (dd *Drivers) Update(ctx context.Context, driver id.ID) error {
	dd.m.Lock()
	defer dd.m.Unlock()

	dd.drivers[driver] = Driver{
		DriverID:  driver,
		UpdatedAt: time.Now(),
	}
	return nil
}

func (dd *Drivers) Lookup(ctx context.Context) (id.ID, error) {
	dd.m.Lock()
	defer dd.m.Unlock()

	for driverID, driver := range dd.drivers {
		if time.Since(driver.UpdatedAt) > 5*time.Second {
			continue
		}

		if tempLockedAt, ok := dd.tempLock[driverID]; ok && time.Since(tempLockedAt) <= dd.lookupTTL {
			continue
		}

		if _, ok := dd.locked[driverID]; ok {
			continue
		}

		dd.tempLock[driverID] = time.Now()

		return driverID, nil
	}
	return "", ErrNotFound
}

func (dd *Drivers) Lock(ctx context.Context, id id.ID) error {
	dd.m.Lock()
	defer dd.m.Unlock()
	delete(dd.tempLock, id)
	dd.locked[id] = struct{}{}
	return nil
}

func (dd *Drivers) Unlock(ctx context.Context, id id.ID) error {
	dd.m.Lock()
	defer dd.m.Unlock()

	delete(dd.locked, id)
	return nil
}

func NewDrivers() *Drivers {
	return &Drivers{
		drivers:   map[id.ID]Driver{},
		tempLock:  map[id.ID]time.Time{},
		locked:    map[id.ID]struct{}{},
		lookupTTL: 5 * time.Second,
	}
}
