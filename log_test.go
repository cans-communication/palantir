package palantir

import (
	"context"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := NewLogger()

	// Use NewLogEntry to automatically set ByteCount and CharCount
	testEntry := NewLogEntry(
		"Test log message",
		"Test text ทดสอบ",
		"123e4567-e89b-12d3-a456-426614174000",
		"TestRef #1",
	)

	ctx := context.Background()
	logger.Info(ctx, *testEntry)
	logger.Debug(ctx, *testEntry)
	logger.Error(ctx, *testEntry)

	t.Log("Logger tested successfully")
}
