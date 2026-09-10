package support

import "os"

// ResolveColor decides whether colored/emoji output should be enabled, per
// the standard precedence order (AI.md PART 7 / global Go conventions):
//  1. --color=yes or --color=no (explicit flag) always wins
//  2. NO_COLOR env var present (any value) disables color, no-color.org
//  3. --color=auto (default) falls back to TTY auto-detection on stdout
func ResolveColor(flag string) bool {
	switch flag {
	case "yes":
		return true
	case "no":
		return false
	default:
		if _, set := os.LookupEnv("NO_COLOR"); set {
			return false
		}
		fi, err := os.Stdout.Stat()
		return err == nil && fi.Mode()&os.ModeCharDevice != 0
	}
}
