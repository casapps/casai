package support

import "testing"

func TestResolveColor(t *testing.T) {
	tests := []struct {
		name    string
		flag    string
		noColor bool
		wantYes bool
		wantNo  bool
	}{
		{"explicit yes wins even with NO_COLOR", "yes", true, true, false},
		{"explicit no wins", "no", false, false, true},
		{"auto with NO_COLOR set disables", "auto", true, false, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.noColor {
				t.Setenv("NO_COLOR", "1")
			}
			got := ResolveColor(tc.flag)
			if tc.wantYes && !got {
				t.Errorf("ResolveColor(%q) = false, want true", tc.flag)
			}
			if tc.wantNo && got {
				t.Errorf("ResolveColor(%q) = true, want false", tc.flag)
			}
		})
	}
}

func TestResolveColorAutoWithoutNoColor(t *testing.T) {
	// Just exercise the TTY auto-detection path without asserting a
	// specific result — that depends on how `go test` is invoked.
	_ = ResolveColor("auto")
}
