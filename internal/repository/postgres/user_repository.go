package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Hacklabs-app/merch-backend/internal/domain"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, email, passwordHash, fullName string, phoneNumber *string) (*domain.User, error) {
	query := `
		INSERT INTO users (email, password_hash, full_name, phone_number, role)
		VALUES ($1, $2, $3, $4, 'customer')
		RETURNING id, email, full_name, phone_number, role, created_at
	`

	var user domain.User
	err := r.db.QueryRowContext(ctx, query, email, passwordHash, fullName, phoneNumber).Scan(
		&user.ID, &user.Email, &user.FullName, &user.PhoneNumber, &user.Role, &user.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
		SELECT id, email, password_hash, full_name, phone_number, role, created_at
		FROM users
		WHERE email = $1
	`

	var user domain.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.FullName, &user.PhoneNumber, &user.Role, &user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return &user, nil
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	query := `
		SELECT id, email, full_name, phone_number, role, created_at
		FROM users
		WHERE id = $1
	`

	var user domain.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Email, &user.FullName, &user.PhoneNumber, &user.Role, &user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return &user, nil
}
