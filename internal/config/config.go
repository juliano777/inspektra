package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the full application configuration
type Config struct {
	Postgres PostgresConfig `yaml:"postgres"`
}

// PostgresConfig holds PostgreSQL connection parameters
type PostgresConfig struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	Database        string `yaml:"database"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	SSLMode         string `yaml:"sslmode"`
	ApplicationName string `yaml:"application_name"`
}

// Load loads the configuration from YAML configuration file
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file %s: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("cannot parse config file %s: %w", path, err)
	}

	return &cfg, nil
}
