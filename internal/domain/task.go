package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid();"`
	Title       string     `json:"title" gorm:"not null"`
	Description *string    `json:"description"`
	Status      string     `json:"status" gorm:"default:'new'"`
	Deadline    *time.Time `json:"deadline"`
	CreatedAt   time.Time  `json:"createdAt"`

	UserID uuid.UUID `json:"userId" gorm:"type:uuid;not null"`
}
