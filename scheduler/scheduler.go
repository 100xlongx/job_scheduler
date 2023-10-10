package scheduler

import (
	"fmt"
	"time"

	"github.com/100xlongx/job_scheduler/job"
)

type Scheduler struct {
	job         job.Job
	ticker      *time.Ticker
	doneChannel chan bool
}

func New(ticker *time.Ticker, job job.Job) *Scheduler {
	return &Scheduler{
		job:         job,
		ticker:      ticker,
		doneChannel: make(chan bool),
	}
}

func (scheduler *Scheduler) Start() {
	fmt.Println("Starting scheduler")

	go func() {
		for {
			select {
			case <-scheduler.doneChannel:
				return
			case t := <-scheduler.ticker.C:
				fmt.Println("Executing ticket at", t)
				err := scheduler.job.Execute()

				if err != nil {
					fmt.Println("Error executing job", err)
				}
			}
		}
	}()
}

func (scheduler *Scheduler) Stop() {
	fmt.Println("Stopping scheduler")
	scheduler.ticker.Stop()
	scheduler.doneChannel <- true
}
