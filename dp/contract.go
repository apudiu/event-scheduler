package dp

import (
	"github.com/apudiu/event-scheduler/event"
	"github.com/apudiu/event-scheduler/event/payload"
	"time"
)

// DataPersistent interface defines ways to interact with persistent layer
type DataPersistent interface {

	// GetOne retrieves any type of event by id
	GetOne(eventId uint) (event.SchedulerEvent, error)

	// GetAll retrieves pending onetime events that has no cron associated with it
	GetAll() ([]event.SchedulerEvent, error)

	// AddWithTime creates one time event with specified time.Time
	AddWithTime(name string, payload payload.TransferablePayload, runAt time.Time) (event.SchedulerEvent, error)
	// AddWithDuration creates one time event with specified time.Duration
	AddWithDuration(name string, payload payload.TransferablePayload, runAfter time.Duration) (event.SchedulerEvent, error)

	// GetAllRecurring retrieves recurring events that has cron associated with it
	GetAllRecurring() ([]event.SchedulerEvent, error)

	// AddRecurring creates recurring event with specified cron. cron should be a valid cron string (you can make one here: https://crontab.guru).
	// Also "@every [unit]" is supported. Ex: @every 5s = every 5 seconds, @every 5d = every 5 days, @every 5w = every 5 weeks ...
	AddRecurring(name string, payload payload.TransferablePayload, cron string) (event.SchedulerEvent, error)
	// AddRecurringWithDuration crates recurring event with specified duration
	AddRecurringWithDuration(name string, payload payload.TransferablePayload, runEvery time.Duration) (event.SchedulerEvent, error)

	// UpdateByName updates one event by event name, as gorm's policy "zero" value field will be excluded from updating
	UpdateByName(eventName string, evt *event.Event) error

	// Delete deletes one event by id
	Delete(eventId uint) (event.SchedulerEvent, error)

	// DeleteAll deletes all events from persistent layer
	DeleteAll() ([]event.SchedulerEvent, error)
}
