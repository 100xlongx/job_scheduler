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

func main() {
	InitLogger()

	job := job.NewHelloWorldJob(2)
	ticker := time.NewTicker(time.Millisecond * 500)

	scheduler := scheduler.New(ticker, job)

	scheduler.Start()
	time.Sleep(1 * time.Minute)
	scheduler.Stop()
}
