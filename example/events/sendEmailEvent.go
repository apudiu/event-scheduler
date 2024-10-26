package events

import (
	"fmt"
	"github.com/apudiu/event-scheduler/event/payload"
	"github.com/apudiu/event-scheduler/example/bootstrap"
	"time"
)

func SendEmailEvent(email string, content string) {
	p := payload.NewGobPayload(
		fmt.Sprintf("Send email to %s with content %s", email, content),
	)
	bootstrap.SCD.Schedule("SendEmail", p, time.Now().Add(time.Second*5))
}
