package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	ServerPort      string
	LogFilePath     string
	ShutdownTimeout time.Duration
}

func Load() Config {
	return Config{
		ServerPort:      getEnv("SERVER_PORT", "8080"),
		LogFilePath:     getEnv("LOG_FILE", "logs.txt"),
		ShutdownTimeout: getEnvAsDuration("SHUTDOWN_TIMEOUT", 5*time.Second),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if valueStr := os.Getenv(key); valueStr != "" {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return time.Duration(value) * time.Second
		}
	}
	return defaultValue
}
