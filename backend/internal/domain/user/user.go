package user

import (
	"fmt"
	"regexp"
)

type User struct {
	ID             int64
	Username       string
	HashedPassword string
}

func New(username, password string) (*User, error) {
	if username == "" {
		return nil, fmt.Errorf("username is required")
	}
	if len(password) < 8 {
		return nil, fmt.Errorf("password must be at least 8 characters long")
	}

	// Regex to ensure username contains only letters, numbers, and underscores
	re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !re.MatchString(username) {
		return nil, fmt.Errorf("username can only contain letters, numbers, and underscores")
	}

	return &User{
		ID:       -1,
		Username: username,
	}, nil
}

func Build(id int64, username string, hashedPassword string) *User {
	return &User{
		ID:             id,
		Username:       username,
		HashedPassword: hashedPassword,
	}
}
