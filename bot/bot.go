package tg_bot

import (
	"fmt"
	"log"

	tg "github.com/Syfaro/telegram-bot-api"
)

type TgBot struct {
	bot    *tg.BotAPI
	keyboard  tg.InlineKeyboardMarkup
}

func NewBot(botToken string) (*TgBot, error) {
	bot, err := tg.NewBotAPI(botToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot API: %w", err)
	}

	return &TgBot{
		bot: bot,
		keyboard: tg.NewInlineKeyboardMarkup(
			tg.NewInlineKeyboardRow(
				tg.NewInlineKeyboardButtonData("🎢", "Парк развлечений"),
				tg.NewInlineKeyboardButtonData("🎳", "Боулинг"),
				tg.NewInlineKeyboardButtonData("🎬", "Кинотеатр"),
				tg.NewInlineKeyboardButtonData("🍽", "Кафе"),
			),
		),
		
	}, nil
}

func (bot *TgBot) RunBot() error {

	upd := tg.NewUpdate(0)
	upd.Timeout = 60

	updates, err := bot.bot.GetUpdatesChan(upd)
	if err != nil {
		return fmt.Errorf("failed to create channel for updates: %w", err)
	}

	for upd := range updates {
		if upd.Message == nil && upd.CallbackQuery != nil {
			msg := tg.NewMessage(upd.CallbackQuery.Message.Chat.ID, fmt.Sprintf("Вы выбрали: %s! 😉", upd.CallbackQuery.Data))
			msg.ReplyMarkup = NewKeyboard(upd.CallbackQuery.Data)
			_, err := bot.bot.Send(msg)
			if err != nil {
				log.Fatalf("failed to send 'button' msg: %s", err.Error())
			}

			continue
		}

		if upd.Message != nil {
			switch upd.Message.Text {
			case "/start":
				msg := tg.NewMessage(upd.Message.Chat.ID, "Привет! 😄\n\nЯ телеграм-бот, созданный для тестового задания!\n\nВыбери место, куда бы ты хотел сходить с другом!")
				msg.ReplyMarkup = bot.keyboard
				_, err := bot.bot.Send(msg)
				if err != nil {
					return fmt.Errorf("failed to send 'start' msg: %w", err)
				}

			default:
				msg := tg.NewMessage(upd.Message.Chat.ID, "Я не могу ответить на это сообщение:(")
				_, err := bot.bot.Send(msg)
				if err != nil {
					log.Fatalf("failed to send 'default' msg: %s", err.Error())
				}
			}
		}
	}
	return nil
}
