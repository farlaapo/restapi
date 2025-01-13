package gateway

import (
	"Somali-Newsletter-App/internal/entity"
	"Somali-Newsletter-App/internal/repository"
	"database/sql"
	"fmt"
	"log"
	

	"github.com/gofrs/uuid"
)

// NewsletterRepositoryImpl  struct
type NewsletterRepositoryImpl struct {
	db *sql.DB
}

// Create implements repository.NewsletterRepository.
func (r *NewsletterRepositoryImpl) Create(newsletter *entity.Newsletter) error {
	// generate uuid
	newUUD, err := uuid.NewV4()
	if err != nil {
		return err
	}

	newsletter.ID = newUUD
	query := `INSERT INTO news_latters(id, title, content, creator_id, puplished, created_at)
 VALUES ($1, $2, $3, $4, $5, $6, )`
	result, err := r.db.Exec(query, newsletter.ID, newsletter.Title, newsletter.Content, newsletter.CreatorID, newsletter.Published, newsletter.CreatedAt)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return nil
	}

	return nil

}

// Delete implements repository.NewsletterRepository.
func (r *NewsletterRepositoryImpl) Delete(NewsletterID uuid.UUID) error {
	query := `DELETE FROM news_latters WHERE id = $1`
	result, err := r.db.Exec(query, NewsletterID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return nil
	}

	return nil
}

// GetAll implements repository.NewsletterRepository.
func (r *NewsletterRepositoryImpl) GetAll() ([]*entity.Newsletter, error) {

	row, err := r.db.Query("SELECT id, title, content, creator_id, puplished, created_at, updated_at FROM news_latters")
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var newsLatters []*entity.Newsletter
	for row.Next() {
		var newsLatter entity.Newsletter
		err := row.Scan(&newsLatter.ID, &newsLatter.Title, &newsLatter.Content, &newsLatter.CreatorID, &newsLatter.Published, &newsLatter.CreatedAt, &newsLatter.UpdatedAt)
		if err != nil {
			return nil, err
		}
		newsLatters = append(newsLatters, &newsLatter)
	}

	return newsLatters, nil

}

// GetByID implements repository.NewsletterRepository.
func (r *NewsletterRepositoryImpl) GetByID(newsletterID uuid.UUID) (*entity.Newsletter, error) {
	var newsLatter entity.Newsletter

 err := r.db.QueryRow(`SELECT id, title, content, creator_id, puplished, created_at, updated_at FROM news_latters WHERE id = $1`, newsletterID ).Scan(
	&newsLatter.ID, &newsLatter.Title, &newsLatter.Content, &newsLatter.CreatorID, &newsLatter.Published,  &newsLatter.CreatedAt, &newsLatter.UpdatedAt)
if err != nil {
	if err == sql.ErrNoRows {
		log.Printf(" No newslatter not found with by ID: %v", newsletterID)
		return nil, fmt.Errorf("  newslatter not found ")
	}
	log.Printf(" Error retrieving newslatter with by ID: %v", err)
	return nil, err
}

return &newsLatter, nil

}

// Update implements repository.NewsletterRepository.
func (r *NewsletterRepositoryImpl) Update(newsLatter *entity.Newsletter) error {
	// sql
	query := ` UPDATE news_latters
	SET title = $1, content = $2, creator_id = $3, puplished = $4, updated_at = $5
	WHERE id = $6`

	result , err := r.db.Exec(query, newsLatter.Title, newsLatter.Content, newsLatter.CreatorID, newsLatter.Published, newsLatter.UpdatedAt )
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return nil
	}


	return nil

}

// NewNewsletterRepositoryImpl  function

func NewNewsletterRepositoryImpl(db *sql.DB) repository.NewsletterRepository {
	return &NewsletterRepositoryImpl{db: db}

}
