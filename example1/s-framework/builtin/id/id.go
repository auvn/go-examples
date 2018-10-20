package id

import (
	"encoding/json"

	"github.com/pborman/uuid"
	"github.com/pkg/errors"
)

type ID string

func (id *ID) UnmarshalJSON(bb []byte) error {
	if len(bb) == 0 {
		return errors.Errorf("empty id")
	}
	var uuidVal uuid.UUID
	if err := json.Unmarshal(bb, &uuidVal); err != nil {
		return err
	}
	*id = ID(string(uuidVal))
	return nil
}

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id)
}

func New() ID {
	return ID(uuid.New())
}
