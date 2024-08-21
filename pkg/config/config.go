package config

import "os"

type Config struct {
	ServerPort string
}

func NewConfig() *Config {
	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
