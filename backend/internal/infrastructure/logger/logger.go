package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"monthly-expenses-handler/internal/config"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Logger
// It is a structured, leveled logger.
type Logger struct {
	out      io.Writer
	minLevel Level
	mu       sync.Mutex
}

// New
// Creates a new Logger.
// `out` is the destination for the log output.
// `minLevel` is the minimum level to log.
func New(config *config.Config) *Logger {
	return &Logger{
		out:      config.Logger.Out,
		minLevel: config.Logger.MinLevel,
	}
}

// internalLog
// It is the private method that handles all log writing.
// Thread safe 😎 (at least until a thread explodes and proves otherwise).
func (l *Logger) internalLog(level Level, message string, fields map[string]interface{}) {
	// 1. Check the level
	if level < l.minLevel {
		return
	}

	// 2. Get the file and line number
	// `runtime.Caller()` returns the execution call stack.
	// If we go up 2 levels, we can collect the function
	// that called the logger:
	// Level 0: runtime.Caller()
	// Level 1: internalLog()
	// Level 2: the function that called the logger
	// - IT SEEMS the compiler "suppresses" the call to Info/Debugf/...
	//   through inlining
	// - I expected to have to go up 3 levels, because I imagined that at
	//   level 2 we would have the call to the public log methods,
	//   but that doesn't seem to be the case
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	} else {
		// 2.1 Clean the file path
		// The file will have the absolute path,
		// "/home/$USER/.../file.go", but since what interests
		// us is only "file.go", we can clean
		// the rest.
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}

	// 3. Define the structure of the log entry
	entry := struct {
		Time    string                 `json:"time"`             // Timestamp of when the log was recorded
		Level   string                 `json:"level"`            // Log level
		Message string                 `json:"message"`          // Log message
		File    string                 `json:"file,omitempty"`   // File that called the logger
		Fields  map[string]interface{} `json:"fields,omitempty"` // Structured information
	}{
		Time:    time.Now().UTC().Format(time.RFC3339Nano), // RFC3339Nano -> YYYY-MM-DDTHH:MM:SS.NNNNNNNNNZ
		Level:   level.String(),
		Message: message,
		File:    fmt.Sprintf("%s:%d", file, line),
		Fields:  fields,
	}

	// 4. Convert the log entry to JSON
	data, err := json.Marshal(entry)
	if err != nil {
		// Fallback for conversion errors
		data = []byte(fmt.Sprintf(
			`{"time":"%s","level":"ERROR","message":"failed to convert log entry to json: %v"}`,
			time.Now().UTC().Format(time.RFC3339Nano),
			err,
		))
	}

	// 5. Lock the mutex to ensure safe writes in concurrency
	// "Everything is thread safe until you get serverless" - not Donald Knuth
	l.mu.Lock()
	defer l.mu.Unlock()

	// 6. Write the log
	_, _ = l.out.Write(data)
	_, _ = l.out.Write([]byte("\n")) // We add a newline for line-based log processing

	// 7. If the level is FATAL, terminate the program
	if level == FATAL {
		os.Exit(1)
	}
}

// --- Public Log Methods ---

// Debug
// Logs a message at the DEBUG level with optional structured fields.
func (l *Logger) Debug(message string, fields map[string]interface{}) {
	l.internalLog(DEBUG, message, fields)
}

// Info
// Logs a message at the INFO level with optional structured fields.
func (l *Logger) Info(message string, fields map[string]interface{}) {
	l.internalLog(INFO, message, fields)
}

// Warning
// Logs a message at the WARNING level with optional structured fields.
func (l *Logger) Warning(message string, fields map[string]interface{}) {
	l.internalLog(WARNING, message, fields)
}

// Error
// Logs a message at the ERROR level with optional structured fields.
func (l *Logger) Error(message string, fields map[string]interface{}) {
	l.internalLog(ERROR, message, fields)
}

// Fatal
// Logs a message at the FATAL level with optional structured fields, and then terminates.
func (l *Logger) Fatal(message string, fields map[string]interface{}) {
	l.internalLog(FATAL, message, fields)
}

// --- Public Formatted Methods ---

// Debugf
// Logs a formatted message at the DEBUG level.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.internalLog(DEBUG, fmt.Sprintf(format, v...), nil)
}

// Infof
// Logs a formatted message at the INFO level.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.internalLog(INFO, fmt.Sprintf(format, v...), nil)
}

// Warningf
// Logs a formatted message at the WARNING level.
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.internalLog(WARNING, fmt.Sprintf(format, v...), nil)
}

// Errorf
// Logs a formatted message at the ERROR level.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.internalLog(ERROR, fmt.Sprintf(format, v...), nil)
}

// Fatalf
// Logs a formatted message at the FATAL level, and then terminates.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.internalLog(FATAL, fmt.Sprintf(format, v...), nil)
}
