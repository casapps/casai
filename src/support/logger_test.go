package support

import "testing"

func TestNewLoggerReturnsUsableLogger(t *testing.T) {
	for _, debug := range []bool{false, true} {
		logger := NewLogger(debug)
		if logger == nil {
			t.Fatalf("NewLogger(%v) = nil, want a logger", debug)
		}
		// Exercise the handler to make sure it doesn't panic; output goes
		// to stderr and isn't asserted here (format depends on the test
		// runner's stderr being a TTY or not).
		logger.Info("test log line", "debug", debug)
	}
}
