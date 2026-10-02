package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email     string    `json:"email" gorm:"not null;unique;"`
	Password  string    `json:"-" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`

	Tasks []Task `json:"-" gorm:"foreignKey:UserID;OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
