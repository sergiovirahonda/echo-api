package echos

import (
	"context"

	"github.com/google/uuid"
	"github.com/sergiovirahonda/echo-api/internal/models"
)

type Service interface {
	Get(ctx context.Context, id uuid.UUID) (*models.Echo, error)
	GetAll(ctx context.Context) (*models.Echos, error)
	Create(ctx context.Context, e *models.Echo) error
	CreateFromRequest(ctx context.Context, r *models.EchoRequest) (*models.Echo, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
