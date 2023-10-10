package job

type Job interface {
	Name() string
	Execute() error
}
