package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var beforeMenu = mainMenu()
var question = &Question{}
var condition = 0

func (b *Bot) handleMessages(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	if condition == 1 {
		if toAddIndex < len(messagesToAdd)-1 {
			fmt.Println("\n\nto add index: ", toAddIndex, "\n\n")
			b.fillQuestion(message)
		}
	}

	if condition == 0 {
		switch message.Text {
		case commandStart:
			msg.Text = "Привет, начнём?"
			msg.ReplyMarkup = mainMenu()
			b.bot.Send(msg)
		case AdminPanel:
			msg.Text = "Привет, админ!"
			msg.ReplyMarkup = adminMenu()
			b.bot.Send(msg)
		}
	}

}

func (b *Bot) handleCallbackQuery(update tgbotapi.Update) {
	callbackQuery := update.CallbackQuery
	// msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")

	switch callbackQuery.Data {
	case AdminPanel, AddQuestion, javaDirection, frontDirection, qaDirection, analystDirection,
		withAnswers, withAnswersCount2, withAnswersCount3, withAnswersCount4, withoutAnswers,
		firstAnswer, secondAnswer, thirdAnswer, fourthAnswer:
		b.handleAdminCallbackQuery(update)

	case FillQuestion:
		// b.fillQuestion(update.Message)

	case Back:
		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите действие:")

		markup := beforeMenu
		editedMsg.ReplyMarkup = &markup

		b.bot.Send(editedMsg)
	}
}

func (b *Bot) handleAdminCallbackQuery(update tgbotapi.Update) {
	callbackQuery := update.CallbackQuery
	// msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")

	switch callbackQuery.Data {

	case AdminPanel:
		beforeMenu = mainMenu()
		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите действие:")

		markup := adminMenu()
		editedMsg.ReplyMarkup = &markup

		b.bot.Send(editedMsg)
	case AddQuestion:
		beforeMenu = adminMenu()
		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите направление для добавления вопроса:")

		markup := addQuestionMenu()
		editedMsg.ReplyMarkup = &markup

		b.bot.Send(editedMsg)

	case javaDirection:
		beforeMenu = addQuestionMenu()
		question.direction = javaDirection

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите тип вопроса:")

		markup := addAnswerType()
		editedMsg.ReplyMarkup = &markup

		b.bot.Send(editedMsg)
	case frontDirection:
		beforeMenu = addQuestionMenu()
		question.direction = frontDirection

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите тип вопроса:")

		markup := addAnswerType()
		editedMsg.ReplyMarkup = &markup

		b.bot.Send(editedMsg)
	case qaDirection:
		beforeMenu = addQuestionMenu()
		question.direction = qaDirection

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите тип вопроса:")

		markup := addAnswerType()
		editedMsg.ReplyMarkup = &markup

		b.bot.Send(editedMsg)
	case analystDirection:
		beforeMenu = addQuestionMenu()
		question.direction = analystDirection

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите тип вопроса:")

		markup := addAnswerType()
		editedMsg.ReplyMarkup = &markup

		b.bot.Send(editedMsg)

	case withAnswers:
		beforeMenu = addAnswerType()
		question.withAnswers = true

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите количество вопросов:")

		markup := addCountAnswers()
		editedMsg.ReplyMarkup = &markup

		b.bot.Send(editedMsg)
	case withAnswersCount2:
		beforeMenu = addAnswerType()
		question.countAnswers = 2

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Пожалуйста, заполните данные\n"+messagesToAdd[toAddIndex])

		b.bot.Send(editedMsg)

		// b.fillQuestion(callbackQuery.Message)
		condition = 1
	case withAnswersCount3:
		beforeMenu = addAnswerType()
		question.countAnswers = 3

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Пожалуйста, заполните данные\n"+messagesToAdd[toAddIndex])

		b.bot.Send(editedMsg)

		// b.fillQuestion(callbackQuery.Message)
		condition = 1
	case withAnswersCount4:
		beforeMenu = addAnswerType()
		question.countAnswers = 4

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Пожалуйста, заполните данные\n"+messagesToAdd[toAddIndex])

		b.bot.Send(editedMsg)

		// b.fillQuestion(callbackQuery.Message)
		condition = 1
	// case whichAnswerRight:

	case firstAnswer:
		question.correctAnswer = "1"
		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Вопрос сохранён! Правильный ответ "+question.correctAnswer+" - "+question.firstAnswer)
		b.bot.Send(editedMsg)
		toAddIndex = 0
		condition = 0
	case secondAnswer:
		question.correctAnswer = "2"
		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Вопрос сохранён! Правильный ответ "+question.correctAnswer+" - "+question.secondAnswer)
		b.bot.Send(editedMsg)
		toAddIndex = 0
		condition = 0
	case thirdAnswer:
		question.correctAnswer = "3"
		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Вопрос сохранён! Правильный ответ "+question.correctAnswer+" - "+question.thirdAnswer)
		b.bot.Send(editedMsg)
		toAddIndex = 0
		condition = 0
	case fourthAnswer:
		question.correctAnswer = "4"
		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Вопрос сохранён! Правильный ответ "+question.correctAnswer+" - "+question.fourthAnswer)
		b.bot.Send(editedMsg)
		toAddIndex = 0
		condition = 0

	case withoutAnswers:
		beforeMenu = addAnswerType()
		question.withAnswers = false
		question.countAnswers = 1

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Пожалуйста, заполните данные\n"+messagesToAdd[toAddIndex])

		b.bot.Send(editedMsg)
		// b.fillQuestion(callbackQuery.Message)
		// toAddIndex++
		condition = 1
	}
}

func (b *Bot) fillQuestion(message *tgbotapi.Message) {
	switch toAddIndex {
	case 0:
		question.text = message.Text
		fmt.Println("\n\nqst.txt", question.text, "\n\n")
		toAddIndex++
		msgTosend := tgbotapi.NewMessage(message.Chat.ID, messagesToAdd[toAddIndex])
		b.bot.Send(msgTosend)
		return
	case 1:
		fmt.Println(toAddIndex, message.Text)
		question.firstAnswer = message.Text
		if !question.withAnswers {
			question.correctAnswer = message.Text
			toAddIndex = 0
			condition = 0

			finalMsg := tgbotapi.NewMessage(message.Chat.ID, "Вопрос успешно сохранён!")
			b.bot.Send(finalMsg)

			fmt.Println("\n\n", question, "\n\n")
			return
		}
		toAddIndex++
	case 2:
		question.secondAnswer = message.Text
		if question.countAnswers == 2 {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Какой ответ правильный?")

			markup := whichAnswerRight()
			msg.ReplyMarkup = &markup

			b.bot.Send(msg)
			return
		}
		toAddIndex++
	case 3:
		question.thirdAnswer = message.Text
		if question.countAnswers == 3 {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Какой ответ правильный?")

			markup := whichAnswerRight()
			msg.ReplyMarkup = &markup

			b.bot.Send(msg)
			return
		}
		toAddIndex++
	case 4:
		question.fourthAnswer = message.Text
		if question.countAnswers == 4 {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Какой ответ правильный?")

			markup := whichAnswerRight()
			msg.ReplyMarkup = &markup

			b.bot.Send(msg)
			return
		}
		toAddIndex++
		// case 5:
		// 	num, err := strconv.Atoi(message.Text)
		// 	if err != nil {
		// 		msgTosend := tgbotapi.NewMessage(message.Chat.ID, "Вы ввели не число\nпопробуйте ещё раз")
		// 		b.bot.Send(msgTosend)
		// 		return
		// 	}
		// 	if num > question.countAnswers {
		// 		msgTosend := tgbotapi.NewMessage(message.Chat.ID, "Номер ответа больше, чем есть всего!\nпопробуйте ещё раз")
		// 		b.bot.Send(msgTosend)
		// 		return
		// 	}
		// 	question.correctAnswer = message.Text
		// 	// toAddIndex++
		// 	finalMsg := tgbotapi.NewMessage(message.Chat.ID, "Вопрос успешно сохранён!")
		// 	b.bot.Send(finalMsg)
		// 	toAddIndex = 0
		// 	condition = 0
		// 	// case 6:
		// 	// 	fmt.Println(question)
		// 	// 	finalMsg := tgbotapi.NewMessage(message.Chat.ID, "Вопрос успешно сохранён!")
		// 	// 	b.bot.Send(finalMsg)
		// 	// 	toAddIndex = 0
		// 	// 	condition = 0
	}
	msgTosend := tgbotapi.NewMessage(message.Chat.ID, messagesToAdd[toAddIndex])
	b.bot.Send(msgTosend)
}
