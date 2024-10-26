package events

import (
	"fmt"
	"github.com/apudiu/event-scheduler/event/payload"
	"github.com/apudiu/event-scheduler/example/bootstrap"
	"time"
)

// receive whatever you need in args

func SendEmailEvent(email string, content string) {

	// prepare payload with provided payload.GobPayload
	// or create one yourself by implementing payload.TransferablePayload interface
	p := payload.NewGobPayload(
		fmt.Sprintf("Send email to %s with content %s", email, content),
	)

	// send the event using the payload
	bootstrap.SCD.Schedule("SendEmail", p, time.Now().Add(time.Second*5))
}
