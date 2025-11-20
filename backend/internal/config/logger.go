package config

import (
	"io"
	"monthly-expenses-handler/internal/infrastructure/logger"
)

type loggerConfig struct {
	Out      io.Writer
	MinLevel logger.Level
}
