package proxy

import (
	"syscall"

	"github.com/AdguardTeam/golibs/errors"
)

// isEPIPE checks if the underlying error is EPIPE.  On Plan 9, we check for
// the error string since syscall.EPIPE doesn't exist there.
func isEPIPE(err error) (ok bool) {
	if errors.Is(err, syscall.EPIPE) {
		return true
	}

	return containsString(err.Error(), "write on closed pipe")
}

// containsString is a simple string contains check.
func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsSubstring(s, substr))
}

// containsSubstring checks if s contains substr.
func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}

	return false
}
