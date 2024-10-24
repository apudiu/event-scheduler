package gormdriver

import (
	"github.com/apudiu/event-scheduler/event"
	"time"
)

type Model struct {
	ID      uint       `json:"id" gorm:"primaryKey"`
	Name    string     `json:"name" gorm:"type:varchar(250);uniqueIndex;not null;"`
	Payload string     `json:"payload" gorm:"type:text;not null;"`
	RunAt   *time.Time `json:"run_at" gorm:"type:timestamp;"`
	Cron    *string    `json:"cron" gorm:"type:varchar(150);"`

	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}

func (m *Model) GetEvent() *event.Event {
	return &event.Event{
		ID:      m.ID,
		Name:    m.Name,
		Payload: m.Payload,
		Cron:    m.Cron,
	}
}
