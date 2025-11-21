package user

import (
	"monthly-expenses-handler/internal/infrastructure/crypto"
	"monthly-expenses-handler/internal/infrastructure/logger"
	userRepo "monthly-expenses-handler/internal/repository/user"
	"monthly-expenses-handler/internal/service/auth"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserController struct {
	Repo userRepo.Repository
}

func NewUserController(pool *pgxpool.Pool, cryptoSvc *crypto.CryptoService, authSvc *auth.AuthService, logger *logger.Logger) *UserController {
	return &UserController{
		Repo: userRepo.NewPostgresRepository(
			pool,
			cryptoSvc,
			authSvc,
			logger,
		),
	}
}
