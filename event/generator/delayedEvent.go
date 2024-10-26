package generator

import (
	"fmt"
	scheduler "github.com/apudiu/event-scheduler"
	"github.com/apudiu/event-scheduler/event/payload"
	"log"
	"time"
)

type DelayedEvent struct {
	name string
	s    *scheduler.Scheduler
}

// Name returns name of the event, it exits when name not set, so be sure to call it after the event is created with NewDelayedEvent
func (t *DelayedEvent) Name() string {
	if t.name == "" {
		log.Fatalln("Event is not initialized with new, first make a new event then try to access the name")
	}
	return t.name
}

// Dispatch dispatches event with specified time.Time
func (t *DelayedEvent) Dispatch(data payload.TransferablePayload, at time.Time) error {
	if t.s == nil {
		return fmt.Errorf("scheduler is not initialized: %v\n", t.s)
	}

	return t.s.Schedule(t.name, data, at)
}

// DispatchDur dispatches event with specified time.Duration
func (t *DelayedEvent) DispatchDur(data payload.TransferablePayload, after time.Duration) error {
	if t.s == nil {
		return fmt.Errorf("scheduler is not initialized: %v\n", t.s)
	}

	at := time.Now().Add(after)
	return t.Dispatch(data, at)
}

// NewDelayedEvent creates a new delayed event that can be dispatched
func NewDelayedEvent(name string, s *scheduler.Scheduler) *DelayedEvent {
	return &DelayedEvent{
		name: name,
		s:    s,
	}
}
