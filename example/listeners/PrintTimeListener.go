package listeners

import (
	scheduler "github.com/apudiu/event-scheduler"
	"github.com/apudiu/event-scheduler/event/payload"
	"log"
	"time"
)

var PrintTimeListener scheduler.ListenFunc = func(d payload.TransferablePayload) {

	// we do not need data in this case

	log.Printf("Listener executing on event: %s with value: %#v \n", d.EventName(), time.Now().Format(time.RFC850))
}
