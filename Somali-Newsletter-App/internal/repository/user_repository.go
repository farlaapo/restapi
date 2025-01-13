package repository

import (
	"Somali-Newsletter-App/internal/entity"

	"github.com/gofrs/uuid"
)

// UserRepository is a generic interface for a data access layer.
type UserRepository interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	GetAll() ([]*entity.User, error)
	GetByID(userID uuid.UUID) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Delete(userID uuid.UUID) error
}