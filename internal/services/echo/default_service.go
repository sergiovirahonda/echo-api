package echos

import (
	"context"

	"github.com/google/uuid"
	"github.com/sergiovirahonda/echo-api/internal/models"
	"github.com/sergiovirahonda/echo-api/internal/repositories/echo"
)

type DefaultService struct {
	Repository echo.Repository
}

func NewDefaultService(repository echo.Repository) *DefaultService {
	return &DefaultService{
		Repository: repository,
	}
}

// Echo service receivers

func (s *DefaultService) Get(
	ctx context.Context,
	id uuid.UUID,
) (*models.Echo, error) {
	return s.Repository.Get(ctx, id)
}

func (s *DefaultService) GetAll(
	ctx context.Context,
) (*models.Echos, error) {
	return s.Repository.GetAll(ctx)
}

func (s *DefaultService) Create(
	ctx context.Context,
	e *models.Echo,
) error {
	return s.Repository.Create(ctx, e)
}

func (s *DefaultService) CreateFromRequest(
	ctx context.Context,
	r *models.EchoRequest,
) (*models.Echo, error) {
	err := r.Validate()
	if err != nil {
		return nil, err
	}
	factory := models.EchoFactory{}
	entity := factory.New(
		r.Value,
	)
	err = s.Create(ctx, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (s *DefaultService) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	return s.Repository.Delete(ctx, id)
}
