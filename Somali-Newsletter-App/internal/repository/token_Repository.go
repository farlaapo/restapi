package repository

import "Somali-Newsletter-App/internal/entity"


type TokenRepository interface{
	FindToken(token string) (*entity.Token, error)
	Create(token *entity.Token) error
} 