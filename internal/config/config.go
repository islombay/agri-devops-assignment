package config

import "os"

type Config struct {
	Port     string
	Env      string
	LogLevel string
}

func Load() Config {
	return Config{
		Port:     getEnv("PORT", "8080"),
		Env:      getEnv("ENV", "dev"),
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
