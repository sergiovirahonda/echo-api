package echos

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/sergiovirahonda/echo-api/internal/models"
	"github.com/sergiovirahonda/echo-api/internal/repositories/echo"
	"gopkg.in/go-playground/assert.v1"
	"gorm.io/gorm"
)

func TestGetReturnsErrIfIdDoesNotExist(t *testing.T) {
	repository := echo.NewDefaultRepository(database)
	service := NewDefaultService(repository)
	_, err := service.Get(
		context.Background(),
		uuid.New(),
	)
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), gorm.ErrRecordNotFound.Error())
}

func TestGetReturnsExpectedEntity(t *testing.T) {
	repository := echo.NewDefaultRepository(database)
	service := NewDefaultService(repository)
	factory := models.NewEchoFactory()
	entity := factory.New(
		"test",
	)
	err := repository.Connection.Create(&entity).Error
	assert.Equal(t, err, nil)
	result, err := service.Get(
		context.Background(),
		entity.ID,
	)
	assert.Equal(t, err, nil)
	assert.Equal(t, result.ID, entity.ID)
	assert.Equal(t, result.Timestamp, entity.Timestamp)
	assert.Equal(t, result.Value, "test")
	err = repository.Connection.Delete(&entity).Error
	assert.Equal(t, err, nil)
}

func TestGetAllReturnsExpectedEntities(t *testing.T) {
	repository := echo.NewDefaultRepository(database)
	service := NewDefaultService(repository)
	factory := models.NewEchoFactory()
	entity := factory.New(
		"test",
	)
	err := repository.Connection.Create(&entity).Error
	assert.Equal(t, err, nil)
	entities, err := service.GetAll(
		context.Background(),
	)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(*entities), 1)
	assert.Equal(t, (*entities)[0].ID, entity.ID)
	assert.Equal(t, (*entities)[0].Timestamp, entity.Timestamp)
	err = repository.Connection.Delete(&entity).Error
	assert.Equal(t, err, nil)
}

func TestCreateReturnsExpectedEntity(t *testing.T) {
	repository := echo.NewDefaultRepository(database)
	service := NewDefaultService(repository)
	factory := models.NewEchoFactory()
	entity := factory.New(
		"test",
	)
	err := service.Create(
		context.Background(),
		entity,
	)
	assert.Equal(t, err, nil)
	result, err := service.Get(
		context.Background(),
		entity.ID,
	)
	assert.Equal(t, err, nil)
	assert.Equal(t, result.ID, entity.ID)
	assert.Equal(t, result.Timestamp, entity.Timestamp)
	assert.Equal(t, result.Value, "test")
	err = repository.Connection.Delete(&entity).Error
	assert.Equal(t, err, nil)
}

func TestCreateFromRequestReturnsExpectedEntity(t *testing.T) {
	repository := echo.NewDefaultRepository(database)
	service := NewDefaultService(repository)
	request := models.EchoRequest{
		Value: "test",
	}
	result, err := service.CreateFromRequest(
		context.Background(),
		&request,
	)
	assert.Equal(t, err, nil)
	assert.Equal(t, result.Timestamp, result.Timestamp)
	assert.Equal(t, result.Value, "test")
	err = repository.Connection.Delete(&result).Error
	assert.Equal(t, err, nil)
}

func TestDeleteReturnsErrIfIdDoesNotExist(t *testing.T) {
	repository := echo.NewDefaultRepository(database)
	service := NewDefaultService(repository)
	err := service.Delete(
		context.Background(),
		uuid.New(),
	)
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), gorm.ErrRecordNotFound.Error())
}

func TestDeleteDeletesEntity(t *testing.T) {
	repository := echo.NewDefaultRepository(database)
	service := NewDefaultService(repository)
	factory := models.NewEchoFactory()
	entity := factory.New(
		"test",
	)
	err := repository.Connection.Create(&entity).Error
	assert.Equal(t, err, nil)
	err = service.Delete(
		context.Background(),
		entity.ID,
	)
	assert.Equal(t, err, nil)
	_, err = service.Get(
		context.Background(),
		entity.ID,
	)
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), gorm.ErrRecordNotFound.Error())
}
