package service

import (
	"Somali-Newsletter-App/internal/entity"
	"Somali-Newsletter-App/internal/repository"
	"time"

	"github.com/gofrs/uuid"
)

type NewsletterService interface {
	CreateNewslatter(title, content string, creatorID uuid.UUID, published bool) (*entity.Newsletter, error)
	UpdateNewslatter(Newsletter *entity.Newsletter) error
	GetAllNewsletter() ([]*entity.Newsletter, error)
	GetNewsletterByID(NewsletterID uuid.UUID) (*entity.Newsletter, error)
	DeleteNewsletter(NewsletterID uuid.UUID) error
}

type newsletterService struct {
	repo      repository.NewsletterRepository
	tokenRepo repository.TokenRepository
}

// CreateNewslatter implements NewsletterService.
func (s *newsletterService) CreateNewslatter(title string, content string, creatorID uuid.UUID, published bool) (*entity.Newsletter, error) {
	// generate uuid 
	newUUID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	Newsletter := &entity.Newsletter{
		ID: newUUID,
		Title: title,
		Content: content,
		CreatorID: creatorID,
		Published: published,
		CreatedAt: time.Now(),
	
	}

	if err := s.repo.Create(Newsletter); err != nil {
		return nil, err
	}

	return Newsletter, err

}

// DeleteNewsletter implements NewsletterService.
func (s *newsletterService) DeleteNewsletter(NewsletterID uuid.UUID) error {
	_, err := s.repo.GetByID(NewsletterID)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(NewsletterID); err != nil {
		return err
	}
	return nil
}

// GetAllNewsletter implements NewsletterService.
func (s *newsletterService) GetAllNewsletter() ([]*entity.Newsletter, error) {
	newslatter, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return newslatter, nil
}

// GetNewsletterByID implements NewsletterService.
func (s *newsletterService) GetNewsletterByID(NewsletterID uuid.UUID) (*entity.Newsletter, error) {
	newslatter, err := s.repo.GetByID(NewsletterID)
	if err != nil {
		return nil, err
	}

	return newslatter, err
}

// UpdateNewslatter implements NewsletterService.
func (s *newsletterService) UpdateNewslatter(Newsletter *entity.Newsletter) error {
	_, err := s.repo.GetByID(Newsletter.ID)
	if err != nil {
	 return err
	}

	if err := s.repo.Update(Newsletter); err != nil {
		return err
	}

	return nil
}

func NewNewsletterService(newsletterRepo repository.NewsletterRepository, tokenRepo repository.TokenRepository) NewsletterService {
	return &newsletterService{
		repo:      newsletterRepo,
		tokenRepo: tokenRepo,
	}
}
