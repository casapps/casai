package support

import (
	"log/slog"
	"os"
)

// Logger is the process-wide structured logger, set once at startup by the
// active UI adapter (see src/ui/cli.NewRootCommand's PersistentPreRunE) from
// the resolved --debug flag.
var Logger *slog.Logger

// ColorEnabled reports whether colored/emoji output is enabled for this
// run, set once at startup from the resolved --color flag (see
// ResolveColor). Defaults to false until an adapter resolves it.
var ColorEnabled bool

// NewLogger returns a structured logger. Text output on a TTY, JSON output
// otherwise (piped/redirected/automation contexts), per AI.md PART 3's
// CLI/automation output rules.
func NewLogger(debug bool) *slog.Logger {
	level := slog.LevelInfo
	if debug {
		level = slog.LevelDebug
	}

	opts := &slog.HandlerOptions{Level: level}

	if fi, err := os.Stderr.Stat(); err == nil && fi.Mode()&os.ModeCharDevice != 0 {
		return slog.New(slog.NewTextHandler(os.Stderr, opts))
	}
	return slog.New(slog.NewJSONHandler(os.Stderr, opts))
}
