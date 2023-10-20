package errors

import (
	"errors"
)

func Is(err, target error) bool {
	return errors.Is(err, target)
}
func New(msg string) error {
	return errors.New(msg)
}

type (
	BadRequest struct {
		err error
	}
)

func (e *BadRequest) Error() string {
	return e.err.Error()
}

func (e *BadRequest) Is(err error) bool {
	_, ok := err.(*BadRequest)
	return ok
}

func NewBadRequest(msg string) *BadRequest {
	return &BadRequest{New(msg)}
}

// Custom errors

var (
	// Echo errors
	InvalidEchoRequestError = errors.New("invalid echo request")
	InvalidValueError       = errors.New("invalid value")
)
