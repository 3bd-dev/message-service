package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	RoleMaker   UserRole = "maker"
	RoleChecker UserRole = "checker"
)

type User struct {
	ID        uuid.UUID `json:"id"`         // Unique identifier for the user
	Username  string    `json:"username"`   // Username of the user
	Email     string    `json:"email"`      // Email of the user
	Role      UserRole  `json:"role"`       // Role of the user (maker or checker)
	CreatedAt time.Time `json:"created_at"` // Timestamp when the user was created
	UpdatedAt time.Time `json:"updated_at"` // Timestamp when the user was last updated
}

func (u *User) IsEmpty() bool {
	return u == nil || u.ID.String() == ""
}
