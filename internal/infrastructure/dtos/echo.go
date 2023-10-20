package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/sergiovirahonda/echo-api/internal/models"
	"github.com/uptrace/bun"
)

type Echo struct {
	bun.BaseModel `bun:"table:echo,alias:echo"`
	ID            uuid.UUID `bun:"id,type:uuid"`
	Timestamp     time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	Value         string    `bun:"value,type:text,notnull"`
}

type Echos []Echo

// Receivers

func (e *Echo) ToEntity() *models.Echo {
	return &models.Echo{
		ID:        e.ID,
		Timestamp: e.Timestamp,
		Value:     e.Value,
	}
}

func (e *Echos) ToEntities() *models.Echos {
	entities := make(models.Echos, len(*e))
	for i, v := range *e {
		entities[i] = *v.ToEntity()
	}
	return &entities
}
