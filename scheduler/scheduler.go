package scheduler

import (
	"fmt"
	"time"

	"github.com/100xlongx/job_scheduler/custom_error"
	"github.com/100xlongx/job_scheduler/job"
	"github.com/rs/zerolog/log"
)

type Scheduler struct {
	job          job.Job
	ticker       *time.Ticker
	doneChannel  chan bool
	errorChannel chan error
}

func New(ticker *time.Ticker, job job.Job) *Scheduler {
	return &Scheduler{
		job:          job,
		ticker:       ticker,
		doneChannel:  make(chan bool),
		errorChannel: make(chan error),
	}
}

func (scheduler *Scheduler) Start() error {
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
					scheduler.errorChannel <- err
				}
			}
		}
	}()

	scheduler.ListenForErrors()

	return nil
}

func (scheduler *Scheduler) ListenForErrors() {
	log.Info().Msg("Starting to listen for errors")

	go func() {
		for err := range scheduler.errorChannel {
			switch e := err.(type) {
			case *custom_error.FatalError:
				log.Fatal().Err(e).Msg("Fatal error encountered, shutting down scheduler")
				scheduler.Stop()
			default:
				log.Error().Err(e).Msg("Received an error from the scheduler")
			}

		}
	}()
}

func (scheduler *Scheduler) Stop() error {
	fmt.Println("Stopping scheduler")
	scheduler.ticker.Stop()
	scheduler.doneChannel <- true

	return nil
}
