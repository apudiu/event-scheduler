package listeners

import (
	scheduler "github.com/apudiu/event-scheduler"
	"github.com/apudiu/event-scheduler/event/payload"
	"log"
)

var SendEmailListener scheduler.ListenFunc = func(data payload.TransferablePayload) {
	var decoded string
	if err := data.Unmarshal(&decoded); err != nil {
		log.Printf("Data unmarshalling failed, value: %v\n", data)
	}
}
