package gormdriver

import (
	"github.com/apudiu/event-scheduler/event"
	"github.com/apudiu/event-scheduler/event/payload"
	"time"
)

// Model is gorm model for events storage, it generates a table named "scheduler_events"
type Model struct {
	ID      uint       `json:"id" gorm:"primaryKey"`
	Name    string     `json:"name" gorm:"type:varchar(250);index;not null;"`
	Payload []byte     `json:"payload" gorm:"type:varbinary(500);not null;"`
	RunAt   *time.Time `json:"run_at" gorm:"type:datetime;"`
	Cron    *string    `json:"cron" gorm:"type:varchar(150);"`

	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}

func (*Model) TableName() string {
	return "scheduler_events"
}

func (m *Model) GetEvent() *event.Event {
	return &event.Event{
		ID:      m.ID,
		Name:    m.Name,
		Payload: payload.NewEncodedGobPayload(m.Payload),
		Cron:    m.Cron,
	}
}
