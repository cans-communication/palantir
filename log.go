package palantir

import (
	"context"
	"os"

	"golang.org/x/exp/slog"
)

type ValueType int

const (
	ValueTypeString ValueType = iota
	ValueTypeInt
	ValueTypeFloat64
	ValueTypeBool
)

// Attr defines a key-value pair for log attributes
type Attr struct {
	valueType ValueType
	key       string
	val       interface{}
}

// String creates a string attribute
func String(k, v string) Attr {
	return Attr{ValueTypeString, k, v}
}

// Int creates an integer attribute
func Int(k string, v int) Attr {
	return Attr{ValueTypeInt, k, v}
}

// Float64 creates a float64 attribute
func Float64(k string, v float64) Attr {
	return Attr{ValueTypeFloat64, k, v}
}

// Bool creates a boolean attribute
func Bool(k string, v bool) Attr {
	return Attr{ValueTypeBool, k, v}
}

var slogger *slog.Logger

func init() {

	level := slog.LevelInfo

	switch os.Getenv("LOG_LEVEL") {
	case "DEBUG":
		level = slog.LevelDebug
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})

	slogger = slog.New(handler)
}

// Info logs an informational message
func Info(ctx context.Context, message string, attrs ...Attr) {
	log(ctx, slog.LevelInfo, message, attrs...)
}

// Debug logs a debug message
func Debug(ctx context.Context, message string, attrs ...Attr) {
	log(ctx, slog.LevelDebug, message, attrs...)
}

// Warn logs a warning message
func Warn(ctx context.Context, message string, attrs ...Attr) {
	log(ctx, slog.LevelWarn, message, attrs...)
}

// Error logs an error message
func Error(ctx context.Context, err error, attrs ...Attr) {
	log(ctx, slog.LevelError, "Error: "+err.Error(), attrs...)
}

// log is a helper function to log messages at the specified level
func log(ctx context.Context, level slog.Level, message string, attrs ...Attr) {

	if slogger == nil {
		return
	}

	sattrs := make([]any, 0, len(attrs))

	for _, attr := range attrs {
		var sval slog.Value

		switch attr.valueType {
		case ValueTypeString:
			val, _ := attr.val.(string)
			sval = slog.StringValue(val)
		case ValueTypeInt:
			val, _ := attr.val.(int)
			sval = slog.IntValue(val)
		case ValueTypeFloat64:
			val, _ := attr.val.(float64)
			sval = slog.Float64Value(val)
		case ValueTypeBool:
			val, _ := attr.val.(bool)
			sval = slog.BoolValue(val)
		default:
			sval = slog.StringValue("")
		}

		sattrs = append(sattrs,
			slog.Attr{Key: attr.key, Value: sval},
		)
	}

	slogger.Log(ctx, level, message, sattrs...)
}
