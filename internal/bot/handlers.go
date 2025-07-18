package bot

import (
	"TelegramToNotion/internal/notion"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleMessage(msg *tgbotapi.Message) {
	chatID := msg.Chat.ID

	switch {
	case msg.IsCommand() && msg.Command() == "start":
		b.sendWelcomeMessage(chatID)

	case msg.IsCommand() && msg.Command() == "new":
		b.startPageCreation(chatID)

	default:
		b.processUserInput(msg)
	}
}

func (b *Bot) handleCallback(callback *tgbotapi.CallbackQuery) {
	chatID := callback.Message.Chat.ID
	data := callback.Data

	switch {
	case strings.HasPrefix(data, "priority_"):
		b.state.SetPriority(chatID, strings.TrimPrefix(data, "priority_"))
		b.sendUrgencyKeyboard(chatID)

	case strings.HasPrefix(data, "urgency_"):
		b.state.SetUrgency(chatID, strings.TrimPrefix(data, "urgency_"))
		b.sendDateRequest(chatID)

	case strings.HasPrefix(data, "status_"):
		status, _ := strconv.ParseBool(strings.TrimPrefix(data, "status_"))
		b.state.SetStatus(chatID, status)
		b.completePageCreation(chatID)
	}

	b.api.Send(tgbotapi.NewCallback(callback.ID, ""))
}

func (b *Bot) sendWelcomeMessage(chatID int64) {
	text := "Привет! Я бот для работы с Notion.\n" +
		"Используй команды:\n" +
		"/new - создать новую страницу\n" +
		"/help - помощь"

	msg := tgbotapi.NewMessage(chatID, text)
	b.api.Send(msg)
}

func (b *Bot) startPageCreation(chatID int64) {
	b.state.StartCreation(chatID)
	msg := tgbotapi.NewMessage(chatID, "Введите заголовок страницы:")
	b.api.Send(msg)
}

func (b *Bot) processUserInput(msg *tgbotapi.Message) {
	chatID := msg.Chat.ID
	state := b.state.Get(chatID)

	if state == nil {
		return
	}

	switch state.Step {
	case "title":
		b.state.SetTitle(chatID, msg.Text)
		b.sendPriorityKeyboard(chatID)

	case "due_date":
		if b.state.SetDueDate(chatID, msg.Text) {
			b.sendStatusKeyboard(chatID)
		} else {
			msg := tgbotapi.NewMessage(chatID, "❌ Неверный формат даты. Используйте ГГГГ-ММ-ДД:")
			b.api.Send(msg)
		}
	}
}

func (b *Bot) completePageCreation(chatID int64) {
	state := b.state.Get(chatID)
	if state == nil {
		return
	}

	err := notion.CreatePage(b.config, state)
	if err != nil {
		log.Printf("Notion API error: %v", err)
		msg := tgbotapi.NewMessage(chatID, "❌ Ошибка при создании страницы в Notion")
		b.api.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(chatID, "✅ Страница успешно создана в Notion!")
		b.api.Send(msg)
	}

	b.state.Complete(chatID)
}

func (b *Bot) sendDateRequest(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "📅 Введите срок выполнения (в формате ГГГГ-ММ-ДД):")
	b.api.Send(msg)
}
