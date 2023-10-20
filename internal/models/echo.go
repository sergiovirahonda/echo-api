package models

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/sergiovirahonda/echo-api/internal/errors"
)

// Errors

var (
	validate = validator.New()
)

type EchoResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Value     string    `json:"value"`
}

type EchoRequest struct {
	Value string `json:"value" validate:"required"`
}

type Echo struct {
	ID        uuid.UUID `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Value     string    `json:"value"`
}

type Echos []Echo

// Receivers

func (e *EchoRequest) Validate() error {
	err := validate.Struct(e)
	if err != nil {
		return errors.InvalidEchoRequestError
	}
	if e.Value == "" {
		return errors.InvalidValueError
	}
	if len(e.Value) > 255 {
		return errors.InvalidValueError
	}
	return nil
}

// Factories

type EchoFactory struct{}

func NewEchoFactory() *EchoFactory {
	return &EchoFactory{}
}

// Factory receivers

func (e *EchoFactory) New(value string) *Echo {
	return &Echo{
		ID:        uuid.New(),
		Timestamp: time.Now().UTC(),
		Value:     value,
	}
}
