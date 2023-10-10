package custom_error

type CustomError interface {
	error
	Reason() string
}

type FatalError struct {
	Msg string
}

func (e *FatalError) Error() string {
	return e.Msg
}

func (e *FatalError) Reason() string {
	return e.Msg
}

func NewFatalError(msg string) *FatalError {
	return &FatalError{Msg: msg}
}

type WarnError struct {
	Msg string
}

func (e *WarnError) Error() string {
	return e.Msg
}

func (e *WarnError) Reason() string {
	return e.Msg
}

func NewWarnError(msg string) *WarnError {
	return &WarnError{Msg: msg}
}

type InfoError struct {
	Msg string
}

func (e *InfoError) Error() string {
	return e.Msg
}

func (e *InfoError) Reason() string {
	return e.Msg
}

func NewInfoError(msg string) *InfoError {
	return &InfoError{Msg: msg}
}
