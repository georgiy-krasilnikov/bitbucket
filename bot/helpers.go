package tg_bot

import (
	tg "github.com/Syfaro/telegram-bot-api"
)

func NewKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("1️⃣", "1 кнопка:)"),
			tg.NewInlineKeyboardButtonData("2️⃣", "2 кнопка:)"),
			tg.NewInlineKeyboardButtonData("3️⃣", "3 кнопка:)"),
			tg.NewInlineKeyboardButtonData("4️⃣", "4 кнопка:)"),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("5️⃣", "5 кнопка:)"),
			tg.NewInlineKeyboardButtonData("6️⃣", "6 кнопка:)"),
		),
	)
}
