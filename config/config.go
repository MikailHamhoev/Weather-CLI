// weather-cli/config/config.go
package config

import (
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	APIKey string
}

func LoadConfig() (*Config, error) {
	// Try to get API key from environment variable first
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		// Try to read from config file
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("failed to get home directory: %w", err)
		}

		configPath := filepath.Join(homeDir, ".weather-cli", "config")
		data, err := os.ReadFile(configPath)
		if err != nil {
			if os.IsNotExist(err) {
				return nil, fmt.Errorf("API key not found. Set WEATHER_API_KEY environment variable or create %s", configPath)
			}
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}

		apiKey = string(data)
	}

	if apiKey == "" {
		return nil, fmt.Errorf("API key is empty")
	}

	return &Config{APIKey: apiKey}, nil
}

func SaveConfig(apiKey string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".weather-cli")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	configPath := filepath.Join(configDir, "config")
	if err := os.WriteFile(configPath, []byte(apiKey), 0600); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
