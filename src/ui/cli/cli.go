// Package cli implements casai's non-interactive command-line surface,
// the required baseline per AI.md PART 3 → "CLI" and IDEA.md → "App
// surfaces in scope".
package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/casapps/casai/src/support"
)

// NewRootCommand builds the casai root cobra command.
func NewRootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:           "casai",
		Short:         "casai is a model-agnostic, Claude-Code-compatible AI coding agent",
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       support.Version,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			debug, err := cmd.Flags().GetBool("debug")
			if err != nil {
				return err
			}
			colorFlag, err := cmd.Flags().GetString("color")
			if err != nil {
				return err
			}
			support.Logger = support.NewLogger(debug)
			support.ColorEnabled = support.ResolveColor(colorFlag)
			return nil
		},
	}

	root.PersistentFlags().Bool("debug", false, "Enable debug output")
	root.PersistentFlags().String("color", "auto", "Color output: auto, yes, no")
	// Cobra only auto-creates a bare --version flag from Command.Version;
	// AI.md PART 7 requires the -v shorthand too, so define it ourselves —
	// cobra skips its own registration when a "version" flag already exists.
	root.Flags().BoolP("version", "v", false, "Show version and exit")

	root.SetVersionTemplate(
		fmt.Sprintf("casai %s (commit %s, built %s)\n", support.Version, support.CommitID, support.BuildDate),
	)

	return root
}

// Run executes the CLI surface for the given args (excluding argv[0]).
func Run(args []string) int {
	root := NewRootCommand()
	root.SetArgs(args)
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return 1
	}
	return 0
}
