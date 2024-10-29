package bootstrap

import (
	scheduler "github.com/apudiu/event-scheduler"
	"github.com/apudiu/event-scheduler/dp/driver/gormdriver"
	"github.com/apudiu/event-scheduler/example/eventListenersMap"
)

var SCD *scheduler.Scheduler

func InitScheduler() {
	// use provided 'gormdriver' or make your own by implementing db.DataPersistent interface
	s := gormdriver.New(DB)

	// init scheduler with listeners mapping
	SCD = scheduler.NewScheduler(s, eventListenersMap.ListenersList, true)
}
