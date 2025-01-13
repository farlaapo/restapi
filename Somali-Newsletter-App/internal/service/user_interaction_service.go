package service

import (
	"Somali-Newsletter-App/internal/entity"
	"Somali-Newsletter-App/internal/repository"
	"time"

	"github.com/gofrs/uuid"
)

// UserInteractionService is a service for user interactions interface
type UserInteractionService interface {
	CreateUserInteraction(NewsletterID, UserID uuid.UUID, liked, read bool, comment string) (*entity.UserInteraction, error)
	GetUserInteractionByID(userInteractionID uuid.UUID) (*entity.UserInteraction, error)
	GetAllUserInteraction() ([]*entity.UserInteraction, error)
	UpdateUserInteraction(userInteraction *entity.UserInteraction) error
	DeleteUserInteraction(userInteractionID uuid.UUID) error
}

// userInteractionService is a service for user interactions struct
type userInteractionService struct {
	repo      repository.UserInteractionRepository
	tokenRepo repository.TokenRepository
}

// CreateUserInteraction implements UserInteractionService.
func (s *userInteractionService) CreateUserInteraction(NewsletterID uuid.UUID, UserID uuid.UUID, liked bool, read bool, comment string) (*entity.UserInteraction, error) {
	// generate uuid
	newUUID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	 userInteraction := &entity.UserInteraction{
		ID:          newUUID,
		NewsletterID: NewsletterID,
		UserID:      UserID,
		Liked:       liked,
		Read:        read,
		Comment:     comment,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	 }

	if err := s.repo.Create(userInteraction); err != nil {
		return nil, err
	}
	return userInteraction, nil

}

// DeleteUserInteraction implements UserInteractionService.
func (s *userInteractionService) DeleteUserInteraction(userInteractionID uuid.UUID) error {
	_, err := s.repo.Get(userInteractionID)
	if err != nil {
		return  err
	}

	if err := s.repo.Delete(userInteractionID); err != nil {
		return err
	}

	return nil
}

// GetAllUserInteraction implements UserInteractionService.
func (s *userInteractionService) GetAllUserInteraction() ([]*entity.UserInteraction, error) {
	userInteraction, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return userInteraction, nil

}

// GetUserInteraction implements UserInteractionService.
func (s *userInteractionService) GetUserInteractionByID(userInteractionID uuid.UUID) (*entity.UserInteraction, error) {
	userInteraction, err := s.repo.Get(userInteractionID)
	if err != nil {
		return nil, err
	}
	return userInteraction, nil
}

// UpdateUserInteraction implements UserInteractionService.
func (s *userInteractionService) UpdateUserInteraction(userInteraction *entity.UserInteraction) error {
	_, err := s.repo.Get(userInteraction.ID)
	if err != nil {
		return err
	}

	if err := s.repo.Update(userInteraction); err != nil {
		return err
	}
	return nil
}

// newuserInteractionService return
func NewuserInteractionService(userInteractionRepo repository.UserInteractionRepository, tokenRepo repository.TokenRepository) UserInteractionService {
	return &userInteractionService{
		repo:  userInteractionRepo,
		tokenRepo: tokenRepo,
	}
}
