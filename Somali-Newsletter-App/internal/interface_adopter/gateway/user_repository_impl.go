package gateway

import (
	"Somali-Newsletter-App/internal/entity"
	"Somali-Newsletter-App/internal/repository"
	"database/sql"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

// UserRepositoryImpl implements repository.UserRepository.
type UserRepositoryImpl struct {
	db *sql.DB
}

// Create implements repository.UserRepository.
func (r *UserRepositoryImpl) Create(user *entity.User) error {
	// generate a new UUID
	newUUID, err := uuid.NewV4()
	if err != nil {
		return err
	}

	user.ID = newUUID

	if user.RoleID == uuid.Nil {
		return fmt.Errorf("invalid role ID %v", user.RoleID)
	}

	// insert the user into the database
	query := `INSERT INTO users (id, name, email, password, role_id, created_at) VALUES ($1, $2, $3, $4, $5, $6)`
	result, err := r.db.Exec(query, user.ID, user.Name, user.Email, user.Password, user.RoleID, user.CreatedAt)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return err
	}
	if rowsAffected == 0 {
		log.Printf(" No rows affected")
		return nil
	}
	return nil
}

// Delete implements repository.UserRepository.
func (r *UserRepositoryImpl) Delete(userID uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.Exec(query, userID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		log.Printf("No rows affected")
		return nil
	}
	return nil
}

// FindByEmail implements repository.UserRepository.
func (r *UserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	query := `SELECT id, name, email, password, role_id, created_at, updated_at FROM users WHERE email = $1`
	row := r.db.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with email %s not found", email)

		}
		return nil, err
	}
	return user, nil
}

// GetAll implements repository.UserRepository.
func (r *UserRepositoryImpl) GetAll() ([]*entity.User, error) {
	query := `SELECT id, name, email, password, role_id, created_at, updated_at FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User

	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, &user)

	}
	if err := rows.Err(); err != nil {
		log.Printf("Error reading rows: %v", err)
		return nil, err
	}
	return users, nil
}

// GetByID implements repository.UserRepository.
func (r *UserRepositoryImpl) GetByID(userID uuid.UUID) (*entity.User, error) {
	var user entity.User

	query := `SELECT id, name, email, password, role_id, created_at, updated_at FROM users WHERE id = $1`
	err := r.db.QueryRow(query, userID).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with ID %s not found", userID)
		}
	}
	log.Printf("User found: %v", user)

	// return the user
	return &user, nil
}

// Update implements repository.UserRepository.
func (r *UserRepositoryImpl) Update(user *entity.User) error {
	query := `UPDATE users SET name = $1, email = $2, password = $3, role_id = $4, updated_at = $5 WHERE id = $6`
	result, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.RoleID, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return err
	}
	if rowsAffected == 0 {
		log.Printf("No rows affected")
		return nil
	}
	return nil
}

// NewUserRepositoryImpl creates a new UserRepositoryImpl.
func NewUserRepositoryImpl(db *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}

}
