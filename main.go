package main

import (
	"log"
	"os"

	bot "tg_bot/bot"
)

func main() {
	botToken := os.Getenv("BOT_TOKEN")

	bot, err := bot.NewBot(botToken)
	if err != nil {
		log.Fatalf("failed to start bot: %s", err.Error())
	}

	if err := bot.RunBot(); err != nil {
		log.Fatalf("error: %s", err.Error())
		os.Exit(1)
	}
}
