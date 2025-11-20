package config

import (
	"io"
	"monthly-expenses-handler/internal/logger"
)

type loggerConfig struct {
	Out      io.Writer
	MinLevel logger.Level
}
