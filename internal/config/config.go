package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func NewConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Print(".env не найден")
	}

	return &Config{
		Port: GetEnv("PORT", "3000"),
	}
}

func GetEnv(key, replacement string) string {
	val := os.Getenv(key)

	if val == "" {
		return replacement
	}

	return val
}
