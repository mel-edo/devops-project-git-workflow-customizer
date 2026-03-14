package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Workflow      string   `yaml:"workflow"`
	DefaultBranch string   `yaml:"default_branch"`
	WebhookURL    string   `yaml:"webhook_url"`
	Repositories  []string `yaml:"repositories"`
}

func expandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return path
		}
		return filepath.Join(home, path[2:])
	}
	return path
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// expand ~ in all repo paths
	for i, r := range cfg.Repositories {
		cfg.Repositories[i] = expandPath(r)
	}

	if err := validate(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func LoadDefaultConfig() (*Config, error) {
	candidates := []string{
		"seryn.yaml",
		"seryn.yml",
		".seryn.yaml",
	}

	for _, name := range candidates {
		if _, err := os.Stat(name); err == nil {
			return LoadConfig(name)
		}
	}

	return nil, nil // no default config found, not an error
}

func validate(cfg *Config) error {
	if cfg.Workflow == "" {
		return fmt.Errorf("config error: 'workflow' field is required")
	}

	if len(cfg.Repositories) == 0 {
		return fmt.Errorf("config error: 'repositories' list must not be empty")
	}

	supported := map[string]bool{
		"centralized": true,
		"feature":     true,
		"gitflow":     true,
		"forking":     true,
	}

	if !supported[cfg.Workflow] {
		return fmt.Errorf("config error: unsupported workflow: '%s'", cfg.Workflow)
	}

	return nil
}
