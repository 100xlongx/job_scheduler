package job

import (
	"github.com/100xlongx/job_scheduler/custom_error"
	"github.com/rs/zerolog/log"
)

type HelloWorldJob struct {
	count    int
	maxCount int
}

func NewHelloWorldJob(maxCount int) *HelloWorldJob {
	return &HelloWorldJob{
		count:    0,
		maxCount: maxCount,
	}
}

func (job *HelloWorldJob) Execute() error {
	job.count++
	if job.count > job.maxCount {
		// return custom_error.NewFatalError("job has been executed too many times")
		return custom_error.NewWarnError("job has been executed too many times")
	}

	log.Info().Msg("Hello World")

	return nil
}

func (job *HelloWorldJob) Name() string {
	return "HelloWorldJob"
}
