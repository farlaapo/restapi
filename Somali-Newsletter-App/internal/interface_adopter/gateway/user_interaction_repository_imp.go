package gateway

import (
	"Somali-Newsletter-App/internal/entity"
	"Somali-Newsletter-App/internal/repository"
	"database/sql"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

// Gateway is a struct that represents a gateway.
type UserInteractionRepositoryImpl struct {
	db *sql.DB
}

// Create implements repository.UserInteractionRepository.
func (r *UserInteractionRepositoryImpl) Create(userInteraction *entity.UserInteraction) error {
 // generate uuid
 newUUID , err := uuid.NewV4()
 if err != nil {
	return err
 }

 userInteraction.ID = newUUID
 // insert into database
  query := `INSERT INTO user_interactions (id, news_latter_id, user_id, liked, read, comment, create_at )
	VALUES($1, $2, $3, $4, $5, $6, $7)`

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
	// database query 
	query := "DELETE FROM userInteractions WHERE id = $1"
	result, err := r.db.Exec(query, userInteractionID)
	if err != nil {
		log.Printf(" Error deleting user interaction: %v", err)
		return err
	}

	rowsAffected, err := result.LastInsertId()
	if err != nil {
		log.Printf(" Error getting rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		log.Printf(" No rows affected")
		return nil
	}

	log.Printf(" User interaction deleted with id: %v", userInteractionID)
	return nil

}

// Get implements repository.UserInteractionRepository.
func (r *UserInteractionRepositoryImpl) Get(userInteractionID uuid.UUID) (*entity.UserInteraction, error) {
	var userInteraction  entity.UserInteraction

	err := r.db.QueryRow(`SELECT id, news_latter_id, user_id, liked, read, comment, create_at, updated_at FROM userInteractions WHERE id = $1`, userInteractionID).Scan(
		&userInteraction.ID,
		&userInteraction.NewsletterID,
		&userInteraction.UserID,
		&userInteraction.Liked,
		&userInteraction.Read,
		&userInteraction.Comment,
		&userInteraction.CreatedAt,
		&userInteraction.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(" User interaction not found: %v", userInteractionID)
		}
		log.Printf(" Error getting user interaction: %v", err)
	}

	return &userInteraction, nil


}

// GetAll implements repository.UserInteractionRepository.
func (r *UserInteractionRepositoryImpl) GetAll() ([]*entity.UserInteraction, error) {
	query := `SELECT id, news_latter_id, user_id, liked, read, comment, create_at, updated_at FROM userInteractions`
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
		return nil , err
	}
 return userInteractions, nil
}

// Update implements repository.UserInteractionRepository.
func (r *UserInteractionRepositoryImpl) Update(userInteraction *entity.UserInteraction) error {
	// database query 
	query := `UPDATE userInteractions
	SET news_latter_id = $1, user_id = $2, liked, read = $3, comment = $4,  updated_at = $5
	WHERE id = $6`
	result, err := r.db.Exec(query, userInteraction.ID, userInteraction.NewsletterID, userInteraction.UserID, userInteraction.Liked, userInteraction.Read, userInteraction.Comment, userInteraction.UpdatedAt)
	if err != nil {
		return err
	}

	rowaAffected, err := result.LastInsertId()
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
