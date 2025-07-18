package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken    string
	NotionToken      string
	NotionDatabaseID string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	return &Config{
		TelegramToken:    getEnv("TELEGRAM_TOKEN", ""),
		NotionToken:      getEnv("NOTION_TOKEN", ""),
		NotionDatabaseID: getEnv("NOTION_DATABASE_ID", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	if defaultValue == "" {
		log.Fatalf("Environment variable %s is required", key)
	}
	return defaultValue
}
