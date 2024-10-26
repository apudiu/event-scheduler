package payload

import (
	"errors"
	"github.com/apudiu/event-scheduler/helper"
)

// GobPayload encodes data to gob & decodes from gob
type GobPayload struct {
	isMarshalled bool
	unmarshalled any
	marshalled   []byte
}

func (tp *GobPayload) Marshal() ([]byte, error) {
	// do not encode multiple times
	if tp.isMarshalled {
		return tp.marshalled, nil
	}

	b, e := helper.EncodeToGob(tp.unmarshalled)
	if e != nil {
		return nil, e
	}

	tp.marshalled = b
	tp.isMarshalled = true
	return b, nil
}

func (tp *GobPayload) IsMarshaled() bool {
	return tp.isMarshalled
}

func (tp *GobPayload) Unmarshal(targetPtr any) error {
	if !tp.isMarshalled {
		return errors.New("value is not marshalled")
	}

	e := helper.DecodeFromGob(tp.marshalled, targetPtr)
	if e != nil {
		return e
	}

	return nil
}

// NewGobPayload is meant to create a payload with unmarshalled data
func NewGobPayload(data any) *GobPayload {
	return &GobPayload{
		unmarshalled: data,
	}
}

// NewEncodedGobPayload is meant to create a payload from gob marshalled data
func NewEncodedGobPayload(gobEncodedData []byte) *GobPayload {
	return &GobPayload{
		marshalled:   gobEncodedData,
		isMarshalled: true,
	}
}