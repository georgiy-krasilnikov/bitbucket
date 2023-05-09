package main

import (
	"log"
	"os"

	bot "tg_bot/bot"
)

func main() {
	if err := bot.RunBot(); err != nil {
		log.Fatalf("error: %s", err.Error())
		os.Exit(1)
	}
}
