// Package config manages whisper client configuration backed by viper.
// The config file is stored at $HOME/.config/whisper/config.yaml.
package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	KeyServerHost = "server.host"
	KeyServerPort = "server.port"
	KeyUsername   = "username"

	defaultHost = "localhost"
	defaultPort = 41002
)

// Load initialises viper, sets defaults, and reads the config file if present.
// It does not return an error when the file simply does not exist yet.
func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(Dir())

	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) {
			return fmt.Errorf("reading config: %w", err)
		}
	}

	return nil
}

// Save writes the current viper state back to the config file, creating
// the config directory if it does not exist yet.
func Save() error {
	if err := os.MkdirAll(Dir(), 0o755); err != nil {
		return fmt.Errorf("creating config dir: %w", err)
	}

	path := filepath.Join(Dir(), "config.yaml")
	if err := viper.WriteConfigAs(path); err != nil {
		return fmt.Errorf("writing config: %w", err)
	}

	return nil
}

// Dir returns the directory that holds the config file.
func Dir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ".whisper"
	}
	return filepath.Join(home, ".config", "whisper")
}

// ServerAddress returns the full host:port address for the gRPC server.
func ServerAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString(KeyServerHost), viper.GetInt(KeyServerPort))
}

// Username returns the saved default username or an empty string if unset.
func Username() string {
	return viper.GetString(KeyUsername)
}

// Set updates a key in viper (does not persist — call Save to write to disk).
func Set(key, value string) {
	viper.Set(key, value)
}

// Get returns the string value for a key.
func Get(key string) string {
	return viper.GetString(key)
}

// Keys returns all known config keys.
func Keys() []string {
	return []string{KeyServerHost, KeyServerPort, KeyUsername}
}

func setDefaults() {
	viper.SetDefault(KeyServerHost, defaultHost)
	viper.SetDefault(KeyServerPort, defaultPort)
	viper.SetDefault(KeyUsername, "")
}
