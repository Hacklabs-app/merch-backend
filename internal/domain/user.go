package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	FullName     string    `json:"full_name"`
	PhoneNumber  *string   `json:"phone_number,omitempty"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

type UserRegistration struct {
	Email       string  `json:"email" validate:"required,email"`
	Password    string  `json:"password" validate:"required,min=8"`
	FullName    string  `json:"full_name" validate:"required"`
	PhoneNumber *string `json:"phone_number,omitempty"`
}

type UserRepository interface {
	Create(ctx context.Context, email, passwordHash, fullName string, phoneNumber *string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
}

type AuthService interface {
	Register(ctx context.Context, email, password, fullName string, phoneNumber *string) (*User, error)
	Login(ctx context.Context, email, password string) (string, error)
}
