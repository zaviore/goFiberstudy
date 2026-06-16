package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type AppConfig struct {
	Name        string `json:"name"`
	Port        string `json:"port"`
	Environment string `json:"environment"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	SSLMode  string `json:"sslmode"`
}

type JWTConfig struct {
	Secret string `json:"secret"`
	Expiry string `json:"expiry"`
}

type Config struct {
	App      AppConfig      `json:"app"`
	Database DatabaseConfig `json:"database"`
	JWT      JWTConfig      `json:"jwt"`
}

var AppCfg *Config

func LoadConfig(environment string) (*Config, error) {
	configPath := filepath.Join("config", fmt.Sprintf("config.%s.json", environment))

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found: %s", configPath)
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	AppCfg = &cfg
	return &cfg, nil
}

func GetDSN(dbCfg DatabaseConfig) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Name,
		dbCfg.SSLMode,
	)
}
