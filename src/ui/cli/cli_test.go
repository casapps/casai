package cli

import (
	"testing"

	"github.com/casapps/casai/src/support"
)

func TestNewRootCommand(t *testing.T) {
	root := NewRootCommand()
	if root.Use != "casai" {
		t.Errorf("root.Use = %q, want %q", root.Use, "casai")
	}
	if !root.SilenceUsage || !root.SilenceErrors {
		t.Error("root command must silence cobra's default usage/error printing")
	}
}

func TestRunReturnsZeroOnSuccess(t *testing.T) {
	if code := Run([]string{"--help"}); code != 0 {
		t.Errorf("Run([--help]) = %d, want 0", code)
	}
}

func TestRunReturnsNonZeroOnUnknownFlag(t *testing.T) {
	if code := Run([]string{"--definitely-not-a-real-flag"}); code == 0 {
		t.Error("Run(unknown flag) = 0, want non-zero")
	}
}

func TestPersistentPreRunEResolvesDebugAndColor(t *testing.T) {
	root := NewRootCommand()
	root.SetArgs([]string{"--debug", "--color=no"})
	if err := root.ParseFlags([]string{"--debug", "--color=no"}); err != nil {
		t.Fatalf("ParseFlags: %v", err)
	}
	if err := root.PersistentPreRunE(root, nil); err != nil {
		t.Fatalf("PersistentPreRunE: %v", err)
	}
	if support.Logger == nil {
		t.Error("support.Logger was not set by PersistentPreRunE")
	}
	if support.ColorEnabled {
		t.Error("support.ColorEnabled = true, want false with --color=no")
	}
}
