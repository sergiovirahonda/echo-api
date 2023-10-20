package echo

import (
	"context"

	"github.com/google/uuid"
	"github.com/sergiovirahonda/echo-api/internal/models"
)

type Repository interface {
	Get(ctx context.Context, id uuid.UUID) (*models.Echo, error)
	GetAll(ctx context.Context) (*models.Echos, error)
	Find(ctx context.Context, filters map[string]interface{}) (*models.Echos, error)
	Create(ctx context.Context, echo *models.Echo) (*models.Echo, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
