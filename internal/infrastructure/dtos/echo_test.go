package dto

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"gopkg.in/go-playground/assert.v1"
)

func TestEchoToEntityReturnsExpectedEntity(t *testing.T) {
	d := Echo{
		ID:        uuid.New(),
		Timestamp: time.Now().UTC(),
		Value:     "test",
	}
	entity := d.ToEntity()
	assert.Equal(t, d.ID, entity.ID)
	assert.Equal(t, d.Timestamp, entity.Timestamp)
	assert.Equal(t, d.Value, entity.Value)
}

func TestEchosToEntitiesReturnsExpectedEntities(t *testing.T) {
	d1 := Echo{
		ID:        uuid.New(),
		Timestamp: time.Now().UTC(),
		Value:     "test1",
	}
	d2 := Echo{
		ID:        uuid.New(),
		Timestamp: time.Now().UTC(),
		Value:     "test2",
	}
	dtos := Echos{d1, d2}
	entities := dtos.ToEntities()
	assert.Equal(t, d1.ID, (*entities)[0].ID)
	assert.Equal(t, d1.Timestamp, (*entities)[0].Timestamp)
	assert.Equal(t, d1.Value, (*entities)[0].Value)
	assert.Equal(t, d2.ID, (*entities)[1].ID)
	assert.Equal(t, d2.Timestamp, (*entities)[1].Timestamp)
	assert.Equal(t, d2.Value, (*entities)[1].Value)
}
