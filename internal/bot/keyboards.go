package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) sendPriorityKeyboard(chatID int64) {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Важно", "priority_важно"),
			tgbotapi.NewInlineKeyboardButtonData("Неважно", "priority_неважно"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, "Выберите приоритет:")
	msg.ReplyMarkup = keyboard
	b.api.Send(msg)
}

func (b *Bot) sendUrgencyKeyboard(chatID int64) {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Спешно", "urgency_спешно"),
			tgbotapi.NewInlineKeyboardButtonData("Неспешно", "urgency_неспешно"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, "Выберите срочность:")
	msg.ReplyMarkup = keyboard
	b.api.Send(msg)
}

func (b *Bot) sendStatusKeyboard(chatID int64) {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Выполнено", "status_true"),
			tgbotapi.NewInlineKeyboardButtonData("Не выполнено", "status_false"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, "Статус выполнения:")
	msg.ReplyMarkup = keyboard
	b.api.Send(msg)
}
