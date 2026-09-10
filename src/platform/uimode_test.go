package platform

import "testing"

// fakeEnv is a test-only Env implementation.
type fakeEnv struct {
	vals     map[string]string
	forced   UIMode
	hasForce bool
}

func (f fakeEnv) Has(key string) bool { return f.vals[key] != "" }

func (f fakeEnv) Is(key, value string) bool { return f.vals[key] == value }

func (f fakeEnv) Truthy(key string) bool {
	switch f.vals[key] {
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}

func (f fakeEnv) HasAny(keys ...string) bool {
	for _, k := range keys {
		if f.Has(k) {
			return true
		}
	}
	return false
}

func (f fakeEnv) ForcedUIMode() (UIMode, bool) { return f.forced, f.hasForce }

func TestUIModeString(t *testing.T) {
	cases := map[UIMode]string{
		UIModeCLI:  "cli",
		UIModeTUI:  "tui",
		UIModeGUI:  "gui",
		UIMode(99): "cli",
	}
	for mode, want := range cases {
		if got := mode.String(); got != want {
			t.Errorf("UIMode(%d).String() = %q, want %q", mode, got, want)
		}
	}
}

func TestDetectUIMode(t *testing.T) {
	tests := []struct {
		name string
		env  fakeEnv
		caps Capabilities
		want UIMode
	}{
		{
			name: "explicit override wins over everything",
			env:  fakeEnv{forced: UIModeTUI, hasForce: true},
			caps: Capabilities{GUISupported: true, TUISupported: true},
			want: UIModeTUI,
		},
		{
			name: "GUI selected when supported, local display, interactive",
			env:  fakeEnv{vals: map[string]string{"DISPLAY": ":0"}},
			caps: Capabilities{GUISupported: true, InteractiveLaunch: true},
			want: UIModeGUI,
		},
		{
			name: "GUI skipped over SSH even with a display",
			env:  fakeEnv{vals: map[string]string{"DISPLAY": ":0", "SSH_TTY": "/dev/pts/0"}},
			caps: Capabilities{GUISupported: true, TUISupported: true, InteractiveLaunch: true, StdinTTY: true, StdoutTTY: true},
			want: UIModeTUI,
		},
		{
			name: "GUI skipped when not compiled in, falls to TUI",
			env:  fakeEnv{vals: map[string]string{"DISPLAY": ":0"}},
			caps: Capabilities{GUISupported: false, TUISupported: true, InteractiveLaunch: true, StdinTTY: true, StdoutTTY: true},
			want: UIModeTUI,
		},
		{
			name: "TUI skipped when stdout is piped",
			env:  fakeEnv{},
			caps: Capabilities{TUISupported: true, StdinTTY: true, StdoutTTY: false},
			want: UIModeCLI,
		},
		{
			name: "TUI skipped under CI",
			env:  fakeEnv{vals: map[string]string{"CI": "true"}},
			caps: Capabilities{TUISupported: true, StdinTTY: true, StdoutTTY: true},
			want: UIModeCLI,
		},
		{
			name: "TUI skipped when TERM=dumb",
			env:  fakeEnv{vals: map[string]string{"TERM": "dumb"}},
			caps: Capabilities{TUISupported: true, StdinTTY: true, StdoutTTY: true},
			want: UIModeCLI,
		},
		{
			name: "TUI skipped when NO_COLOR set and formatting required",
			env:  fakeEnv{vals: map[string]string{"NO_COLOR": "1"}},
			caps: Capabilities{TUISupported: true, StdinTTY: true, StdoutTTY: true, TUIRequiresFormatting: true},
			want: UIModeCLI,
		},
		{
			name: "TUI selected on a plain interactive terminal",
			env:  fakeEnv{},
			caps: Capabilities{TUISupported: true, StdinTTY: true, StdoutTTY: true},
			want: UIModeTUI,
		},
		{
			name: "CLI is the default fallback",
			env:  fakeEnv{},
			caps: Capabilities{},
			want: UIModeCLI,
		},
		{
			name: "machine output request forces CLI over TUI",
			env:  fakeEnv{},
			caps: Capabilities{TUISupported: true, StdinTTY: true, StdoutTTY: true, MachineOutputRequested: true},
			want: UIModeCLI,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := DetectUIMode(tc.env, tc.caps); got != tc.want {
				t.Errorf("DetectUIMode() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestNewOSEnvForcedUIMode(t *testing.T) {
	tests := []struct {
		flag     string
		wantMode UIMode
		wantOK   bool
	}{
		{"gui", UIModeGUI, true},
		{"TUI", UIModeTUI, true},
		{"  cli  ", UIModeCLI, true},
		{"", UIModeCLI, false},
		{"bogus", UIModeCLI, false},
	}

	for _, tc := range tests {
		env := NewOSEnv(tc.flag)
		mode, ok := env.ForcedUIMode()
		if ok != tc.wantOK {
			t.Errorf("NewOSEnv(%q) ForcedUIMode ok = %v, want %v", tc.flag, ok, tc.wantOK)
		}
		if ok && mode != tc.wantMode {
			t.Errorf("NewOSEnv(%q) ForcedUIMode mode = %v, want %v", tc.flag, mode, tc.wantMode)
		}
	}
}

func TestOSEnvHasIsTruthyHasAny(t *testing.T) {
	t.Setenv("CASAI_TEST_KEY", "true")
	t.Setenv("CASAI_TEST_EMPTY", "")

	env := NewOSEnv("")

	if !env.Has("CASAI_TEST_KEY") {
		t.Error("Has(CASAI_TEST_KEY) = false, want true")
	}
	if env.Has("CASAI_TEST_EMPTY") {
		t.Error("Has(CASAI_TEST_EMPTY) = true, want false")
	}
	if !env.Is("CASAI_TEST_KEY", "true") {
		t.Error("Is(CASAI_TEST_KEY, true) = false, want true")
	}
	if !env.Truthy("CASAI_TEST_KEY") {
		t.Error("Truthy(CASAI_TEST_KEY) = false, want true")
	}
	if !env.HasAny("CASAI_TEST_EMPTY", "CASAI_TEST_KEY") {
		t.Error("HasAny(...) = false, want true")
	}
	if env.HasAny("CASAI_TEST_EMPTY", "CASAI_TEST_NOPE") {
		t.Error("HasAny(...) = true, want false")
	}
}

func TestStdTTYCapabilities(t *testing.T) {
	// In the test harness stdin/stdout are typically not a character
	// device; this just exercises the code path without asserting a
	// specific TTY state (that depends on how `go test` is invoked).
	_, _ = StdTTYCapabilities()
}
