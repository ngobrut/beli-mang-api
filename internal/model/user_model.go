package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/ngobrut/beli-mang-api/constant"
)

type User struct {
	UserID    uuid.UUID     `json:"user_id" db:"user_id"`
	Username  string        `json:"username" db:"username"`
	Email     string        `json:"email" db:"email"`
	Password  string        `json:"-" db:"password"`
	Role      constant.Role `json:"role" db:"role"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time    `json:"-" db:"deleted_at"`
}
