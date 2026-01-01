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
    FrontendOrigin    string
}

func Load() *Config {
    return &Config{
        DBSource:          getEnv("DATABASE_URL", "postgresql://postgres:1234@localhost:5432/emagne?sslmode=disable"),
        ServerAddress:     getEnv("SERVER_ADDRESS", ":8080"),
        JWTSecret:         getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
        JWTExpiration:     24 * time.Hour,
        PasswordCost:      12,
        FrontendOrigin:    getEnv("FRONTEND_ORIGIN", "http://localhost:3000"),
    }
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
