package errors

import (
	"fmt"
	"runtime"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	ErrorJSON = iota
	ErrorUnauthorized
	ErrorInvalidRequest
	ErrorHTTPRequest
	ErrorInvalidConfig
	ErrorInternal
)

const (
	SeverityDebug = iota
	SeverityInfo
	SeverityWarning
	SeverityError
	SeverityFatal
)

type AppError struct {
	// Error message to be displayed when Error() is called
	Message string
	// Context is a map for potentially useful extra information regarding the error
	Context map[string]interface{}
	// Stack info about where the error was created
	File string
	Func string
	// Error code for tracking what type of error this is
	Code     int
	Severity int
}

func (err *AppError) Error() string { return err.Message }

// LogEvent prepares a `*zerolog.Event` object with the error's information
func (err *AppError) LogEvent(logger zerolog.Logger) *zerolog.Event {
	event := logger.WithLevel(zerolog.Level(err.Severity)).
		Str("function", err.Func).
		Str("file", err.File)

	for key, val := range err.Context {
		event.Str(key, fmt.Sprintf("%v", val))
	}

	return event
}

func (err *AppError) Log(message string) {
	event := err.LogEvent(log.Logger)
	event.Msg(message)
}

// Multi slice for slurping up multiple errors and returning them at the end of some unit of work
type Multi []*AppError

// NewError creates a new application error and records the location of its creation.
func NewError(code int, message string, args ...interface{}) *AppError {
	formatted := fmt.Sprintf(message, args...)
	programCounter, filename, line, _ := runtime.Caller(1)
	return &AppError{
		Message: formatted,
		Code:    code,
		File:    fmt.Sprintf("%s:%d", filename, line),
		Func:    runtime.FuncForPC(programCounter).Name(),
		Context: make(map[string]interface{}),
	}
}
