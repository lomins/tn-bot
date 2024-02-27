package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{
		bot: bot,
		// answersChannel: make(chan tgbotapi.Update, 16),
	}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates := b.initUpdateChannel()

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil { // If we got a message
			// b.handleUpdate(update)
			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Начнём?")
			// msg.ReplyMarkup = mainMenu()
			// b.bot.Send(msg)
			b.handleMessages(update.Message)
		} else if update.CallbackQuery != nil { // If we got a message
			b.handleCallbackQuery(update)
		}
	}
}

func (b *Bot) handleUpdate(update tgbotapi.Update) {
	log.Printf("[%s] %s chID %d", update.Message.From.UserName, update.Message.Text, update.Message.Chat.ID)

	b.handleMessages(update.Message)

	// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	// msg.ReplyToMessageID = update.Message.	MessageID
}

func (b *Bot) initUpdateChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
