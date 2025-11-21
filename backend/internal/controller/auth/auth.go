package auth

import (
	"errors"
	"monthly-expenses-handler/internal/api_error"
	"monthly-expenses-handler/internal/domain/user"
	"monthly-expenses-handler/internal/infrastructure/logger"
	"monthly-expenses-handler/internal/service/auth"
	userSvc "monthly-expenses-handler/internal/usecase/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	authSvc *auth.AuthService
	userSvc *userSvc.UseCase
	logger  *logger.Logger
}

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAuthController(authSvc *auth.AuthService, userSvc *userSvc.UseCase, logger *logger.Logger) *Controller {
	return &Controller{
		authSvc,
		userSvc,
		logger,
	}
}

func (ac *Controller) Register(c *gin.Context) {
	ac.logger.Info("Received a request to register a new user", nil)

	request, err := collectRequest(c)
	if err != nil {
		ac.logger.Error("Failed to bind request", map[string]interface{}{"error": err})
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind request",
			"detail":  err,
		})
		return
	}

	newUser, err := buildUserObj(request)
	if err != nil {
		ac.logger.Error("Failed to build new user", map[string]interface{}{"error": err})
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to build new user",
			"detail":  err,
		})
		return
	}

	ac.logger.Info("Attempting to register a new user...", map[string]interface{}{"username": newUser.Username})

	token, err := ac.userSvc.Register(newUser)
	if err != nil {
		ac.logger.Error("Failed to register user", map[string]interface{}{"error": err})
		var validationErr *api_error.ValidationError
		if errors.As(err, &validationErr) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed register user",
				"detail":  validationErr,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to register user",
				"detail":  err,
			})
		}
		return
	}

	ac.logger.Info("Successfully registered user", map[string]interface{}{
		"token":    token,
		"username": newUser.Username,
	})

	c.JSON(http.StatusCreated, map[string]string{"token": token})
}

func (ac *Controller) Login(c *gin.Context) {
	ac.logger.Info("Received a request to log-in a user", nil)

	requestDatamap, err := collectRequest(c)
	if err != nil {
		ac.logger.Error("Failed to bind request", map[string]interface{}{"error": err})
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind request",
			"detail":  err,
		})
	}

	requestedUser, err := buildUserObj(requestDatamap)
	if err != nil {
		ac.logger.Error("Failed to build new user", map[string]interface{}{"error": err})
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to build new user",
			"detail":  err,
		})
		return
	}

	ac.logger.Info("Attempting to login user...", map[string]interface{}{"username": requestedUser.Username})

	token, err := ac.userSvc.Login(requestedUser)
	if err != nil {
		ac.logger.Error("Failed to log-in user", map[string]interface{}{"error": err})

		var authErr *api_error.AuthError
		var validationErr *api_error.ValidationError

		if errors.As(err, &authErr) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Failed to log-in user",
				"detail":  authErr,
			})
		} else if errors.As(err, &validationErr) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Failed to log-in user",
				"detail":  validationErr,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to log-in user",
				"detail":  err,
			})
		}

		return
	}

	ac.logger.Info("Successfully logged in user", map[string]interface{}{"username": requestedUser.Username})
	c.JSON(http.StatusOK, map[string]string{"token": token})
}

// --- Helpers

func collectRequest(c *gin.Context) (*authRequest, error) {
	var request *authRequest

	// Parse the request body
	if err := c.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func buildUserObj(request *authRequest) (*user.User, error) {
	newUser, err := user.New(request.Username, request.Password)
	if err != nil {
		return nil, api_error.NewValidationError("invalid credentials", err)
	}

	return newUser, nil
}
