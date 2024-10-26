package bootstrap

import (
	scheduler "github.com/apudiu/event-scheduler"
	"github.com/apudiu/event-scheduler/db/driver/gormdriver"
	"github.com/apudiu/event-scheduler/example/listeners"
)

var SCD *scheduler.Scheduler

// register listeners to events
var eventListeners = scheduler.Listeners{
	"SendEmail": {listeners.SendEmailListener},
}

func InitScheduler() {
	// use provided 'gormdriver' or make your own by implementing db.DataPersistent interface
	s := gormdriver.New(DB)

	// init scheduler with listeners mapping
	SCD = scheduler.NewScheduler(s, eventListeners)
}
