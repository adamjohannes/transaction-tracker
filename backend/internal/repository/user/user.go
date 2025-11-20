package user

import (
	"context"
	"errors"
	"fmt"
	"monthly-expenses-handler/internal/api_error"
	"monthly-expenses-handler/internal/domain/user"
	"monthly-expenses-handler/internal/infrastructure/crypto"
	"monthly-expenses-handler/internal/infrastructure/logger"
	"monthly-expenses-handler/internal/service/auth"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, u *user.User) (int64, error)
	GetByUsername(ctx context.Context, username string) (*user.User, error)
}

type postgresRepository struct {
	authSvc *auth.AuthService
	crypto  *crypto.CryptoService
	db      *pgxpool.Pool
	logger  *logger.Logger
}

func NewPostgresRepository(db *pgxpool.Pool, crypto *crypto.CryptoService, authSvc *auth.AuthService, logger *logger.Logger) Repository {
	return &postgresRepository{
		authSvc,
		crypto,
		db,
		logger,
	}
}

func (r *postgresRepository) Create(ctx context.Context, u *user.User) (int64, error) {
	var id int64

	r.logger.Debug("Encrypting username...", map[string]interface{}{"username": u.Username})

	encryptedUsername, err := r.crypto.Encrypt([]byte(u.Username))
	if err != nil {
		return 0, api_error.NewAnyError("failed to encrypt username", err)
	}

	r.logger.Debug("Creating search hash...", map[string]interface{}{"username": u.Username})

	usernameHash := r.authSvc.CreateSearchHash(u.Username)

	query := `INSERT INTO users (username, username_search_hash, hashed_password) VALUES ($1, $2, $3) RETURNING id`

	r.logger.Debug("Querying database...", map[string]interface{}{"query": query})

	err = r.db.QueryRow(ctx, query, encryptedUsername, usernameHash, u.HashedPassword).Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "users_username_hash_unique") {
			return 0, api_error.NewValidationError("username already taken", err)
		}
		return 0, api_error.NewAnyError("failed to create user: %w", err)
	}

	return id, nil
}

func (r *postgresRepository) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	var u user.User
	var encryptedUsername []byte

	r.logger.Debug("Creating search hash...", map[string]interface{}{"username": username})

	usernameHash := r.authSvc.CreateSearchHash(username)

	query := `SELECT id, username, hashed_password FROM users WHERE username_search_hash = $1`

	r.logger.Debug("Querying database...", map[string]interface{}{"query": query})

	err := r.db.QueryRow(ctx, query, usernameHash).Scan(&u.ID, &encryptedUsername, &u.HashedPassword)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user by hash: %w", err)
	}

	r.logger.Debug("Decrypting username...", nil)

	decryptedUsername, err := r.crypto.Decrypt(encryptedUsername)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt username: %w", err)
	}

	u.Username = string(decryptedUsername)
	return &u, nil
}
