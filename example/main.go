package main

import (
	"context"
	"github.com/apudiu/event-scheduler/example/bootstrap"
	"github.com/apudiu/event-scheduler/example/events"
	"log"
	"os"
	"os/signal"
	"time"
)

func init() {
	// init db conn
	bootstrap.ConnectDB()

	// init scheduler
	bootstrap.InitScheduler()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	stopCron := bootstrap.SCD.StartCron()
	defer stopCron()

	bootstrap.SCD.CheckEventsInInterval(ctx, time.Second*5)

	//bootstrap.SCD.Schedule("SendEmail", "mail: nilkantha.dipesh@gmail.com", time.Now().Add(1*time.Minute))
	//bootstrap.SCD.Schedule("PayBills", "paybills: $4,000 bill", time.Now().Add(2*time.Minute))
	//
	//bootstrap.SCD.ScheduleRecurring("SendEmail", "mail: dipesh.dulal+new@wesionary.team", "* * * * *")

	events.SendEmailEvent("nilkantha.dipesh@gmail.com", "thank you!")

	go func() {
		for range interrupt {
			log.Println("\n❌ Interrupt received closing...")
			cancel()
		}
	}()

	<-ctx.Done()
}
