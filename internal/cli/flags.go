package cli

import (
	"flag"
	"os"
	"path/filepath"
)

// Flags holds all CLI flags
type Flags struct {
	ConfigPath string
}

// defaultConfigPath returns ~/.inspektra/config.yaml
func defaultConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "config.yaml"
	}

	return filepath.Join(home, ".inspektra", "config.yaml")
}

// ParseFlags parses command-line flags
func ParseFlags() *Flags {
	flags := &Flags{}

	flag.StringVar(
		&flags.ConfigPath,
		"config",
		defaultConfigPath(),
		"Path to configuration file",
	)

	flag.Parse()
	return flags
}
