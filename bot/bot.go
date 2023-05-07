package tg_bot

import (
	"fmt"

	tg "github.com/Syfaro/telegram-bot-api"
)

var botToken = "6066706411:AAGv3h0xsBEayXTuHnGR_q3XSShUJwry4iQ"

func NewBot() (*tg.BotAPI, error) {
	bot, err := tg.NewBotAPI(botToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot API: %s", err.Error())
	}

	return bot, nil
}

func RunBot() error {
	bot, err := NewBot()
	if err != nil {
		return fmt.Errorf("failed to run bot: %s", err.Error())
	}

	upd := tg.NewUpdate(0)
	upd.Timeout = 60

	updates, err := bot.GetUpdatesChan(upd)
	if err != nil {
		return fmt.Errorf("failed to create channel for updates: %w", err)
	}

	for upd := range updates {
		if upd.Message == nil {
			continue
		}

		if upd.Message.Text != "" {

		}

	}

}
