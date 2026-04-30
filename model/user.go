package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" binding:"omitempty"`
	CreatedAt time.Time `json:"created_at" binding:"omitempty"`
	UpdatedAt time.Time `json:"updated_at" binding:"omitempty"`
	UserID    string    `json:"user_id" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Passwd    string    `json:"password" binding:"required"`
}
