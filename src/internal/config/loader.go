package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Workflow       string   `yaml:"workflow"`
	DefaultBranch  string   `yaml:"default_branch"`
	RequireReviews bool     `yaml:"require_reviews"`
	WebhookURL     string   `yaml:"webhook_url"`
	Repositories   []string `yaml:"repositories"`
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

	if err := validate(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
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
