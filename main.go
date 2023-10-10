package main

import (
	"time"

	"github.com/100xlongx/job_scheduler/job"
	"github.com/100xlongx/job_scheduler/scheduler"
)

func main() {
	job := job.NewHelloWorldJob(2)
	ticker := time.NewTicker(time.Millisecond * 500)

	scheduler := scheduler.New(ticker, job)

	scheduler.Start()
	time.Sleep(1600 * time.Millisecond)
	scheduler.Stop()
}
