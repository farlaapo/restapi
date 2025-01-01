package gateway

import (
	"Somali-Newsletter-App/internal/entity"
	"Somali-Newsletter-App/internal/repository"
	"database/sql"
)

// token is a token that can be used to authenticate with the gateway.
type TokenRepositoryImpl struct {
	db *sql.DB
}

// Create implements repository.TokenRepository.
func (t *TokenRepositoryImpl) Create(token *entity.Token) error {
	panic("unimplemented")
}

// FindToken implements repository.TokenRepository.
func (t *TokenRepositoryImpl) FindToken(token string) (*entity.Token, error) {
	panic("unimplemented")
}

// 	
func NewTokenRepositoryImpl(db *sql.DB) repository.TokenRepository {
	return &TokenRepositoryImpl{db: db}
}
