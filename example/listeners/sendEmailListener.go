package listeners

import (
	scheduler "github.com/apudiu/event-scheduler"
	"github.com/apudiu/event-scheduler/event/payload"
	"log"
)

var SendEmailListener scheduler.ListenFunc = func(data payload.TransferablePayload) {

	// decode payload to whatever you want (ideally to that it was encoded & sent to the event)
	// like: if you sent a User (struct) then try to decode to that

	var decoded string
	if err := data.Unmarshal(&decoded); err != nil {
		log.Printf("Data unmarshalling failed on event: %s, value: %v\n", data.EventName(), data)
	}

	// do whatever you need with the decoded payload

	log.Printf("Listener executing with on event: %s with value: %#v \n", data.EventName(), decoded)
}
