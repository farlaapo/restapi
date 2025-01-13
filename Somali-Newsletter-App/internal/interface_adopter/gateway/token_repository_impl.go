package gateway

import (
	"Somali-Newsletter-App/internal/entity"
	"Somali-Newsletter-App/internal/repository"
	"database/sql"
	"errors"
	"log"
	"time"
)

// token is a token that can be used to authenticate with the gateway.
type TokenRepositoryImpl struct {
	db *sql.DB
}

// Create implements repository.TokenRepository.
func (r *TokenRepositoryImpl) Create(token *entity.Token) error {
	 query := `INSERT INTO tokens (id, user_id, token, expires_at, created_at, updated_at)
	 VALUES ($1, $2, $3, $4, $5, $6)`
	result, err := r.db.Exec(query, token.ID, token.UserID, token.Token, token.ExpiredAt, time.Now(), time.Now())
	if err != nil {
		log.Printf(" Error creating token: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf(" Error getting rows affected: %v", err)
		return err
	}

	log.Printf(" Rows affected: %d", rowsAffected)
	return nil
}

// FindToken implements repository.TokenRepository.
func (r *TokenRepositoryImpl) FindToken(token string) (*entity.Token, error) {
	t := &entity.Token{}
	query := `SELECT id, user_id, token, expires_at, created_at, updated_at FROM tokens WHERE token = $1`
	row := r.db.QueryRow(query, token)

	err := row.Scan(
		&t.ID, 
		&t.UserID,
		&t.Token,
		&t.ExpiredAt,
		&t.CreatedAt,
		&t.UpdatedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("token not found")
			}
			return nil, err
		}

		return t, nil
}

// 	
func NewTokenRepository(db *sql.DB) repository.TokenRepository {
	return &TokenRepositoryImpl{db: db}
}
