package service

import (
	"Somali-Newsletter-App/internal/entity"
	"Somali-Newsletter-App/internal/repository"
	"Somali-Newsletter-App/pkg/utils"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

// UserService is a generic interface for a data access layer.
type UserService interface {
	RegesterUser(Name, Email, Password string, RoleID uuid.UUID) (*entity.User, error)
	AuthenticateUser(email, Password string) (*entity.User, error)
	UpdateUser(user *entity.User) error
	GetAllUser() ([]*entity.User, error)
	GetUserByID(userID uuid.UUID) (*entity.User, error)
	DeleteUser(userID uuid.UUID) error
}

// userService implements the UserService interface
type userService struct {
	repo      repository.UserRepository
	tokenRepo repository.TokenRepository
}

// DeleteUser implements UserService.
func (s *userService) DeleteUser(userID uuid.UUID) error {
	_, err := s.repo.GetByID(userID)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(userID); err != nil {
		log.Printf(" Error deleting user: %v", err)
		return err

	}
	return nil
}

// GetUserAll implements UserService.
func (s *userService) GetAllUser() ([]*entity.User, error) {
	user, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByID implements UserService.
func (s *userService) GetUserByID(userID uuid.UUID) (*entity.User, error) {
	user, err := s.repo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// RegesterUser implements UserService.
func (s *userService) RegesterUser(Name string, Email string, Password string, RoleID uuid.UUID) (*entity.User, error) {
	// check if user already exists
	if _, err := s.repo.FindByEmail(Email); err == nil {
		return nil, errors.New(" user with this email already exists")
	}

	// hash password
	hashPassword, err := utils.HashPassword(Password)
	if err != nil {
		return nil, err
	}

	// create new user
	user := &entity.User{
		Name:     Name,
		Email:    Email,
		Password: hashPassword,
		RoleID:   RoleID,
	}

	// save user
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser implements UserService.
func (s *userService) UpdateUser(user *entity.User) error {
	_, err := s.repo.GetByID(user.ID)
	if err != nil {
		return err
	}

	if err := s.repo.Update(user); err != nil {
		return fmt.Errorf(" failed to update user: %v", err)
	}
	return nil
}

// authenticateUser implements UserService.
func (s *userService) AuthenticateUser(email string, Password string) (*entity.User, error) {
	// find user email
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New(" user not found")
	}

	// generate err
	newToken, err := uuid.NewV4()
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	// create token
	token := &entity.Token{
		ID:        newToken,
		UserID:    user.ID,
		Token:     newToken.String(),
		ExpiredAt: time.Now().Add(24 * time.Hour),
	}

	// store token in database
	if err := s.tokenRepo.Create(token); err != nil {
		return nil, errors.New(" failed to create token")
	}
	return user, nil
}

// NewUserService creates a new UserService
func NewUserService(userRepo repository.UserRepository, tokenRepo repository.TokenRepository) UserService {
	return &userService{
		repo:      userRepo,
		tokenRepo: tokenRepo,
	}
}
