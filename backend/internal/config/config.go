package config
package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
	JWTSecret   string
	Env         string
}

func LoadConfig() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://weaver_user:weaver_password@localhost:5432/weaver_db?sslmode=disable"),
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		Env:         getEnv("ENV", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
