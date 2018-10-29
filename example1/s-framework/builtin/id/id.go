package id

import (
	"encoding/json"

	"github.com/pborman/uuid"
	"github.com/pkg/errors"
)

var (
	errEmptyID = errors.New("empty id")
)

type ID string

func (id *ID) UnmarshalJSON(bb []byte) error {

	if string(bb) == `""` {
		return errEmptyID
	}
	var uuidVal uuid.UUID
	if err := json.Unmarshal(bb, &uuidVal); err != nil {
		return err
	}
	*id = ID(uuidVal.String())
	return nil
}

func (id ID) MarshalJSON() ([]byte, error) {
	if id == "" {
		return nil, errEmptyID
	}
	uuidValue := uuid.Parse(string(id))
	return json.Marshal(uuidValue)
}

func New() ID {
	return ID(uuid.New())
}
