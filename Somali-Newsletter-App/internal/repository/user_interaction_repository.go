package repository

import (
	"Somali-Newsletter-App/internal/entity"

	"github.com/gofrs/uuid"
)

// Repository is a generic interface for data access.
type UserInteractionRepository interface {
	Create(userInteraction *entity.UserInteraction) error
	Update(userInteraction *entity.UserInteraction) error
	Delete(userInteractionID uuid.UUID) error
	Get(userInteractionID uuid.UUID) (*entity.UserInteraction, error)
	GetAll() ([]*entity.UserInteraction, error)
}

