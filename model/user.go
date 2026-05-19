package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    string
	Email     string
	Passwd    string
}

// Used for create a user
type CreateUserForm struct {
	UserID string `json:"user_id" binding:"required"`
	Email  string `json:"email" binding:"required,email"`
	Passwd string `json:"password" binding:"required"`
}

type LoginUserForm struct {
	Email  string `json:"email" binding:"required,email"`
	Passwd string `json:"password" binding:"required"`
}
