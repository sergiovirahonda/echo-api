package echo

import (
	"context"

	"github.com/google/uuid"
	dto "github.com/sergiovirahonda/echo-api/internal/infrastructure/dtos"
	"github.com/sergiovirahonda/echo-api/internal/models"
	"gorm.io/gorm"
)

// Repository

type DefaultRepository struct {
	Connection *gorm.DB
}

// Repository factories

func NewDefaultRepository(connection *gorm.DB) *DefaultRepository {
	return &DefaultRepository{
		Connection: connection,
	}
}

// Repository receivers

func (r *DefaultRepository) Get(
	ctx context.Context,
	id uuid.UUID,
) (*models.Echo, error) {
	var echo dto.Echo
	err := r.Connection.Where("id = ?", id).First(&echo).Error
	if err != nil {
		return nil, err
	}
	return echo.ToEntity(), nil
}

func (r *DefaultRepository) GetAll(
	ctx context.Context,
) (*models.Echos, error) {
	var echos dto.Echos
	err := r.Connection.Find(&echos).Error
	if err != nil {
		return nil, err
	}
	return echos.ToEntities(), nil
}

func (r *DefaultRepository) Create(
	ctx context.Context,
	e *models.Echo,
) error {
	echo := dto.Echo{
		ID:        e.ID,
		Timestamp: e.Timestamp,
		Value:     e.Value,
	}
	err := r.Connection.Create(&echo).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *DefaultRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	instance := dto.Echo{}
	result := r.Connection.Where("id = ?", id).First(&instance)
	if result.Error != nil {
		return gorm.ErrRecordNotFound
	}
	r.Connection.Delete(&instance)
	return nil
}
