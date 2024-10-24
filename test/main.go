package main

import (
	"context"
	"github.com/apudiu/event-scheduler/customevents"
	"github.com/apudiu/event-scheduler/db/driver/gormdriver"
	"github.com/apudiu/event-scheduler/scheduler"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
	"time"
)

var eventListeners = scheduler.Listeners{
	"SendEmail": {customevents.SendEmail},
	"PayBills":  {customevents.PayBills},
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	var db *gorm.DB

	s := gormdriver.New(db)
	scd := scheduler.NewScheduler(s, eventListeners)

	stopCron := scd.StartCron()
	defer stopCron()

	scd.CheckEventsInInterval(ctx, time.Minute)

	scd.Schedule("SendEmail", "mail: nilkantha.dipesh@gmail.com", time.Now().Add(1*time.Minute))
	scd.Schedule("PayBills", "paybills: $4,000 bill", time.Now().Add(2*time.Minute))

	scd.ScheduleCron("SendEmail", "mail: dipesh.dulal+new@wesionary.team", "* * * * *")

	go func() {
		for range interrupt {
			log.Println("\n❌ Interrupt received closing...")
			cancel()
		}
	}()

	<-ctx.Done()
}
