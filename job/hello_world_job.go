package job

import (
	"errors"

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
		return errors.New("job has been executed too many times")
	}

	log.Info().Msg("Hello World")

	return nil
}

func (job *HelloWorldJob) Name() string {
	return "HelloWorldJob"
}
