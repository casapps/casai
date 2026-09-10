// Package support holds logging, error helpers, and build/version metadata
// shared by every UI surface (GUI/TUI/CLI). See AI.md PART 6 for the
// metadata injection contract this file implements.
package support

import (
	"strconv"
	"time"
)

// Build info - BuildEpoch is set via -ldflags at build time; BuildDate is derived from it.
var (
	// Version is overridden by: -ldflags="-X main.Version=$(cat release.txt)"
	Version = "devel"
	// CommitID is overridden by: -ldflags="-X main.CommitID=$(git rev-parse --short=7 HEAD)"
	CommitID = "N/A"
	// BuildDate is derived from BuildEpoch in init(); "N/A" when BuildEpoch is unset.
	BuildDate = "N/A"
	// BuildEpoch is overridden by: -ldflags="-X 'main.BuildEpoch=${BUILD_EPOCH}'" - Unix build timestamp (seconds, UTC).
	BuildEpoch = "0"
	// OfficialSite is overridden by: -ldflags="-X main.OfficialSite=$(cat site.txt)"
	OfficialSite = ""
)

// buildEpoch parses the embedded BuildEpoch ldflag; 0 when unset or invalid.
func buildEpoch() int64 {
	n, err := strconv.ParseInt(BuildEpoch, 10, 64)
	if err != nil {
		return 0
	}
	return n
}

// DeriveBuildDate recomputes BuildDate (RFC 3339 UTC) from BuildEpoch. The
// ldflags -X mechanism can only target string vars declared in package
// main, so main.go copies its own Version/CommitID/BuildEpoch/OfficialSite
// into this package at startup and then calls DeriveBuildDate.
func DeriveBuildDate() {
	if n := buildEpoch(); n > 0 {
		BuildDate = time.Unix(n, 0).UTC().Format("2006-01-02T15:04:05Z")
	}
}
