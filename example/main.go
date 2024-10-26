package main

import (
	"context"
	"fmt"
	"github.com/apudiu/event-scheduler/event/payload"
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

	p := payload.NewGobPayload(
		fmt.Sprintf("Send email to %s with content %s", "nilkantha.dipesh@gmail.com", "CONTENT"),
	)
	if err := events.SendEmailEvent().Dispatch(p, time.Now().Add(time.Second*5)); err != nil {
		log.Fatal("event dispatch failed: ", err)
	}

	if err := events.PrintTimeEvent().DispatchDur(payload.NewGobPayload(""), time.Second*7); err != nil {
		log.Fatal("event dispatch failed: ", err)
	}

	go func() {
		for range interrupt {
			log.Println("\n❌ Interrupt received closing...")
			cancel()
		}
	}()

	<-ctx.Done()
}
