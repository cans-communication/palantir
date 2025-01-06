package palantir

import (
	"context"
	"os"
	"unicode/utf8"

	"golang.org/x/exp/slog"
)

// Logger wraps slog.Logger for custom logging
type Logger struct {
	slogger *slog.Logger
}

// NewLogger initializes a new Logger
func NewLogger() *Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	return &Logger{slogger: slog.New(handler)}
}

// Info logs an informational message
func (l *Logger) Info(ctx context.Context, message string, args map[string]interface{}) {
	l.log(ctx, slog.LevelInfo, message, args)
}

// Debug logs a debug message
func (l *Logger) Debug(ctx context.Context, message string, args map[string]interface{}) {
	l.log(ctx, slog.LevelDebug, message, args)
}

// Warn logs a warning message
func (l *Logger) Warn(ctx context.Context, message string, args map[string]interface{}) {
	l.log(ctx, slog.LevelWarn, message, args)
}

// Error logs an error message
func (l *Logger) Error(ctx context.Context, err error, args map[string]interface{}) {
	if args == nil {
		args = make(map[string]interface{})
	}
	args["error"] = err.Error()
	l.log(ctx, slog.LevelError, "Error: "+err.Error(), args)
}

// log is a helper function to log messages at the specified level
func (l *Logger) log(ctx context.Context, level slog.Level, message string, args map[string]interface{}) {
	// Check for "text" key and calculate byte_count and char_count
	if text, ok := args["text"].(string); ok {
		args["byte_count"] = len(text)
		args["char_count"] = utf8.RuneCountInString(text)
	}

	attrs := make([]any, 0, len(args)*2)
	for key, value := range args {
		attrs = append(attrs, key, value)
	}
	l.slogger.Log(ctx, level, message, attrs...)
}
