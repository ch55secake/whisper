// Package cli provides the cobra-based command-line interface for the whisper client.
package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ch55secake/whisper/pkg/client"
	"github.com/ch55secake/whisper/pkg/config"
)

// Execute is the single entry point called from main. It loads config then
// hands off to cobra.
func Execute() {
	if err := config.Load(); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "whisper: %v\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}

	if err := rootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}

// rootCmd is the default command — running `whisper` with no sub-command
// launches the TUI.
func rootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "whisper",
		Short: "A terminal chat client",
		Long:  "whisper — end-to-end encrypted terminal chat over gRPC.",
		RunE: func(cmd *cobra.Command, args []string) error {
			client.StartClient()
			return nil
		},
	}

	root.AddCommand(configCmd())
	return root
}
