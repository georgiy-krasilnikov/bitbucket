package tg_bot

import (
	tg "github.com/Syfaro/telegram-bot-api"
)

func NewKeyboard() tg.ReplyKeyboardMarkup {
	return tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton("1️⃣"),
			tg.NewKeyboardButton("2️⃣"),
			tg.NewKeyboardButton("3️⃣"),
			tg.NewKeyboardButton("4️⃣"),
		),
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton("5️⃣"),
			tg.NewKeyboardButton("6️⃣"),
		),
	)
}
