package entity

import (
	"time"


	"github.com/gofrs/uuid"
)

type Newsletter struct {
    ID          uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"` // Unique ID for the newsletter
    Title       string     `json:"title" binding:"required" gorm:"size:255"`                 // Title of the newsletter
    Content     string     `json:"content" binding:"required"`                               // Main content of the newsletter
    CreatorID   uuid.UUID  `json:"creator_id" gorm:"type:uuid;not null"`                     // ID of the user (content creator)
    Published   bool       `json:"published" gorm:"default:false"`                           // Indicates if the newsletter is published
    DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"index"`                        // Timestamp for soft deletes
    CreatedAt   time.Time  `json:"created_at" gorm:"default:current_timestamp"`              // Timestamp when the newsletter was created
    UpdatedAt   time.Time  `json:"updated_at" gorm:"default:current_timestamp"`              // Timestamp when the newsletter was last updated
}
