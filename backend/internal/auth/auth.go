package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	jwtSecret []byte
	hashKey   []byte
}

func NewAuthService(secret, hashKey string) *AuthService {
	return &AuthService{
		jwtSecret: []byte(secret),
		hashKey:   []byte(hashKey),
	}
}

// CreateSearchHash
// Creates an HMAC-SHA256 hash for blind indexing.
func (s *AuthService) CreateSearchHash(data string) []byte {
	h := hmac.New(sha256.New, s.hashKey)
	h.Write([]byte(data))
	return h.Sum(nil)
}

// HashPassword
// Generates a bcrypt hash of the password.
func (s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash
// Compares a password with a hash.
func (s *AuthService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWT
// Creates a new JWT for a given userID.
func (s *AuthService) GenerateJWT(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}

// ValidateJWT
// Validates a token string and returns the userID (sub claim).
func (s *AuthService) ValidateJWT(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.jwtSecret, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if sub, ok := claims["sub"].(float64); ok {
			return int64(sub), nil
		}
	}

	return 0, fmt.Errorf("invalid token")
}
