package config

import (
	"encoding/json"
	"flag"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
)

// Config holds all configuration for the little-share server
type Config struct {
	// Server configuration
	Host string `json:"host"`
	Port int    `json:"port"`

	// Storage configuration
	FrontendPath  string `json:"frontend_path"`
	FileStorePath string `json:"file_store_path"`

	// Storage limits (in bytes, 0 means unlimited)
	MaxStorageSize int64 `json:"max_storage_size"`
}

// DefaultConfig returns a Config with sensible defaults
func DefaultConfig() Config {
	return Config{
		Host:           "0.0.0.0",
		Port:           19999,
		FrontendPath:   "front",
		FileStorePath:  "files",
		MaxStorageSize: 10 * 1024 * 1024 * 1024, // 10GB
	}
}

// Load loads configuration from a JSON file or returns defaults with warning
func Load(configPath string) Config {
	config := DefaultConfig()

	// If config file exists, load it
	if configPath != "" {
		if err := loadFromFile(configPath, &config); err != nil {
			slog.Error("Failed to load config from file", "path", configPath, "error", err)
			return config
		}
		return config
	}

	// Try default path if no path specified
	defaultPath := filepath.Join(".", "little-share-conf.json")
	if _, err := os.Stat(defaultPath); err == nil {
		if err := loadFromFile(defaultPath, &config); err != nil {
			slog.Error("Failed to load config from default file", "path", defaultPath, "error", err)
		} else {
			slog.Info("Loaded config from file", "path", defaultPath)
			return config
		}
	}

	// Generate default config file if it doesn't exist
	slog.Warn("No config file found, using defaults and generating config file",
		"config_path", defaultPath)

	if err := saveToFile(defaultPath, config); err != nil {
		slog.Warn("Failed to save default config file", "path", defaultPath, "error", err)
	} else {
		slog.Info("Generated default config file", "path", defaultPath, "please_modify", "Please modify it as needed")
	}

	return config
}

// LoadWithFlag loads configuration from command-line flag or defaults
func LoadWithFlag() Config {
	var configPath string
	flag.StringVar(&configPath, "config", "", "Path to configuration file (default: ./little-share-conf.json)")
	flag.Parse()

	return Load(configPath)
}

func loadFromFile(path string, config *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, config); err != nil {
		return err
	}

	return nil
}

func saveToFile(path string, config Config) error {
	// Create directory if it doesn't exist
	if dir := filepath.Dir(path); dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// GetAddr returns the address to listen on in the format "host:port"
func (c Config) GetAddr() string {
	if c.Port == 0 {
		c.Port = 19999
	}
	return c.Host + ":" + strconv.Itoa(c.Port)
}

func init() {
	// Set up structured logging
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewTextHandler(os.Stdout, opts)
	slog.SetDefault(slog.New(handler))
}
