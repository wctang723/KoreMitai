package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    string    `json:"user_id" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Passwd    string    `json:"password" binding:"required"`
}
