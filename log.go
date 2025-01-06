package palantir

import (
	"context"
	"os"
	"unicode/utf8"

	"golang.org/x/exp/slog"
)

// Attr defines a key-value pair for log attributes
type Attr struct {
	key string
	val interface{}
}

// String creates a string attribute
func String(k, v string) Attr {
	return Attr{k, v}
}

// Int creates an integer attribute
func Int(k string, v int) Attr {
	return Attr{k, v}
}

// Float64 creates a float64 attribute
func Float64(k string, v float64) Attr {
	return Attr{k, v}
}

// Bool creates a boolean attribute
func Bool(k string, v bool) Attr {
	return Attr{k, v}
}

// ILogger defines the logging interface
type ILogger interface {
	Info(ctx context.Context, message string, attrs ...Attr)
	Debug(ctx context.Context, message string, attrs ...Attr)
	Warn(ctx context.Context, message string, attrs ...Attr)
	Error(ctx context.Context, err error, attrs ...Attr)
}

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
func (l *Logger) Info(ctx context.Context, message string, attrs ...Attr) {
	l.log(ctx, slog.LevelInfo, message, attrs...)
}

// Debug logs a debug message
func (l *Logger) Debug(ctx context.Context, message string, attrs ...Attr) {
	l.log(ctx, slog.LevelDebug, message, attrs...)
}

// Warn logs a warning message
func (l *Logger) Warn(ctx context.Context, message string, attrs ...Attr) {
	l.log(ctx, slog.LevelWarn, message, attrs...)
}

// Error logs an error message
func (l *Logger) Error(ctx context.Context, err error, attrs ...Attr) {
	errorAttr := Attr{"error", err.Error()}
	l.log(ctx, slog.LevelError, "Error: "+err.Error(), append(attrs, errorAttr)...)
}

// log is a helper function to log messages at the specified level
func (l *Logger) log(ctx context.Context, level slog.Level, message string, attrs ...Attr) {
	attributes := make([]any, 0, len(attrs)*2)
	textValue := ""

	for _, attr := range attrs {
		attributes = append(attributes, attr.key, attr.val)
		if attr.key == "text" {
			if text, ok := attr.val.(string); ok {
				textValue = text
			}
		}
	}

	// If text is provided, calculate and append bytes_count and char_count
	if textValue != "" {
		attributes = append(attributes,
			"bytes_count", len(textValue),
			"char_count", utf8.RuneCountInString(textValue),
		)
	}

	l.slogger.Log(ctx, level, message, attributes...)
}
