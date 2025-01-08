package palantir_test

import (
	"context"
	"testing"

	"github.com/cans-communication/palantir"
)

func TestLogger(t *testing.T) {
	ctx := context.Background()

	// Info log test
	palantir.Info(ctx, "Test log message",
		palantir.String("text", "Test text ทดสอบ"),
		palantir.String("domain_id", "123e4567-e89b-12d3-a456-426614174000"),
		palantir.String("ref", "TestRef #1"),
	)

	// Debug log test
	palantir.Debug(ctx, "Test debug message",
		palantir.String("text", "Debug text ทดสอบ"),
		palantir.String("domain_id", "123e4567-e89b-12d3-a456-426614174000"),
		palantir.String("ref", "DebugRef #2"),
	)

	// Warn log test
	palantir.Warn(ctx, "Test warning message",
		palantir.String("text", "Warning text ทดสอบ"),
		palantir.String("domain_id", "123e4567-e89b-12d3-a456-426614174000"),
		palantir.String("ref", "WarnRef #3"),
	)

	// Error log test
	palantir.Error(ctx, context.DeadlineExceeded,
		palantir.String("text", "Error text ทดสอบ"),
		palantir.String("domain_id", "123e4567-e89b-12d3-a456-426614174000"),
		palantir.String("ref", "ErrorRef #4"),
	)

	t.Log("Logger tested successfully")
}
