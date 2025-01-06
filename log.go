package palantir

import (
	"context"
	"os"
	"unicode/utf8"

	"golang.org/x/exp/slog"
)

type LogEntry struct {
	Message   string
	ByteCount int
	CharCount int
	Text      string
	DomainID  string
	Ref       string
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

// NewLogEntry initializes a new LogEntry and calculates ByteCount and CharCount
func NewLogEntry(message, text, domainID, ref string) *LogEntry {
	return &LogEntry{
		Message:   message,
		Text:      text,
		DomainID:  domainID,
		Ref:       ref,
		ByteCount: len(text),
		CharCount: utf8.RuneCountInString(text),
	}
}

// Info logs an informational message
func (l *Logger) Info(ctx context.Context, entry LogEntry) {
	l.slogger.InfoContext(ctx, entry.Message, "byte_count", entry.ByteCount, "char_count", entry.CharCount, "text", entry.Text, "domain_id", entry.DomainID, "ref", entry.Ref)
}

// Debug logs a debug message
func (l *Logger) Debug(ctx context.Context, entry LogEntry) {
	l.slogger.DebugContext(ctx, entry.Message, "byte_count", entry.ByteCount, "char_count", entry.CharCount, "text", entry.Text, "domain_id", entry.DomainID, "ref", entry.Ref)
}

// Error logs an error message
func (l *Logger) Error(ctx context.Context, entry LogEntry) {
	l.slogger.ErrorContext(ctx, entry.Message, "byte_count", entry.ByteCount, "char_count", entry.CharCount, "text", entry.Text, "domain_id", entry.DomainID, "ref", entry.Ref)
}
