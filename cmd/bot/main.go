package main

import (
	"TelegramToNotion/internal/bot"
	"TelegramToNotion/internal/config"
	"log"
)

func main() {
	cfg := config.Load()

	notionBot, err := bot.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	log.Println("Bot is running...")
	notionBot.Start()
}
