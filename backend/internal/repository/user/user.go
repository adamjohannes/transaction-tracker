package user

import (
	"context"
	"errors"
	"fmt"
	"monthly-expenses-handler/internal/domain/user"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, user *user.User) (int64, error)
	GetByUsername(ctx context.Context, username string) (*user.User, error)
}

type postgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(ctx context.Context, u *user.User) (int64, error) {
	var id int64
	query := `INSERT INTO users (username, hashed_password) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(ctx, query, u.Username, u.HashedPassword).Scan(&id)
	if err != nil {
		// Basic check for unique constraint violation
		if strings.Contains(err.Error(), "users_username_unique") {
			return 0, fmt.Errorf("username '%s' is already taken", u.Username)
		}
		return 0, fmt.Errorf("failed to create user: %w", err)
	}
	return id, nil
}

func (r *postgresRepository) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	var u user.User
	query := `SELECT id, username, hashed_password FROM users WHERE username = $1`
	err := r.db.QueryRow(ctx, query, username).Scan(&u.ID, &u.Username, &u.HashedPassword)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &u, nil
}
