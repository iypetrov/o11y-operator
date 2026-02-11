package testutils

import (
	"testing"

	"go.uber.org/goleak"
)

// VerifyGoLeaks verifies that unit tests do not leak any goroutines.
// It should be called in TestMain.
func VerifyGoLeaks(m *testing.M) {
	goleak.VerifyTestMain(m)
}
