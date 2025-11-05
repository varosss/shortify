package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  string
	DatabaseURL string
	JWTSecret   string
	RedisAddr   string
}

func Load() (*Config, error) {
	_ = godotenv.Load("/srv/.env")

	cfg := &Config{
		ServerPort:  getEnv("SERVER_PORT", "80"),
		DatabaseURL: os.Getenv("POSTGRES_URL"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		RedisAddr:   fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
