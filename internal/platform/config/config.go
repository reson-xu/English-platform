package config

import (
	"os"
	"time"
)

const (
	defaultAPIAddr        = "127.0.0.1:8080"
	defaultAccessTokenTTL = 24 * time.Hour

	// Development fallback only. Production deployments must set JWT_SECRET.
	defaultJWTSecret = "dev-change-me-english-platform-jwt-secret"
)

type Config struct {
	APIAddr        string
	JWTSecret      string
	AccessTokenTTL time.Duration
}

func Load() Config {
	return Config{
		APIAddr:        getEnv("API_ADDR", defaultAPIAddr),
		JWTSecret:      getEnv("JWT_SECRET", defaultJWTSecret),
		AccessTokenTTL: getDurationEnv("ACCESS_TOKEN_TTL", defaultAccessTokenTTL),
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func getDurationEnv(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}

	return duration
}
