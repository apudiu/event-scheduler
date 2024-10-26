package events

import (
	"github.com/apudiu/event-scheduler/event/generator"
	"github.com/apudiu/event-scheduler/example/bootstrap"
	"github.com/apudiu/event-scheduler/example/eventListenersMap"
)

// create event like this by directly calling scheduler
//func SendEmailEvent(email string, content string) {
//
//	// prepare payload with provided payload.GobPayload
//	// or create one yourself by implementing payload.TransferablePayload interface
//	p := payload.NewGobPayload(
//		fmt.Sprintf("Send email to %s with content %s", email, content),
//	)
//
//	// send the event using the payload
//	bootstrap.SCD.Schedule("SendEmail", p, time.Now().Add(time.Second*5))
//}

// create event using event generator

var SendEmailEvent = func() *generator.DelayedEvent {
	return generator.NewDelayedEvent(eventListenersMap.EventSendEmail, bootstrap.SCD)
}
