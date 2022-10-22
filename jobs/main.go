package job

import (
	"time"

	"github.com/go-co-op/gocron"
)

var Scheduler *gocron.Scheduler = gocron.NewScheduler(time.UTC)

func RegisterJobs() {
	RegisterEndpointJob()
	go Scheduler.StartAsync()
}
