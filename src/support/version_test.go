package support

import "testing"

func TestDeriveBuildDate(t *testing.T) {
	origEpoch, origDate := BuildEpoch, BuildDate
	defer func() { BuildEpoch, BuildDate = origEpoch, origDate }()

	tests := []struct {
		name  string
		epoch string
		want  string
	}{
		{"unset epoch stays N/A", "0", "N/A"},
		{"invalid epoch stays N/A", "not-a-number", "N/A"},
		{"valid epoch derives RFC3339 UTC", "1700000000", "2023-11-14T22:13:20Z"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			BuildEpoch = tc.epoch
			BuildDate = "N/A"
			DeriveBuildDate()
			if BuildDate != tc.want {
				t.Errorf("DeriveBuildDate() with epoch %q => BuildDate = %q, want %q", tc.epoch, BuildDate, tc.want)
			}
		})
	}
}
