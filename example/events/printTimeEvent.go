package events

import (
	"github.com/apudiu/event-scheduler/event/generator"
	"github.com/apudiu/event-scheduler/example/bootstrap"
	"github.com/apudiu/event-scheduler/example/eventListenersMap"
)

var PrintTimeEvent = func() *generator.RecurringEvent {
	return generator.NewRecurringEvent(eventListenersMap.EventPrintTime, bootstrap.SCD)
}
