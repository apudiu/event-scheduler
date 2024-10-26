package generator

import (
	"fmt"
	scheduler "github.com/apudiu/event-scheduler"
	"github.com/apudiu/event-scheduler/event/payload"
	"log"
	"time"
)

type RecurringEvent struct {
	name string
	s    *scheduler.Scheduler
}

// Name returns name of the event, it exits when name not set, so be sure to call it after the event is created with NewDelayedEvent
func (t *RecurringEvent) Name() string {
	if t.name == "" {
		log.Fatalln("Event is not initialized with new, first make a new event then try to access the name")
	}
	return t.name
}

// Dispatch dispatches event with specified cronStr, cronStr need to be a valid cronStr
func (t *RecurringEvent) Dispatch(data payload.TransferablePayload, cronStr string) error {
	if t.s == nil {
		return fmt.Errorf("scheduler is not initialized: %v\n", t.s)
	}

	return t.s.ScheduleRecurring(t.name, data, cronStr)
}

// DispatchDur dispatches event with specified 'every' time.Duration
func (t *RecurringEvent) DispatchDur(data payload.TransferablePayload, every time.Duration) error {
	if t.s == nil {
		return fmt.Errorf("scheduler is not initialized: %v\n", t.s)
	}

	cronStr := "@every " + every.String()
	return t.s.ScheduleRecurring(t.name, data, cronStr)
}

// NewRecurringEvent creates a new event that can be dispatched for recurring execution
func NewRecurringEvent(name string, s *scheduler.Scheduler) *RecurringEvent {
	return &RecurringEvent{
		s:    s,
		name: name,
	}
}
