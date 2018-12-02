package encoding

import (
	"encoding/json"
	"io"
)

type RawMessage json.RawMessage

func (m *RawMessage) UnmarshalJSON(bb []byte) error {
	var msg json.RawMessage
	if err := msg.UnmarshalJSON(bb); err != nil {
		return err
	}

	*m = RawMessage(msg)
	return nil
}

func (m RawMessage) MarshalJSON() ([]byte, error) {
	return json.RawMessage(m).MarshalJSON()
}

func (m RawMessage) UnmarshalTo(dest interface{}) error {
	return json.Unmarshal([]byte(m), dest)
}

func MarshalToWriter(v interface{}, w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(v)
}

func UnmarshalReader(r io.Reader, dest interface{}) error {
	dec := json.NewDecoder(r)
	return dec.Decode(dest)
}
