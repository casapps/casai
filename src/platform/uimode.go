// Package platform provides OS/platform integration, including the
// GUI/TUI/CLI runtime-mode detection defined in AI.md PART 3.
package platform

import (
	"os"
	"strings"
)

// UIMode is the selected presentation surface for this invocation.
type UIMode int

const (
	UIModeCLI UIMode = iota
	UIModeTUI
	UIModeGUI
)

func (m UIMode) String() string {
	switch m {
	case UIModeGUI:
		return "gui"
	case UIModeTUI:
		return "tui"
	default:
		return "cli"
	}
}

// Capabilities describes what this build and this invocation can actually do.
// GUISupported/TUISupported reflect what was compiled in; the rest reflect
// the live invocation.
type Capabilities struct {
	GUISupported           bool
	TUISupported           bool
	LocalWindowsSession    bool
	LocalMacOSSession      bool
	InteractiveLaunch      bool
	StdinTTY               bool
	StdoutTTY              bool
	TUIRequiresFormatting  bool
	MachineOutputRequested bool
}

// Env is the environment/override source consulted during detection. The
// default implementation (Getenv) reads real process environment variables
// and explicit --ui overrides; tests supply a fake.
type Env interface {
	Has(key string) bool
	Is(key, value string) bool
	Truthy(key string) bool
	HasAny(keys ...string) bool
	// ForcedUIMode returns an explicit override (CLI flag > config > env),
	// per AI.md PART 3 → "Override priority".
	ForcedUIMode() (UIMode, bool)
}

// DetectUIMode implements AI.md PART 3 → "Reference Detection Logic":
// explicit override, then GUI, then TUI, else CLI.
func DetectUIMode(env Env, caps Capabilities) UIMode {
	if forced, ok := env.ForcedUIMode(); ok {
		return forced
	}

	remoteShell := env.HasAny("SSH_CONNECTION", "SSH_CLIENT", "SSH_TTY", "MOSH_IP", "MOSH_KEY")
	displayAvail := env.Has("WAYLAND_DISPLAY") || env.Has("DISPLAY") ||
		caps.LocalWindowsSession || caps.LocalMacOSSession

	if caps.GUISupported && !remoteShell && displayAvail && caps.InteractiveLaunch {
		return UIModeGUI
	}

	plainOrNoninteractive := !caps.StdinTTY || !caps.StdoutTTY ||
		env.Is("TERM", "dumb") || env.Truthy("CI") ||
		(env.Truthy("NO_COLOR") && caps.TUIRequiresFormatting) ||
		caps.MachineOutputRequested

	if caps.TUISupported && !plainOrNoninteractive {
		return UIModeTUI
	}

	return UIModeCLI
}

// osEnv is the real-process Env implementation.
type osEnv struct {
	forced   UIMode
	hasForce bool
}

// NewOSEnv builds an Env backed by real process environment variables.
// forcedFlag is the --ui flag value ("gui"/"tui"/"cli"), empty when unset.
func NewOSEnv(forcedFlag string) Env {
	e := osEnv{}
	switch strings.ToLower(strings.TrimSpace(forcedFlag)) {
	case "gui":
		e.forced, e.hasForce = UIModeGUI, true
	case "tui":
		e.forced, e.hasForce = UIModeTUI, true
	case "cli":
		e.forced, e.hasForce = UIModeCLI, true
	}
	return e
}

func (e osEnv) Has(key string) bool { return os.Getenv(key) != "" }

func (e osEnv) Is(key, value string) bool { return os.Getenv(key) == value }

func (e osEnv) Truthy(key string) bool {
	switch strings.ToLower(os.Getenv(key)) {
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}

func (e osEnv) HasAny(keys ...string) bool {
	for _, k := range keys {
		if e.Has(k) {
			return true
		}
	}
	return false
}

func (e osEnv) ForcedUIMode() (UIMode, bool) { return e.forced, e.hasForce }

// StdTTYCapabilities inspects the real stdin/stdout of this process.
func StdTTYCapabilities() (stdinTTY, stdoutTTY bool) {
	if fi, err := os.Stdin.Stat(); err == nil {
		stdinTTY = fi.Mode()&os.ModeCharDevice != 0
	}
	if fi, err := os.Stdout.Stat(); err == nil {
		stdoutTTY = fi.Mode()&os.ModeCharDevice != 0
	}
	return stdinTTY, stdoutTTY
}
