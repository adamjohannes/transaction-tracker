package config

// postgresConfig
// Holds all configuration variables for the database connection.
type postgresConfig struct {
	Host          string
	Port          string
	User          string
	Password      string
	Name          string
	EncryptionKey string
	JWTSecret     string
	SearchHashKey string
}
