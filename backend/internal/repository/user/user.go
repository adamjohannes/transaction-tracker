package user

import (
	"context"
	"errors"
	"fmt"
	"monthly-expenses-handler/internal/auth"
	"monthly-expenses-handler/internal/crypto"
	"monthly-expenses-handler/internal/domain/user"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, u *user.User) (int64, error)
	GetByUsername(ctx context.Context, username string) (*user.User, error)
}

type postgresRepository struct {
	db      *pgxpool.Pool
	crypto  *crypto.CryptoService
	authSvc *auth.AuthService
}

func NewPostgresRepository(db *pgxpool.Pool, crypto *crypto.CryptoService, authSvc *auth.AuthService) Repository {
	return &postgresRepository{db: db, crypto: crypto, authSvc: authSvc}
}

func (r *postgresRepository) Create(ctx context.Context, u *user.User) (int64, error) {
	var id int64

	encryptedUsername, err := r.crypto.Encrypt([]byte(u.Username))
	if err != nil {
		return 0, fmt.Errorf("failed to encrypt username: %w", err)
	}

	usernameHash := r.authSvc.CreateSearchHash(u.Username)

	query := `INSERT INTO users (username, username_search_hash, hashed_password) VALUES ($1, $2, $3) RETURNING id`
	err = r.db.QueryRow(ctx, query, encryptedUsername, usernameHash, u.HashedPassword).Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "users_username_hash_unique") {
			return 0, fmt.Errorf("username '%s' is already taken", u.Username)
		}
		return 0, fmt.Errorf("failed to create user: %w", err)
	}
	return id, nil
}

func (r *postgresRepository) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	var u user.User
	var encryptedUsername []byte
	usernameHash := r.authSvc.CreateSearchHash(username)

	query := `SELECT id, username, hashed_password FROM users WHERE username_search_hash = $1`
	err := r.db.QueryRow(ctx, query, usernameHash).Scan(&u.ID, &encryptedUsername, &u.HashedPassword)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user by hash: %w", err)
	}

	decryptedUsername, err := r.crypto.Decrypt(encryptedUsername)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt username: %w", err)
	}
	u.Username = string(decryptedUsername)

	return &u, nil
}
