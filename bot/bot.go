package tg_bot

import (
	"fmt"
	"log"

	tg "github.com/Syfaro/telegram-bot-api"
)

var botToken = "6066706411:AAGv3h0xsBEayXTuHnGR_q3XSShUJwry4iQ"

func NewBot() (*tg.BotAPI, error) {
	bot, err := tg.NewBotAPI(botToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot API: %w", err)
	}

	return bot, nil
}

func RunBot() error {
	bot, err := NewBot()
	if err != nil {
		return fmt.Errorf("failed to run bot: %w", err)
	}

	upd := tg.NewUpdate(0)
	upd.Timeout = 60

	updates, err := bot.GetUpdatesChan(upd)
	if err != nil {
		return fmt.Errorf("failed to create channel for updates: %w", err)
	}

	keyboard := NewKeyboard()

	for upd := range updates {
		if upd.Message == nil && upd.CallbackQuery != nil {
			msg := tg.NewMessage(upd.CallbackQuery.Message.Chat.ID, upd.CallbackQuery.Data)
			_, err := bot.Send(msg)
			if err != nil {
				log.Fatalf("failed to send 'button' msg: %s", err.Error())
			}
			
			continue
		}

		if upd.Message != nil {
			switch upd.Message.Text {
			case "/start":
				msg := tg.NewMessage(upd.Message.Chat.ID, "Привет! 😄\nЯ телеграм-бот, созданный для тестового задания!\nПопробуй нажать на кнопки:)")
				msg.ReplyMarkup = keyboard
				_, err := bot.Send(msg)
				if err != nil {
					return fmt.Errorf("failed to send 'start' msg: %w", err)
				}

			default:
				msg := tg.NewMessage(upd.Message.Chat.ID, "Я не могу ответить на это сообщение:(")
				_, err := bot.Send(msg)
				if err != nil {
					log.Fatalf("failed to send 'default' msg: %s", err.Error())
				}
			}
		}
	}
	return nil
}
