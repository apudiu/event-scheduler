package gormdriver

import (
	"fmt"
	"github.com/apudiu/event-scheduler/event"
	"github.com/apudiu/event-scheduler/event/payload"
	"gorm.io/gorm"
	"time"
)

// Driver implements dp.DataPersistent and uses gorm supported DB's for persistence
type Driver struct {
	db *gorm.DB
}

func (d *Driver) GetOne(eventId uint) (event.SchedulerEvent, error) {
	m := new(Model)
	if err := d.db.First(m, eventId).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (d *Driver) GetAll() ([]event.SchedulerEvent, error) {
	models := make([]Model, 0)
	if err := d.db.Find(&models, "cron IS NULL").Error; err != nil {
		return nil, err
	}

	events := make([]event.SchedulerEvent, 0)
	for i := range models {
		events = append(events, &models[i])
	}

	return events, nil
}

func (d *Driver) AddWithTime(name string, data payload.TransferablePayload, runAt time.Time) (event.SchedulerEvent, error) {
	p, e := data.Marshal()
	if e != nil {
		return nil, e
	}

	m := Model{
		Name:    name,
		Payload: p,
		RunAt:   &runAt,
	}

	if err := d.db.Create(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (d *Driver) AddWithDuration(name string, data payload.TransferablePayload, runAfter time.Duration) (event.SchedulerEvent, error) {
	t := time.Now().Add(runAfter)
	return d.AddWithTime(name, data, t)
}

func (d *Driver) GetAllRecurring() ([]event.SchedulerEvent, error) {
	models := make([]Model, 0)
	if err := d.db.Find(&models, "cron IS NOT NULL").Error; err != nil {
		return nil, err
	}

	var events []event.SchedulerEvent
	for i := range models {
		events = append(events, &models[i])
	}

	return events, nil
}

func (d *Driver) AddRecurring(name string, data payload.TransferablePayload, cron string) (event.SchedulerEvent, error) {
	p, e := data.Marshal()
	if e != nil {
		return nil, e
	}

	m := Model{
		Name:    name,
		Payload: p,
		Cron:    &cron,
	}

	if err := d.db.Create(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (d *Driver) AddRecurringWithDuration(name string, data payload.TransferablePayload, runEvery time.Duration) (event.SchedulerEvent, error) {
	now := time.Now()
	dur := now.Add(runEvery).Sub(now)

	cronTxt := fmt.Sprintf("@every %s", dur)

	return d.AddRecurring(name, data, cronTxt)
}

func (d *Driver) UpdateByName(eventName string, evt *event.Event) error {
	p, err := evt.Payload.Marshal()
	if err != nil {
		return err
	}

	m := &Model{
		Payload: p,
		Cron:    evt.Cron,
	}
	if err2 := d.db.Where("name = ?", eventName).Updates(m).Error; err2 != nil {
		return err2
	}
	return nil
}

func (d *Driver) Delete(eventId uint) (event.SchedulerEvent, error) {
	m := new(Model)
	if err := d.db.Delete(m, eventId).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (d *Driver) DeleteAll() ([]event.SchedulerEvent, error) {
	models := make([]Model, 0)
	if err := d.db.Delete(&models, "1 = 1").Error; err != nil {
		return nil, err
	}

	events := make([]event.SchedulerEvent, 0)
	for i := range models {
		events = append(events, &models[i])
	}
	return events, nil
}

// New creates a new gorm driver with passed gorm.DB connection for event operation/ persistence
func New(db *gorm.DB) *Driver {
	return &Driver{
		db: db,
	}
}
