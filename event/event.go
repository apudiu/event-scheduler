package event

import "github.com/apudiu/event-scheduler/event/payload"

type SchedulerEvent interface {
	GetEvent() *Event
}

// Event structure
type Event struct {
	ID      uint                        `json:"id"`
	Name    string                      `json:"name"`
	Payload payload.TransferablePayload `json:"payload"`
	Cron    *string                     `json:"cron"`
}
