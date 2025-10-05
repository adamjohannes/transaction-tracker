package auth

import (
	"context"
	"fmt"
	"monthly-expenses-handler/internal/apierror"
	"monthly-expenses-handler/internal/auth"
	userCtrl "monthly-expenses-handler/internal/controller/user"
	"monthly-expenses-handler/internal/domain/user"
	userRepo "monthly-expenses-handler/internal/repository/user"
	"strings"
)

type AuthController struct {
	ctx      context.Context
	authSvc  *auth.AuthService
	userRepo userRepo.Repository
}

func NewAuthController(ctx context.Context, authSvc *auth.AuthService, userController *userCtrl.UserController) *AuthController {
	return &AuthController{
		ctx:      ctx,
		authSvc:  authSvc,
		userRepo: userController.Repo,
	}
}

func (c *AuthController) Register(payload map[string]any) (string, error) {
	username, okU := payload["username"].(string)
	password, okP := payload["password"].(string)

	if !okU || !okP {
		return "", apierror.NewValidationError("username and password are required")
	}

	newUser, err := user.New(username, password)
	if err != nil {
		return "", apierror.NewValidationError(err.Error())
	}

	hashedPassword, err := c.authSvc.HashPassword(password)
	if err != nil {
		return "", fmt.Errorf("could not hash password: %w", err)
	}
	newUser.HashedPassword = hashedPassword

	userID, err := c.userRepo.Create(c.ctx, newUser)
	if err != nil {
		// Check if it's a validation-style error
		if strings.Contains(err.Error(), "already taken") {
			return "", apierror.NewValidationError(err.Error())
		}
		return "", fmt.Errorf("could not create user in db: %w", err)
	}

	token, err := c.authSvc.GenerateJWT(userID)
	if err != nil {
		return "", fmt.Errorf("could not generate token: %w", err)
	}

	return token, nil
}

func (c *AuthController) Login(payload map[string]any) (string, error) {
	username, okU := payload["username"].(string)
	password, okP := payload["password"].(string)

	if !okU || !okP {
		return "", apierror.NewValidationError("username and password are required")
	}

	existingUser, err := c.userRepo.GetByUsername(c.ctx, username)
	if err != nil {
		return "", apierror.NewValidationError("invalid credentials")
	}

	if !c.authSvc.CheckPasswordHash(password, existingUser.HashedPassword) {
		return "", apierror.NewValidationError("invalid credentials")
	}

	token, err := c.authSvc.GenerateJWT(existingUser.ID)
	if err != nil {
		return "", fmt.Errorf("could not generate token: %w", err)
	}

	return token, nil
}
