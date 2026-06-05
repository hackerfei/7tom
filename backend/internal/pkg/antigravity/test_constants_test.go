package antigravity

import (
	"os"
	"testing"
)

const testAntigravityClientID = "test-antigravity-client-id"

func TestMain(m *testing.M) {
	defaultClientID = testAntigravityClientID
	defaultClientSecret = "test-antigravity-client-secret"
	os.Exit(m.Run())
}
