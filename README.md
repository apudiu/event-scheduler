### What it is
This is a scheduler with support for events persistence. Following is key features
1. Schedule event for single execution (with `time.Time` or `time.Duration`)
2. Schedule event for recurring execution (with cron string ([generate here](https://crontab.guru)) or `time.Duration`)
3. Events persistence with `gorm` supported db's or your own persistent storage (create by implementing single interface)
4. Customized event payload, Use included `gob` payload or create your own payload (create by implementing single interface)

#### To run example;
- Clone repo & set your (`gorm` supported db) credentials in `example/bootstrap/db.go:21`
- Run `go run example/main.go` 

#### Installation
Install the package in your project using go
```shell
go get github.com/apudiu/event-scheduler
```

### Usage

##### Persistence driver
First a persistent driver is needed to create a scheduler. You can use existing `gormdriver` to persist events in `gorm` supported databases.
```go
import (
    "github.com/apudiu/event-scheduler/dp/driver/gormdriver"
)

var DB *gorm.DB // your existing/ created gorm connection
var persis = gormdriver.New(DB)
```
If you dont want `gorm` then you can create your own persistence driver by implementing `dp.DataPersistent` interface.
For better understanding on how to do it, can check `gormdriver` implementation

##### Scheduler creation
One you've persistence driver use that to create a scheduler
```go
import (
    scheduler "github.com/apudiu/event-scheduler"
)

// ListenersList register listeners to events 
// NOTE: here we're registering listeners to events
var listenersMap = scheduler.Listeners{
    "EventSendEmail": { // event name
        listeners.SendEmailListener, // listener function which is type of scheduler.ListenFunc
        // can register more listeners for this event here ...
    },
    "EventPrintTime": {
        listeners.PrintTimeListener,
    },
    // can have more mapping here ...
}

// create new scheduler
var SCD = scheduler.NewScheduler(s, listenersMap)
```

##### Event creation
Once the scheduler is ready, we can create & dispatch events like following
```go
import (
    "github.com/apudiu/event-scheduler/event/generator"
    "github.com/apudiu/event-scheduler/example/bootstrap"
    "github.com/apudiu/event-scheduler/example/eventListenersMap"
)

// create the event using generator
var SendEmailEvent = func() *generator.DelayedEvent {
	return generator.NewDelayedEvent(eventListenersMap.EventSendEmail, bootstrap.SCD)
}

// later dispatch the event
if err := SendEmailEvent().Dispatch(p, time.Now().Add(time.Second*5)); err != nil {
    log.Fatal("event dispatch failed: ", err)
}
```
`generator` provides easy way to create event. You can create event without using generators too.
First create the event payload
```go
func SendEmailEvent(email string, content string) {

	// crete event payload
	p := payload.NewGobPayload(
		fmt.Sprintf("Send email to %s with content %s", email, content),
	)

	// send the event using the payload, or make a way to send it later, for this take a look at generators (like: generator.NewDelayedEvent)
	bootstrap.SCD.Schedule("SendEmail", p, time.Now().Add(time.Second*5))
}
```
##### Listener creation
```go
var SendEmailListener scheduler.ListenFunc = func(data payload.TransferablePayload) {

	// decode payload to whatever you want (ideally to that it was encoded & sent to the event)
	// like: if you sent a User (struct) then you may want to decode to that

	var decoded string
	if err := data.Unmarshal(&decoded); err != nil {
		log.Printf("Data unmarshalling failed, value: %v\n", data)
	}

	// do whatever you need with the decoded payload

	log.Printf("Listener executing with value: %#v \n", decoded)
}
```
Once the listener(s) is ready for an event, you need to register that to the scheduler (when creating scheduler).
We've registered couple of listeners with events in **Scheduler creation** section above

If something is still not clear for you, simply check **example** directory.

#### State and roadmap
This lib is in it's very early stage. Theres no tests, some parts of the pkg can be more idiomatic etc.
You get the idea that this package does exactly what I need now.

Following is current roadmap from my side.
1. Add tests
2. Add redis `DataPersistence`
3. Add more payloads `TransferablePayload`
4. Add benchmarks

I'll do this over time when I get time for it.

#### Contribution
Any kind of contribution is welcome. Just fork the repo, do your changes & submit PR. I'll try to respond as soon as I can.