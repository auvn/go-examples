package encoding

import (
	"encoding/json"
	"io"
)

func MarshalToWriter(v interface{}, w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(v)
}

func UnmarshalReader(r io.Reader, dest interface{}) error {
	dec := json.NewDecoder(r)
	return dec.Decode(dest)
}
