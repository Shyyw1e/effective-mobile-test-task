package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost		string
	DBPort		string
	DBUser		string
	DBPassword	string
	DBName		string
	Port		string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		slog.Error("failed to load .env", slog.String("error", err.Error()))
	}

	return Config{
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		Port: os.Getenv("PORT"),
	}
}

