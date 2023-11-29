package log

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// InitZeroLogger returns new instance of zero logger
func InitZeroLogger() zerolog.Logger {
	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}, // pretty print
	).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	return logger
}
