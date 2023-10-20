package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/sergiovirahonda/echo-api/internal/models"
	"gorm.io/gorm"
)

type Echo struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;"`
	Timestamp time.Time `gorm:"type:timestamp;not null"`
	Value     string    `gorm:"type:varchar(255);not null"`
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
