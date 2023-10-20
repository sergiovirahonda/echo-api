package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/sergiovirahonda/echo-api/internal/errors"
)

// Errors

var (
	validate = validator.New()
)

// HTTP models

type EchoResponse struct {
	Time  time.Time `json:"time"`
	Value string    `json:"value"`
}

type EchoResponses struct {
	Echos []EchoResponse `json:"whats-echoed"`
}

type EchoRequest struct {
	Value string `json:"echo-me" validate:"required"`
}

type EchoResponseFromRequest struct {
	Value string `json:"echo-you"`
}

// Entities

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

func (e *Echo) ToResponse() *EchoResponse {
	return &EchoResponse{
		Time:  e.Timestamp,
		Value: e.Value,
	}
}

func (e *Echos) ToResponses() *EchoResponses {
	var responses []EchoResponse
	for _, echo := range *e {
		responses = append(responses, *echo.ToResponse())
	}
	return &EchoResponses{
		Echos: responses,
	}
}

func (e *Echo) ToResponseFromRequest() *EchoResponseFromRequest {
	return &EchoResponseFromRequest{
		Value: fmt.Sprintf("%s echo", e.Value),
	}
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
