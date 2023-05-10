package tg_bot

import (
	tg "github.com/Syfaro/telegram-bot-api"
)

func NewKeyboard(choice string) tg.InlineKeyboardMarkup {
	var name1, name2, url1, url2 string
	switch choice {
	case "Парк развлечений":
		name1, url1, name2, url2 = "Сочи Парк", "https://www.sochipark.ru/", "SkyPark", "https://skypark.ru/"
	case "Боулинг":
		name1, url1, name2, url2 = "Торо", "torobowling-sochi.ru", "Хали-гали", "haligalisochi.ru"
	case "Кинотеатр":
		name1, url1, name2, url2 = "City Stars", "https://kinomonitor.ru/cinemas/city-stars/", "Кинотеатр Сочи", "sochi-sputnik.ru"
	case "Кафе":
		name1, url1, name2, url2 = "Мидийное место", "midiynoe-mesto.ru", "Дельфин и Русалка", "delfinrusalka.ru"
	}

	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonURL(name1, url1),
			tg.NewInlineKeyboardButtonURL(name2, url2),
		),
	)
}
