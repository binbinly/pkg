package logger

var logger = NewLogger()

// Logger is a generic logging interface.
type Logger interface {
	// Init initializes options
	Init(options ...Option) error
	// Fields set fields to always be logged
	Fields(fields map[string]any) Logger
	// Log writes a log entry
	Log(level string, v ...any)
	// Logf writes a formatted log entry
	Logf(level string, format string, v ...any)
}

func Init(opts ...Option) error {
	return logger.Init(opts...)
}

func Fields(fields map[string]any) Logger {
	return logger.Fields(fields)
}

func Log(level string, v ...any) {
	logger.Log(level, v...)
}

func Logf(level string, format string, v ...any) {
	logger.Logf(level, format, v...)
}
