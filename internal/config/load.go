package config

import (
	"fmt"
	"os"
)

type Config struct {
	JwtSignKey string
}

var AppConfig Config

func LoadConfig() error {
	config := Config{
		JwtSignKey: os.Getenv("JWT_SIGN_KEY"),
	}

	if err := validateConfig(config); err != nil {
		return fmt.Errorf("config validation failed: %v", err)
	}

	AppConfig = config
	return nil
}

func validateConfig(config Config) error {
	if config.JwtSignKey == "" {
		return fmt.Errorf("JWT_SIGN_KEY is required")
	}
	return nil
}

func GetConfig() Config {
	return AppConfig
}

func SetConfig(config Config) {
	AppConfig = config
}
