package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name       string    `json:"name" binding:"required" gorm:"size:255"`
	Email      string    `json:"email" binding:"required" gorm:"size:255;unique"`
	Password   string    `json:"password" binding:"required" gorm:"size:255"`
	RoleID     uuid.UUID `json:"role_id" gorm:"type:uuid;not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:current_timestamp"`
	Deleted_at *time.Time
}
