package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name" `
	Email      string     `json:"email" binding:"required"`
	Password   string     `json:"password" binding:"required"`
	RoleID     uuid.UUID  `json:"role_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	Deleted_at *time.Time `json:"deleted_at,omitempty"`
}
