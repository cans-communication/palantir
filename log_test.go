package palantir

import (
	"context"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := NewLogger()

	ctx := context.Background()

	// Info log test
	logger.Info(ctx, "Test log message",
		String("text", "Test text ทดสอบ"),
		String("domain_id", "123e4567-e89b-12d3-a456-426614174000"),
		String("ref", "TestRef #1"),
	)

	// Debug log test
	logger.Debug(ctx, "Test debug message",
		String("text", "Debug text ทดสอบ"),
		String("domain_id", "123e4567-e89b-12d3-a456-426614174000"),
		String("ref", "DebugRef #2"),
	)

	// Warn log test
	logger.Warn(ctx, "Test warning message",
		String("text", "Warning text ทดสอบ"),
		String("domain_id", "123e4567-e89b-12d3-a456-426614174000"),
		String("ref", "WarnRef #3"),
	)

	// Error log test
	logger.Error(ctx, context.DeadlineExceeded,
		String("text", "Error text ทดสอบ"),
		String("domain_id", "123e4567-e89b-12d3-a456-426614174000"),
		String("ref", "ErrorRef #4"),
	)

	t.Log("Logger tested successfully")
}
