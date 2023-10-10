package main

import (
	"os"
	"time"

	"github.com/100xlongx/job_scheduler/job"
	"github.com/100xlongx/job_scheduler/scheduler"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	output := zerolog.ConsoleWriter{Out: os.Stderr}

	log.Logger = log.Output(output)
}

func ListenForFeedback(s *scheduler.Scheduler) {
	feedback := s.Feedback()

	go func() {
		for msg := range feedback {
			log.Info().Msg(msg)
		}
	}()
}

func main() {
	InitLogger()

	job := job.NewHelloWorldJob(2)
	ticker := time.NewTicker(time.Millisecond * 500)

	scheduler := scheduler.New(ticker, job)

	if err := scheduler.Start(); err != nil {
		log.Error().Err(err).Msg("Failed to start the scheduler")
		return
	}

	ListenForFeedback(scheduler)

	time.Sleep(1 * time.Minute)

	if err := scheduler.Stop(); err != nil {
		log.Error().Err(err).Msg("Failed to stop the scheduler")
		return
	}
}
