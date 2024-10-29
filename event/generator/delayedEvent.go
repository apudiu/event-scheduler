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
func (de *DelayedEvent) Name() string {
	if de.name == "" {
		log.Fatalln("Event is not initialized with new, first make a new event then try to access the name")
	}
	return de.name
}

// Dispatch dispatches event with specified time.Time
func (de *DelayedEvent) Dispatch(data payload.TransferablePayload, at time.Time) error {
	if de.s == nil {
		return fmt.Errorf("scheduler is not initialized: %v\n", de.s)
	}

	// set event name
	data.SetEventName(de.name)

	e := de.s.Schedule(de.name, data, at)
	return e
}

// DispatchDur dispatches event with specified time.Duration
func (de *DelayedEvent) DispatchDur(data payload.TransferablePayload, after time.Duration) error {
	if de.s == nil {
		return fmt.Errorf("scheduler is not initialized: %v\n", de.s)
	}

	data.SetEventName(de.name)

	at := time.Now().Add(after)
	return de.Dispatch(data, at)
}

// NewDelayedEvent creates a new delayed event that can be dispatched
func NewDelayedEvent(name string, s *scheduler.Scheduler) *DelayedEvent {
	return &DelayedEvent{
		name: name,
		s:    s,
	}
}
