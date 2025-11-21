package user

import (
	"context"
	"errors"
	"monthly-expenses-handler/internal/api_error"
	"monthly-expenses-handler/internal/domain/user"
	"monthly-expenses-handler/internal/infrastructure/logger"
	userRepo "monthly-expenses-handler/internal/repository/user"
	"monthly-expenses-handler/internal/service/auth"
)

type UseCase struct {
	authSvc  *auth.AuthService
	ctx      context.Context
	logger   *logger.Logger
	userRepo userRepo.Repository
}

func NewUserService(authSvc *auth.AuthService, ctx context.Context, logger *logger.Logger, userRepo userRepo.Repository) *UseCase {
	return &UseCase{
		authSvc,
		ctx,
		logger,
		userRepo,
	}
}

// Register
// Handles user registration.
func (u *UseCase) Register(newUser *user.User) (string, error) {
	u.logger.Debug("Hashing new user password...", map[string]interface{}{"username": newUser.Username})

	hashedPassword, err := u.authSvc.HashPassword(newUser.HashedPassword)
	if err != nil {
		return "", api_error.NewAnyError("failed to hash password", err)
	}
	newUser.HashedPassword = hashedPassword

	u.logger.Debug("Adding user to the database...", map[string]interface{}{"username": newUser.Username})

	userID, err := u.userRepo.Create(u.ctx, newUser)
	if err != nil {
		var validationError *api_error.ValidationError
		if errors.As(err, &validationError) {
			return "", api_error.NewValidationError("failed to add new user", err)
		}
		return "", api_error.NewAnyError("failed to add new user", err)
	}

	u.logger.Debug("Generating token...", map[string]interface{}{"username": newUser.Username})

	token, err := u.authSvc.GenerateJWT(userID)
	if err != nil {
		return "", api_error.NewAnyError("failed to generate token", err)
	}

	return token, nil
}

// Login
// Handles user login.
func (u *UseCase) Login(user *user.User) (string, error) {
	u.logger.Debug("Checking if user exists...", map[string]interface{}{"username": user.Username})

	existingUser, err := u.userRepo.GetByUsername(u.ctx, user.Username)
	if err != nil {
		return "", api_error.NewAuthError("user not found")
	}

	u.logger.Debug("Checking user password...", map[string]interface{}{"username": user.Username})

	if !u.authSvc.CheckPasswordHash(user.HashedPassword, existingUser.HashedPassword) {
		return "", api_error.NewAuthError("password incorrect")
	}

	u.logger.Debug("Generating token...", map[string]interface{}{"username": user.Username})

	token, err := u.authSvc.GenerateJWT(existingUser.ID)
	if err != nil {
		return "", api_error.NewAnyError("failed to generate token", err)
	}

	return token, nil
}
