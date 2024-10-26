package helper

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
)

// EncodeToGob encoded p to gob, it is meant to be used with message payload encoding
func EncodeToGob(p any) (data []byte, err error) {
	buff := new(bytes.Buffer)
	if gob.NewEncoder(buff).Encode(p) != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

// DecodeFromGob decodes p (gob encoded) to t, it is meant to be used with message payload decoding.
// pass t as ptr so it is possible to decode to t
func DecodeFromGob[T any](p []byte, t T) error {
	if reflect.ValueOf(t).Kind() != reflect.Ptr {
		return fmt.Errorf("t must be a pointer of any type, provided %T", t)
	}

	buff := bytes.NewBuffer(p)
	return gob.NewDecoder(buff).Decode(t)
}
