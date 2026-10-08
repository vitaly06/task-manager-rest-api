package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DSN       string
	JwtSecret string
}

func NewConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Print(".env не найден")
	}

	return &Config{
		Port:      GetEnv("PORT", "3000"),
		DSN:       GetEnv("DSN", "host=localhost user=postgres password=your_password dbname=name_db port=5432"),
		JwtSecret: GetEnv("JWT_SECRET", "XXXXXXXXXX"),
	}
}

func GetEnv(key, replacement string) string {
	val := os.Getenv(key)

	if val == "" {
		return replacement
	}

	return val
}
