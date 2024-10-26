package payload

// TransferablePayload guarantees a payload type that can e encoded and decoded for storage
type TransferablePayload interface {
	IsMarshaled() bool
	Marshal() error
	Unmarshal(targetPtr any) error
}
