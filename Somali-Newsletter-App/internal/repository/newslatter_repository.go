package repository

import (
	"Somali-Newsletter-App/internal/entity"

	"github.com/gofrs/uuid"
)

// Repository is a generic interface for a data access layer.
type newsletterRepository interface {
	Create(ContentCreator *entity.Newsletter) error
	Update(ContentCreator *entity.Newsletter) error
	GetAll() ([]*entity.Newsletter, error)
	GetByID(ContentCreatorID uuid.UUID) (*entity.Newsletter, error)
	Delete(ContentCreatorID uuid.UUID) error
}
