// Package scheduler allows delayed and recurring event dispatching
// with ways to persist events, so events doesn't gets lost when server restarted
package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/apudiu/event-scheduler/db"
	"github.com/apudiu/event-scheduler/event"
)

// Scheduler data structure
type Scheduler struct {
	dp          db.DataPersistent
	listeners   Listeners
	cron        *cron.Cron
	cronEntries map[string]cron.EntryID
}

// Listeners has attached event listeners
type Listeners map[string][]ListenFunc

// ListenFunc function that listens to events
type ListenFunc func(string)

// NewScheduler creates a new scheduler
func NewScheduler(dp db.DataPersistent, listeners Listeners) *Scheduler {
	return &Scheduler{
		dp:          dp,
		listeners:   listeners,
		cron:        cron.New(),
		cronEntries: map[string]cron.EntryID{},
	}
}

// AddListener adds the listener function to Listeners
func (s Scheduler) AddListener(event string, listenFunctions []ListenFunc) {
	s.listeners[event] = listenFunctions
}

// callListeners calls the event listener of provided event
func (s Scheduler) callListeners(event event.Event) {
	listenerFns, ok := s.listeners[event.Name]
	if ok {
		for _, fn := range listenerFns {
			go fn(event.Payload)
		}

		_, err := s.dp.Delete(event.ID)
		if err != nil {
			log.Print("💀 error: ", err)
		}
	} else {
		log.Print("💀 error: couldn't find event listeners attached to ", event.Name)
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
				for _, e := range *events {
					s.callListeners(e)
				}
			}

		}
	}()
}

// checkDueEvents checks and returns due events
func (s Scheduler) checkDueEvents() *[]event.Event {
	sEvents, err := s.dp.GetAll()
	if err != nil {
		log.Print("💀 error: ", err)
		return nil
	}

	events := make([]event.Event, 0)
	for _, se := range sEvents {
		events = append(events, *se.GetEvent())
	}

	return &events
}

// Schedule schedules the provided events
func (s Scheduler) Schedule(event string, payload string, runAt time.Time) {
	log.Print("🚀 Scheduling event ", event, " to run at ", runAt)
	_, err := s.dp.AddWithTime(event, payload, runAt)
	if err != nil {
		log.Print("schedule insert error: ", err)
	}
}

func (s Scheduler) ScheduleDur(event string, payload string, runAfter time.Duration) {
	log.Print("🚀 Scheduling event with dur ", event, " to run after ", runAfter)
	_, err := s.dp.AddWithDuration(event, payload, runAfter)
	if err != nil {
		log.Print("schedule insert error: ", err)
	}
}

// ScheduleRecurring schedules a cron job
func (s Scheduler) ScheduleRecurring(evtName string, payload string, cronStr string) {
	log.Print("🚀 Scheduling event ", evtName, " with cron string ", cronStr)
	entryID, ok := s.cronEntries[evtName]
	if ok {
		s.cron.Remove(entryID)
		err := s.dp.UpdateByName(evtName, &event.Event{
			Payload: payload,
			Cron:    &cronStr,
		})
		if err != nil {
			log.Print("schedule cron update error: ", err)
		}
	} else {
		_, err := s.dp.AddRecurring(evtName, payload, cronStr)
		if err != nil {
			log.Print("schedule cron insert error: ", err)
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
		}
	}
}

// attachCronJobs attaches cron jobs
func (s Scheduler) attachCronJobs() {
	log.Printf("Attaching cron jobs")
	sEvents, err := s.dp.GetAllRecurring()
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
