package config

import "errors"

type PostgresConfig struct {
	DatabaseURL  string
	DatabaseName string
	DatabaseUser string
	DatabasePass string
}

func New(url, name, user, password string) (*PostgresConfig, error) {
	if url == "" {
		return nil, errors.New("postgres url required")
	}

	if name == "" {
		return nil, errors.New("postgres database name required")
	}

	if user == "" {
		return nil, errors.New("postgres username required")
	}

	if password == "" {
		return nil, errors.New("postgres password required")
	}

	return &PostgresConfig{
		DatabaseURL:  url,
		DatabaseName: name,
		DatabaseUser: user,
		DatabasePass: password,
	}, nil
}
