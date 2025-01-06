package palantir

import (
	"context"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := NewLogger()

	ctx := context.Background()

	// Info log test
	text1 := "Test text ทดสอบ"
	logger.Info(ctx, "Test log message", map[string]interface{}{
		"text":      text1,
		"domain_id": "123e4567-e89b-12d3-a456-426614174000",
		"ref":       "TestRef #1",
	})

	// Debug log test
	text2 := "Debug text ทดสอบ"
	logger.Debug(ctx, "Test debug message", map[string]interface{}{
		"text":      text2,
		"domain_id": "123e4567-e89b-12d3-a456-426614174000",
		"ref":       "DebugRef #2",
	})

	// Warn log test
	text3 := "Warning text ทดสอบ"
	logger.Warn(ctx, "Test warning message", map[string]interface{}{
		"text":      text3,
		"domain_id": "123e4567-e89b-12d3-a456-426614174000",
		"ref":       "WarnRef #3",
	})

	// Error log test
	text4 := "Error text ทดสอบ"
	logger.Error(ctx, context.DeadlineExceeded, map[string]interface{}{
		"text":      text4,
		"domain_id": "123e4567-e89b-12d3-a456-426614174000",
		"ref":       "ErrorRef #4",
	})

	t.Log("Logger tested successfully")
}
