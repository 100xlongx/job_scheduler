package scheduler

import (
	"fmt"
	"time"

	"github.com/100xlongx/job_scheduler/job"
	"github.com/rs/zerolog/log"
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
	log.Info().Msg("Starting scheduler")

	go func() {
		for {
			select {
			case <-scheduler.doneChannel:
				return
			case t := <-scheduler.ticker.C:
				log.Info().Time("time", t).Str("jobName", scheduler.job.Name()).Msg("Executing job")
				err := scheduler.job.Execute()

				if err != nil {
					log.Error().Err(err).Msg("Error executing job")
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
