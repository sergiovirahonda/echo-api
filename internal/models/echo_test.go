package models

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"gopkg.in/go-playground/assert.v1"
)

func TestEchoRequestValidationSuccessScenario(t *testing.T) {
	echo := EchoRequest{
		Value: "test",
	}
	err := echo.Validate()
	if err != nil {
		t.Errorf("Echo validation failed: %v", err)
	}
}

func TestEchoValidationFailsScenario1(t *testing.T) {
	echo := EchoRequest{
		Value: "",
	}
	err := echo.Validate()
	if err == nil {
		t.Errorf("Echo validation should have failed")
	}
}

func TestEchoValidationFailsScenario2(t *testing.T) {
	var value string
	for i := 0; i < 256; i++ {
		value += "a"
	}
	echo := EchoRequest{
		Value: value,
	}
	err := echo.Validate()
	assert.NotEqual(t, err, nil)
}

func TestEchoEntityToResponseSucceeds(t *testing.T) {
	timestamp := time.Now().UTC()
	echo := Echo{
		ID:        uuid.New(),
		Timestamp: timestamp,
		Value:     "test-1",
	}
	response := echo.ToResponse()
	assert.Equal(t, response.Time, echo.Timestamp)
	assert.Equal(t, response.Value, echo.Value)
}

func TestEchoEntitiesToResponsesSucceeds(t *testing.T) {
	timestamp := time.Now().UTC()
	echo1 := Echo{
		ID:        uuid.New(),
		Timestamp: timestamp,
		Value:     "test-1",
	}
	echo2 := Echo{
		ID:        uuid.New(),
		Timestamp: timestamp,
		Value:     "test-2",
	}
	entities := Echos{echo1, echo2}
	responses := entities.ToResponses()
	assert.Equal(t, len(responses.Echos), 2)
	assert.Equal(t, responses.Echos[0].Time, echo1.Timestamp)
	assert.Equal(t, responses.Echos[0].Value, echo1.Value)
	assert.Equal(t, responses.Echos[1].Time, echo2.Timestamp)
	assert.Equal(t, responses.Echos[1].Value, echo2.Value)
}

func TestEchoToResponseFromRequestSucceeds(t *testing.T) {
	echo := Echo{
		Value: "test",
	}
	response := echo.ToResponseFromRequest()
	assert.Equal(t, response.Value, "test echo")
}

func TestEchoFactorySucceeds(t *testing.T) {
	factory := NewEchoFactory()
	echo := factory.New("test")
	assert.Equal(t, echo.Value, "test")
	assert.NotEqual(t, echo.ID, uuid.Nil)
	assert.NotEqual(t, echo.Timestamp, time.Time{})
}
