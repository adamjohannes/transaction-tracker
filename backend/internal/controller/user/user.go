package user

import (
	"monthly-expenses-handler/internal/infrastructure/crypto"
	userRepo "monthly-expenses-handler/internal/repository/user"
	"monthly-expenses-handler/internal/service/auth"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserController struct {
	Repo userRepo.Repository
}

func NewUserController(pool *pgxpool.Pool, cryptoSvc *crypto.CryptoService, authSvc *auth.AuthService) *UserController {
	return &UserController{
		Repo: userRepo.NewPostgresRepository(pool, cryptoSvc, authSvc),
	}
}
