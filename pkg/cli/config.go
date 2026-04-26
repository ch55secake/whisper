package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ch55secake/whisper/pkg/config"
)

// configCmd returns the `whisper config` sub-command tree.
func configCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage whisper configuration",
		Long: `Read and write the whisper config file (~/.config/whisper/config.yaml).

Available keys:
  server.host   Hostname or IP of the whisper server  (default: localhost)
  server.port   Port the server listens on             (default: 41002)
  username      Default username pre-filled at login   (default: "")`,
	}

	cmd.AddCommand(configShowCmd(), configGetCmd(), configSetCmd())
	return cmd
}

// configShowCmd prints all config keys and their current values.
func configShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Print all configuration values",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("config file: %s/config.yaml\n\n", config.Dir())
			for _, key := range config.Keys() {
				fmt.Printf("  %-15s = %s\n", key, config.Get(key))
			}
			return nil
		},
	}
}

// configGetCmd prints the value of a single key.
func configGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <key>",
		Short: "Get the value of a config key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]
			if !isKnownKey(key) {
				return fmt.Errorf("unknown key %q — run `whisper config show` to list valid keys", key)
			}
			fmt.Println(config.Get(key))
			return nil
		},
	}
}

// configSetCmd writes a new value for a key and saves the config file.
func configSetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set a config key and save to disk",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			key, value := args[0], args[1]
			if !isKnownKey(key) {
				return fmt.Errorf("unknown key %q — run `whisper config show` to list valid keys", key)
			}
			config.Set(key, value)
			if err := config.Save(); err != nil {
				_, err := fmt.Fprintf(os.Stderr, "whisper: %v\n", err)
				if err != nil {
					return err
				}
				return err
			}
			fmt.Printf("set %s = %s\n", key, value)
			return nil
		},
	}
}

func isKnownKey(key string) bool {
	for _, k := range config.Keys() {
		if strings.EqualFold(k, key) {
			return true
		}
	}
	return false
}
