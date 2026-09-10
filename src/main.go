// Command casai is the single entry point for casai's GUI, TUI, and CLI
// surfaces. It auto-detects which surface to present (AI.md PART 3) unless
// explicitly overridden by flag, config, or environment.
package main

import (
	"os"

	"github.com/casapps/casai/src/platform"
	"github.com/casapps/casai/src/support"
	"github.com/casapps/casai/src/ui/cli"
)

// Build info - set via `-ldflags -X main.Version=...` (AI.md PART 6 →
// "Build Metadata"); ldflags -X only works against string vars declared
// directly in the target package, so these live here and are copied into
// the support package (the shared source every UI adapter reads from) in
// init below.
var (
	Version      = "devel"
	CommitID     = "N/A"
	BuildEpoch   = "0"
	OfficialSite = ""
)

func init() {
	support.Version = Version
	support.CommitID = CommitID
	support.BuildEpoch = BuildEpoch
	support.OfficialSite = OfficialSite
	support.DeriveBuildDate()
}

func main() {
	os.Exit(run(os.Args[1:]))
}

// run selects a UI surface and dispatches to it. GUI and TUI adapters are
// not yet implemented (see TODO.AI.md); until they land, any GUI/TUI
// selection falls back to the CLI surface, which is the required baseline.
func run(args []string) int {
	uiFlag := extractUIFlag(args)

	stdinTTY, stdoutTTY := platform.StdTTYCapabilities()
	caps := platform.Capabilities{
		// GUISupported becomes true once src/ui/gui is implemented.
		GUISupported: false,
		// TUISupported becomes true once src/ui/tui is implemented.
		TUISupported:           false,
		InteractiveLaunch:      stdinTTY && stdoutTTY,
		StdinTTY:               stdinTTY,
		StdoutTTY:              stdoutTTY,
		TUIRequiresFormatting:  true,
		MachineOutputRequested: false,
	}

	switch platform.DetectUIMode(platform.NewOSEnv(uiFlag), caps) {
	case platform.UIModeGUI, platform.UIModeTUI:
		// Neither surface is implemented yet; CLI is the mandatory fallback.
		return cli.Run(args)
	default:
		return cli.Run(args)
	}
}

// extractUIFlag does a minimal pre-scan for --ui=<mode>/--ui <mode> so mode
// detection can honor it before cobra parses the rest of the CLI's flags.
func extractUIFlag(args []string) string {
	for i, a := range args {
		switch {
		case a == "--ui" && i+1 < len(args):
			return args[i+1]
		case len(a) > 5 && a[:5] == "--ui=":
			return a[5:]
		}
	}
	return ""
}
