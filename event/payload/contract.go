package payload

// TransferablePayload guarantees a payload type that can e encoded and decoded for storage
type TransferablePayload interface {
	// EventName get associated event name
	EventName() string

	// SetEventName sets event name, it is mainly inside events dispatch method to set event name in payload
	SetEventName(string)

	// IsMarshaled confirms if the data is already marshalled or not
	IsMarshaled() bool

	// Marshal encodes the data
	Marshal() ([]byte, error)

	// Unmarshal decodes the data to the 'targetPtr'
	Unmarshal(targetPtr any) error
}
