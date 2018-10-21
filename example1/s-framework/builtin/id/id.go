package id

import (
	"encoding/json"

	"github.com/pborman/uuid"
	"github.com/pkg/errors"
)

type ID string

func (id *ID) UnmarshalJSON(bb []byte) error {
	if string(bb) == `""` {
		return errors.Errorf("id is empty")
	}
	var uuidVal uuid.UUID
	if err := json.Unmarshal(bb, &uuidVal); err != nil {
		return err
	}
	*id = ID(uuidVal.String())
	return nil
}

func (id ID) MarshalJSON() ([]byte, error) {
	uuidValue := uuid.Parse(string(id))
	return json.Marshal(uuidValue)
}

func New() ID {
	return ID(uuid.New())
}
