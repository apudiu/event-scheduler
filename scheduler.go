// Package scheduler allows delayed and recurring event dispatching
// with ways to persist events, so events doesn't gets lost when server restarted
package scheduler

import (
	"context"
	"fmt"
	"github.com/apudiu/event-scheduler/dp"
	"log"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/apudiu/event-scheduler/event"
	"github.com/apudiu/event-scheduler/event/payload"
)

// Scheduler data structure
type Scheduler struct {
	p           dp.DataPersistent
	listeners   Listeners
	cron        *cron.Cron
	cronEntries map[string]cron.EntryID
}

// Listeners has attached event listeners
type Listeners map[string][]ListenFunc

// ListenFunc function that listens to events
type ListenFunc func(data payload.TransferablePayload)

// NewScheduler creates a new scheduler
func NewScheduler(p dp.DataPersistent, listeners Listeners) *Scheduler {
	return &Scheduler{
		p:           p,
		listeners:   listeners,
		cron:        cron.New(),
		cronEntries: map[string]cron.EntryID{},
	}
}

// AddListener adds the listener function to Listeners
func (s Scheduler) AddListener(evtName string, listenFunctions []ListenFunc) {
	s.listeners[evtName] = listenFunctions
}

// callListeners calls the event listener of provided event
func (s Scheduler) callListeners(evt event.Event) {
	listenerFns, ok := s.listeners[evt.Name]
	if ok {
		for _, fn := range listenerFns {
			go fn(evt.Payload)
		}

		_, err := s.p.Delete(evt.ID)
		if err != nil {
			log.Print("💀 error: ", err)
		}
	} else {
		log.Print("💀 error: couldn't find event listeners attached to ", evt.Name)
	}

}

// CheckEventsInInterval checks the event in given interval
func (s Scheduler) CheckEventsInInterval(ctx context.Context, duration time.Duration) {
	ticker := time.NewTicker(duration)
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				log.Println("⏰ Ticks Received...")
				events := s.checkDueEvents()
				//fmt.Printf("------------------- %#v \n", events)
				for _, e := range events {
					s.callListeners(e)
				}
			}

		}
	}()
}

// checkDueEvents checks and returns due events
func (s Scheduler) checkDueEvents() []event.Event {
	sEvents, err := s.p.GetAll()
	if err != nil {
		fmt.Printf("--------------------> %#v \n", err)
		log.Print("💀 error: ", err)
		return nil
	}

	events := make([]event.Event, 0)
	for _, se := range sEvents {
		events = append(events, *se.GetEvent())
	}

	return events
}

// Schedule schedules the provided events
func (s Scheduler) Schedule(event string, payload payload.TransferablePayload, runAt time.Time) error {
	log.Print("🚀 Scheduling event ", event, " to run at ", runAt)

	_, err := s.p.AddWithTime(event, payload, runAt)
	return err
}

func (s Scheduler) ScheduleDur(event string, payload payload.TransferablePayload, runAfter time.Duration) error {
	log.Print("🚀 Scheduling event with dur ", event, " to run after ", runAfter)

	_, err := s.p.AddWithDuration(event, payload, runAfter)
	return err
}

// ScheduleRecurring schedules a cron job
func (s Scheduler) ScheduleRecurring(evtName string, payload payload.TransferablePayload, cronStr string) error {
	log.Print("🚀 Scheduling event ", evtName, " with cron string ", cronStr)

	entryID, ok := s.cronEntries[evtName]
	if ok {
		s.cron.Remove(entryID)
		err := s.p.UpdateByName(evtName, &event.Event{
			Payload: payload,
			Cron:    &cronStr,
		})
		if err != nil {
			log.Print("schedule cron update error: ", err)
			return err
		}
	} else {
		_, err := s.p.AddRecurring(evtName, payload, cronStr)
		if err != nil {
			log.Print("schedule cron insert error: ", err)
			return err
		}
	}

	listenerFns, ok := s.listeners[evtName]
	if ok {
		cronId, err := s.cron.AddFunc(cronStr, func() {
			for _, fn := range listenerFns {
				fn(payload)
			}
		})
		s.cronEntries[evtName] = cronId
		if err != nil {
			log.Print("💀 error: ", err)
			return err
		}
	}

	return nil
}

// attachCronJobs attaches cron jobs
func (s Scheduler) attachCronJobs() {
	log.Printf("Attaching cron jobs")
	sEvents, err := s.p.GetAllRecurring()
	if err != nil {
		log.Print("💀 error: ", err)
	}

	for _, se := range sEvents {
		evt := *se.GetEvent()
		listenerFns, ok := s.listeners[evt.Name]
		if ok {
			entryID, err2 := s.cron.AddFunc(*evt.Cron, func() {
				for _, fn := range listenerFns {
					fn(evt.Payload)
				}
			})
			if err2 != nil {
				log.Print("💀 error: ", err2)
				continue
			}

			s.cronEntries[evt.Name] = entryID
		}
	}
}

// StartCron starts cron job. It returns cleanup to stop the cron
func (s Scheduler) StartCron() (cleanup func()) {
	s.attachCronJobs()
	s.cron.Start()

	return func() {
		s.cron.Stop()
	}
}
