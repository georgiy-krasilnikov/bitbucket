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
		if upd.Message == nil {
			continue
		}

		if upd.Message.Text != "" {
			switch upd.Message.Text {
			case "/start":
				msg := tg.NewMessage(upd.Message.Chat.ID, "Привет! 😄\nЯ телеграм-бот, созданный для тестового задания!")
				msg.ReplyMarkup = keyboard
				_, err := bot.Send(msg)
				if err != nil {
					return fmt.Errorf("failed to send 'start' msg: %w", err)
				}

			case keyboard.Keyboard[0][0].Text:
				msg := tg.NewMessage(upd.Message.Chat.ID, "Это 1 кнопка:)")
				_, err := bot.Send(msg)
				if err != nil {
					log.Fatalf("failed to send '1 button' msg: %s", err.Error())
				}

			case keyboard.Keyboard[0][1].Text:
				msg := tg.NewMessage(upd.Message.Chat.ID, "Это 2 кнопка:)")
				_, err := bot.Send(msg)
				if err != nil {
					log.Fatalf("failed to send '2 button' msg: %s", err.Error())
				}

			case keyboard.Keyboard[0][2].Text:
				msg := tg.NewMessage(upd.Message.Chat.ID, "Это 3 кнопка:)")
				_, err := bot.Send(msg)
				if err != nil {
					log.Fatalf("failed to send '3 button' msg: %s", err.Error())
				}

			case keyboard.Keyboard[0][3].Text:
				msg := tg.NewMessage(upd.Message.Chat.ID, "Это 4 кнопка:)")
				_, err := bot.Send(msg)
				if err != nil {
					log.Fatalf("failed to send '4 button' msg: %s", err.Error())
				}

			case keyboard.Keyboard[1][0].Text:
				msg := tg.NewMessage(upd.Message.Chat.ID, "Это 5 кнопка:)")
				_, err := bot.Send(msg)
				if err != nil {
					log.Fatalf("failed to send '5 button' msg: %s", err.Error())
				}

			case keyboard.Keyboard[1][1].Text:
				msg := tg.NewMessage(upd.Message.Chat.ID, "Это 6 кнопка:)")
				_, err := bot.Send(msg)
				if err != nil {
					log.Fatalf("failed to send '6 button' msg: %s", err.Error())
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
