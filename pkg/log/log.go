package log

import "github.com/rs/zerolog"

// ILoggerGeneral general contract methods that fits most of the loggers (zerolog, logrus...)
type ILoggerGeneral interface {
	Printf(format string, v interface{})
	Print(v interface{})
	Fatal(v string)
	Fatalf(format string, v interface{})
	Panic(v string)
	Panicf(format string, v interface{})
}

// IZeroLogger zero logger specific contract
type IZeroLogger interface {
	Error(err error)
	Errorf(format string, v interface{})
}

// ILogger contract
type ILogger interface {
	ILoggerGeneral
	IZeroLogger
}

// zerolog Adapter
type zereoLoggerAdapter struct {
	zeroLogger zerolog.Logger
}

// NewLogger returns zeroLogger
func NewZeroLogger(zeroLogger zerolog.Logger) ILogger {
	return &zereoLoggerAdapter{
		zeroLogger: zeroLogger,
	}
}

// Printf sends a log event using debug level and no extra field. Arguments are handled in the manner of fmt.Printf.
func (l zereoLoggerAdapter) Printf(format string, v interface{}) {
	l.zeroLogger.Printf(format, v)
}

// Print sends a log event using debug level and no extra field. Arguments are handled in the manner of fmt.Print.
func (l zereoLoggerAdapter) Print(v interface{}) {
	l.zeroLogger.Print(v)
}

// Fatal starts a new message with fatal level. The os.Exit(1) function is called by the Msg method, which terminates the program immediately.
func (l zereoLoggerAdapter) Fatal(v string) {
	l.zeroLogger.Fatal().Msg(v)
}

// Fatal starts a new message with fatal level. The os.Exit(1) function is called by the Msg method, which terminates the program immediately.
func (l zereoLoggerAdapter) Fatalf(format string, v interface{}) {
	l.zeroLogger.Fatal().Msgf(format, v)
}

// Panic starts a new message with panic level. The panic() function is called by the Msg method, which stops the ordinary flow of a goroutine.
func (l zereoLoggerAdapter) Panic(v string) {
	l.zeroLogger.Panic().Msg(v)
}

// Panic starts a new message with panic level. The panic() function is called by the Msg method, which stops the ordinary flow of a goroutine.
func (l zereoLoggerAdapter) Panicf(format string, v interface{}) {
	l.zeroLogger.Panic().Msgf(format, v)
}

// Error starts a new message with error level. Err adds the field "error" with serialized err to the *Event context. If err is nil, no field is added.
func (l zereoLoggerAdapter) Error(err error) {
	l.zeroLogger.Error().Err(err)
}

// Error starts a new message with error level. Msgf sends the event with formatted msg added as the message field if not empty.
func (l zereoLoggerAdapter) Errorf(format string, v interface{}) {
	l.zeroLogger.Error().Msgf(format, v)
}
