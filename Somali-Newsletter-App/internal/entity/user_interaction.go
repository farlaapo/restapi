package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type UserInteraction struct {
    ID           uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"` // Unique ID for the interaction
    NewsletterID uuid.UUID  `json:"newsletter_id" gorm:"type:uuid;not null"`                  // ID of the newsletter
    UserID       uuid.UUID  `json:"user_id" gorm:"type:uuid;not null"`                        // ID of the user interacting with the newsletter
    Liked        bool       `json:"liked" gorm:"default:false"`                               // Whether the user liked the newsletter
    Read         bool       `json:"read" gorm:"default:false"`                                // Whether the user read the newsletter
    Comment      string     `json:"comment" gorm:"type:text"`                                 // Comment made by the user (optional)
    CreatedAt    time.Time  `json:"created_at" gorm:"default:current_timestamp"`              // Timestamp for when the interaction occurred
    UpdatedAt    time.Time  `json:"updated_at" gorm:"default:current_timestamp"`              // Timestamp for updates to the interaction
}
