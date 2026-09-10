package main

import "testing"

func TestExtractUIFlag(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{"no flag", []string{"status"}, ""},
		{"space-separated", []string{"--ui", "tui", "status"}, "tui"},
		{"space-separated at end with no value", []string{"status", "--ui"}, ""},
		{"equals form", []string{"--ui=gui"}, "gui"},
		{"equals form empty value", []string{"--ui="}, ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := extractUIFlag(tc.args); got != tc.want {
				t.Errorf("extractUIFlag(%v) = %q, want %q", tc.args, got, tc.want)
			}
		})
	}
}

func TestRunFallsBackToCLI(t *testing.T) {
	if code := run([]string{"--help"}); code != 0 {
		t.Errorf("run([--help]) = %d, want 0", code)
	}
}
