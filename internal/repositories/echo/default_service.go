package echo

import (
	"context"
	"strings"

	"github.com/google/uuid"
	dto "github.com/sergiovirahonda/echo-api/internal/infrastructure/dtos"
	"github.com/sergiovirahonda/echo-api/internal/models"
	"github.com/uptrace/bun"
)

// Repository

type DefaultRepository struct {
	Connection *bun.DB
}

// Repository factories

func NewDefaultRepository(connection *bun.DB) *DefaultRepository {
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
	err := r.Connection.NewSelect().
		Model(&echo).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return echo.ToEntity(), nil
}

func (r *DefaultRepository) GetAll(
	ctx context.Context,
) (*models.Echos, error) {
	var echos dto.Echos
	err := r.Connection.NewSelect().
		Model(&echos).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return echos.ToEntities(), nil
}

func (r *DefaultRepository) Find(
	ctx context.Context,
	filters map[string]interface{},
) (*models.Echos, error) {
	var instances dto.Echos
	query := r.Connection.NewSelect().Model(&instances)
	for key, value := range filters {
		if strings.HasSuffix(key, "__gt") {
			key = strings.ReplaceAll(key, "__gt", "")
			query = query.Where(key+" > ?", value)
		} else if strings.HasSuffix(key, "__gte") {
			key = strings.ReplaceAll(key, "__gte", "")
			query = query.Where(key+" >= ?", value)
		} else if strings.HasSuffix(key, "__lt") {
			key = strings.ReplaceAll(key, "__lt", "")
			query = query.Where(key+" < ?", value)
		} else if strings.HasSuffix(key, "__lte") {
			key = strings.ReplaceAll(key, "__lte", "")
			query = query.Where(key+" <= ?", value)
		} else if strings.HasSuffix(key, "__like") {
			key = strings.ReplaceAll(key, "__like", "")
			query = query.Where(key+" LIKE ?", value)
		} else {
			query = query.Where(key+" = ?", value)
		}
	}
	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}
	return instances.ToEntities(), nil
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
	_, err := r.Connection.NewInsert().
		Model(&echo).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *DefaultRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	_, err := r.Connection.NewDelete().
		Model(&dto.Echo{}).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
