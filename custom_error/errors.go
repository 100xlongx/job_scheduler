package custom_error

type FatalError struct {
	Reason string
}

func (e *FatalError) Error() string {
	return e.Reason
}

func NewFatalError(reason string) *FatalError {
	return &FatalError{Reason: reason}
}
