package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tn-bot/internal/models"
)

var beforeMenu = mainMenu()
var question = models.NewNilQuestion()

// var condition = standardConditon
// var adminMode = ""

func (b *Bot) handleMessages(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	if users[message.Chat.ID] == nil {
		users[message.Chat.ID] = models.NewNilUserWithCounter()
		fmt.Println("\n\n", users, "\n\n")
	}

	if users[message.Chat.ID].Condition == userCondition {
		if users[message.Chat.ID].Counter <= len(messagesToAddUser) {
			b.fillUser(message)
		}
	} else if users[message.Chat.ID].Condition == standardConditon {
		switch message.Text {
		case commandStart:
			msg.Text = "Привет, начнём?"
			msg.ReplyMarkup = mainMenu()
			b.bot.Send(msg)
		case AdminPanel:
			beforeMenu = mainMenu()
			msg.Text = helloToAdmin
			msg.ReplyMarkup = adminMenu()
			b.bot.Send(msg)
		}
	} else if users[message.Chat.ID].Condition == adminCondition && users[message.Chat.ID].User.IsAdmin {
		if toAddIndex <= question.CountAnswers {
			fmt.Println("\nto add index: ", toAddIndex, "\nCountAnswers: ", countAnswers)
			b.fillQuestion(message)
		}
	}
}

func (b *Bot) handleCallbackQuery(update tgbotapi.Update) {

	if users[update.CallbackQuery.From.ID] == nil {
		users[update.CallbackQuery.From.ID] = models.NewNilUserWithCounter()
		fmt.Println("\n\n", users, "\n\n")
	}

	callbackQuery := update.CallbackQuery
	// msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")

	// switch callbackQuery.Data {
	switch {

	case callbackQuery.Data == commandStart:
		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите действие:")
		markup := mainMenu()
		editedMsg.ReplyMarkup = &markup
		b.bot.Send(editedMsg)

	case userPanel[callbackQuery.Data]:
		b.handleUserCallbackQuery(update)

	case adminPanel[callbackQuery.Data]:
		b.handleAdminCallbackQuery(update)

	case callbackQuery.Data == Back:
		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите действие:")

		markup := beforeMenu
		editedMsg.ReplyMarkup = &markup

		b.bot.Send(editedMsg)

		// case callbackQuery.Data == FillQuestion:
		// b.fillQuestion(update.Message)
	}
}
