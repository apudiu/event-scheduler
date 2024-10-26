package eventListenersMap

import (
	scheduler "github.com/apudiu/event-scheduler"
	"github.com/apudiu/event-scheduler/example/listeners"
)

const (
	EventSendEmail = "SendEmail"
	EventPrintTime = "PrintTime"
)

// ListenersList register listeners to events
var ListenersList = scheduler.Listeners{
	EventSendEmail: {
		listeners.SendEmailListener,
		// can register more listeners for this event here ...
	},
	EventPrintTime: {
		listeners.PrintTimeListener,
	},
	// can have more mapping here ...
}
