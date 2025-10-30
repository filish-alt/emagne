package config

import (
	"os"
	"time"
)

type Config struct {
	DBSource          string
	ServerAddress     string
	JWTSecret         string
	JWTExpiration     time.Duration
	PasswordCost      int
}

func Load() *Config {
	return &Config{
		DBSource:          getEnv("DATABASE_URL", "postgresql://postgres:1234@localhost:5432/escrow?sslmode=disable"),
		ServerAddress:     getEnv("SERVER_ADDRESS", ":8080"),
		JWTSecret:         getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		JWTExpiration:     24 * time.Hour,
		PasswordCost:      12,
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
