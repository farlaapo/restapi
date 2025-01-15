package gateway

import (
	"Somali-Newsletter-App/internal/entity"
	"Somali-Newsletter-App/internal/repository"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

// Gateway is a struct that represents a gateway.
type UserInteractionRepositoryImpl struct {
	db *sql.DB
}

// Create implements repository.UserInteractionRepository.
func (r *UserInteractionRepositoryImpl) Create(userInteraction *entity.UserInteraction) error {
	// generate uuid
	newUUID, err := uuid.NewV4()
	if err != nil {
		return err
	}

	userInteraction.ID = newUUID
	// insert into database
	query := `INSERT INTO user_interactions (id, newslatter_id, user_id, liked, read, comment, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7);
`

	result, err := r.db.Exec(query, userInteraction.ID, userInteraction.NewsletterID, userInteraction.UserID, userInteraction.Liked, userInteraction.Read, userInteraction.Comment, userInteraction.CreatedAt)
	if err != nil {
		log.Printf(" Error inserting user interaction: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf(" Error getting rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		log.Printf(" No rows affected")
		return nil
	}

	log.Printf(" User interaction created with id: %v", userInteraction.ID)
	return nil
}

// Delete implements repository.UserInteractionRepository.
func (r *UserInteractionRepositoryImpl) Delete(userInteractionID uuid.UUID) error {
	// Try to delete the record
	res, err := r.db.Exec(`
			DELETE FROM user_interactions
			WHERE id = $1
	`, userInteractionID)

	if err != nil {
		return fmt.Errorf("error deleting user interaction: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting number of rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user interaction found with id %v", userInteractionID)
	}

	return nil
}

// Get implements repository.UserInteractionRepository.
func (r *UserInteractionRepositoryImpl) Get(userInteractionID uuid.UUID) (*entity.UserInteraction, error) {
	var userInteraction entity.UserInteraction

	// sql statment
	query := `SELECT id, newslatter_id, user_id, liked, read, comment, created_at, updated_at
  FROM user_interactions
  WHERE id = $1`
	row := r.db.QueryRow(query, userInteractionID).Scan(
		&userInteraction.ID,
		&userInteraction.NewsletterID,
		&userInteraction.UserID,
		&userInteraction.Liked,
		&userInteraction.Read,
		&userInteraction.Comment,
		&userInteraction.CreatedAt,
		&userInteraction.UpdatedAt,
	)
	if row != nil {
		log.Printf("Error scanning userInteraction: %v", row)
		return nil, row
	}
	return &userInteraction, nil

}

// GetAll implements repository.UserInteractionRepository.
func (r *UserInteractionRepositoryImpl) GetAll() ([]*entity.UserInteraction, error) {
	query := `SELECT id, newslatter_id, user_id, liked, read, comment, created_at, updated_at FROM user_interactions`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var userInteractions []*entity.UserInteraction
	for rows.Next() {
		var userInteraction entity.UserInteraction
		if err := rows.Scan(&userInteraction.ID, &userInteraction.NewsletterID, &userInteraction.UserID, &userInteraction.Liked, &userInteraction.Read, &userInteraction.Comment, &userInteraction.CreatedAt, &userInteraction.UpdatedAt); err != nil {
			return nil, err
		}
		userInteractions = append(userInteractions, &userInteraction)
	}
	if err := rows.Err(); err != nil {
		log.Printf(" Error getting user interactions: %v", err)
		return nil, err
	}
	return userInteractions, nil
}

// Update implements repository.UserInteractionRepository.
func (r *UserInteractionRepositoryImpl) Update(userInteraction *entity.UserInteraction) error {

	result, err := r.db.Exec(`
    UPDATE user_interactions
    SET newslatter_id = $1, user_id = $2, liked = $3, read = $4, comment = $5, updated_at = $6
    WHERE id = $7
`, userInteraction.NewsletterID, userInteraction.UserID, userInteraction.Liked, userInteraction.Read, userInteraction.Comment, time.Now(), userInteraction.ID)
	if err != nil {
		return fmt.Errorf("error updating user interaction: %v", err)
	}

	rowaAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowaAffected == 0 {
		return nil
	}

	return nil
}

// NewUserInteractionRepositoryImpl return func

func NewUserInteractionRepositoryImpl(db *sql.DB) repository.UserInteractionRepository {
	return &UserInteractionRepositoryImpl{db: db}
}
