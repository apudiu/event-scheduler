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

var SCD = scheduler.NewScheduler(s, listenersMap)
```
todo: complete readme...