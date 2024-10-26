package listeners

import (
	scheduler "github.com/apudiu/event-scheduler"
	"github.com/apudiu/event-scheduler/event/payload"
	"log"
	"time"
)

var PrintTimeListener scheduler.ListenFunc = func(_ payload.TransferablePayload) {

	// we do not need data in this case

	log.Printf("Listener executing with value: %#v \n", time.Now().Format(time.RFC850))
}
