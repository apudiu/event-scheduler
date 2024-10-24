package event

type SchedulerEvent interface {
	GetEvent() *Event
}

// Event structure
type Event struct {
	ID      uint    `json:"id"`
	Name    string  `json:"name"`
	Payload string  `json:"payload"`
	Cron    *string `json:"cron"`
}
