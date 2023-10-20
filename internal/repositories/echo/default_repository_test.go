package echo

import (
	"context"
	"testing"

	"github.com/google/uuid"
	dto "github.com/sergiovirahonda/echo-api/internal/infrastructure/dtos"
	"github.com/sergiovirahonda/echo-api/internal/models"
	"gopkg.in/go-playground/assert.v1"
	"gorm.io/gorm"
)

func TestGetEchoEntityWithInvalidIDReturnsError(t *testing.T) {
	repository := NewDefaultRepository(database)
	id := uuid.New()
	entity, err := repository.Get(
		context.Background(),
		id,
	)
	assert.NotEqual(t, err, nil)
	assert.Equal(t, entity, nil)
	assert.Equal(t, err.Error(), gorm.ErrRecordNotFound.Error())
}

func TestGetEchoEntityWithValidIDReturnsExpectedEntity(t *testing.T) {
	repository := NewDefaultRepository(database)
	factory := models.NewEchoFactory()
	echo := factory.New(
		"test",
	)
	err := repository.Connection.Create(&echo).Error
	assert.Equal(t, err, nil)
	entity, err := repository.Get(
		context.Background(),
		echo.ID,
	)
	assert.Equal(t, err, nil)
	assert.Equal(t, entity.ID, echo.ID)
	assert.Equal(t, entity.Timestamp, echo.Timestamp)
	assert.Equal(t, entity.Value, "test")
	err = repository.Connection.Delete(&echo).Error
	assert.Equal(t, err, nil)
}

func TestGetAllEchoEntitiesReturnsExpectedEntities(t *testing.T) {
	repository := NewDefaultRepository(database)
	factory := models.NewEchoFactory()
	entity := factory.New(
		"test",
	)
	err := repository.Connection.Create(&entity).Error
	assert.Equal(t, err, nil)
	entities, err := repository.GetAll(
		context.Background(),
	)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(*entities), 1)
	assert.Equal(t, (*entities)[0].ID, entity.ID)
	assert.Equal(t, (*entities)[0].Timestamp, entity.Timestamp)
	assert.Equal(t, (*entities)[0].Value, "test")
	err = repository.Connection.Delete(&entity).Error
	assert.Equal(t, err, nil)
}

func TestCreateEchoEntityReturnsExpectedEntity(t *testing.T) {
	repository := NewDefaultRepository(database)
	factory := models.NewEchoFactory()
	entity := factory.New(
		"test",
	)
	err := repository.Create(
		context.Background(),
		entity,
	)
	assert.Equal(t, err, nil)
	echo := dto.Echo{}
	err = repository.Connection.Where("id = ?", entity.ID).First(&echo).Error
	assert.Equal(t, err, nil)
	assert.Equal(t, echo.ID, entity.ID)
	assert.Equal(t, echo.Timestamp, entity.Timestamp)
	assert.Equal(t, echo.Value, "test")
	err = repository.Connection.Delete(&echo).Error
	assert.Equal(t, err, nil)
}

func TestDeleteEchoEntityWithValidIDReturnsNoError(t *testing.T) {
	repository := NewDefaultRepository(database)
	factory := models.NewEchoFactory()
	entity := factory.New(
		"test",
	)
	err := repository.Connection.Create(&entity).Error
	assert.Equal(t, err, nil)
	err = repository.Delete(
		context.Background(),
		entity.ID,
	)
	assert.Equal(t, err, nil)
}

func TestDeleteEchoEntityWithInvalidIDReturnsError(t *testing.T) {
	repository := NewDefaultRepository(database)
	id := uuid.New()
	err := repository.Delete(
		context.Background(),
		id,
	)
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), gorm.ErrRecordNotFound.Error())
}
