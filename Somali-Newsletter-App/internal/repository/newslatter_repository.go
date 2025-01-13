package repository

import (
	"Somali-Newsletter-App/internal/entity"

	"github.com/gofrs/uuid"
)

// Repository is a generic interface for a data access layer.
type NewsletterRepository interface {
	Create(newsletter *entity.Newsletter) error
	Update(newsLatter *entity.Newsletter) error
	GetAll() ([]*entity.Newsletter, error)
	GetByID(newsletterID uuid.UUID) (*entity.Newsletter, error)
	Delete(NewsletterID uuid.UUID) error
}


