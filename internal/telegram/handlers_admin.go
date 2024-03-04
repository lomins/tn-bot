package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleAdminCallbackQuery(update tgbotapi.Update) {
	callbackQuery := update.CallbackQuery
	// msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")

	switch callbackQuery.Data {

	case AdminPanel:
		beforeMenu = mainMenu()
		users[update.CallbackQuery.Message.Chat.ID].Condition = adminCondition
		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите действие:")

		markup := adminMenu()
		editedMsg.ReplyMarkup = &markup

		b.bot.Send(editedMsg)
	case AddQuestion:
		beforeMenu = adminMenu()
		users[update.CallbackQuery.Message.Chat.ID].User.AdminMode = AddQuestion
		// adminMode = AddQuestion
		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите направление для добавления вопроса:")

		markup := chooseTypeQuestionMenu()
		editedMsg.ReplyMarkup = &markup

		b.bot.Send(editedMsg)
	case EditQuestion:
		beforeMenu = adminMenu()
		users[update.CallbackQuery.Message.Chat.ID].User.AdminMode = EditQuestion
		// adminMode = EditQuestion

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите направление для редактирования вопроса:")

		markup := chooseTypeQuestionMenu()
		editedMsg.ReplyMarkup = &markup

		b.bot.Send(editedMsg)
	case DeleteQuestion:
		beforeMenu = adminMenu()
		users[update.CallbackQuery.Message.Chat.ID].User.AdminMode = DeleteQuestion
		// adminMode = DeleteQuestion

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите направление для удаления вопроса:")

		markup := chooseTypeQuestionMenu()
		editedMsg.ReplyMarkup = &markup

		b.bot.Send(editedMsg)

	case JavaDirection:
		beforeMenu = adminMenu()
		question.Direction = JavaDirection

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите тип вопроса:")
		adminMode := users[update.CallbackQuery.Message.Chat.ID].User.AdminMode
		if adminMode == AddQuestion {
			markup := addAnswerType()
			editedMsg.ReplyMarkup = &markup
		} else if adminMode == EditQuestion {
			markup := editQuestion(b.getQuestions(question.Direction))
			editedMsg = tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите номер вопроса:")
			editedMsg.ReplyMarkup = &markup
		} else if adminMode == DeleteQuestion {
			if b.getCount(question.Direction) == 0 || b.getCount(question.Direction) == 1 {
				finalMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Вопросов по данному направлению нету :(")
				b.bot.Send(finalMsg)

				adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
				beforeMenu = mainMenu()
				adminMsg.ReplyMarkup = adminMenu()
				b.bot.Send(adminMsg)
			}

			markup := deleteQuestion(b.getQuestions(question.Direction))
			editedMsg = tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите номер вопроса:")
			editedMsg.ReplyMarkup = &markup
		}

		b.bot.Send(editedMsg)
	case FrontDirection:
		beforeMenu = adminMenu()
		question.Direction = FrontDirection

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите тип вопроса:")
		adminMode := users[update.CallbackQuery.Message.Chat.ID].User.AdminMode
		if adminMode == AddQuestion {
			markup := addAnswerType()
			editedMsg.ReplyMarkup = &markup
		} else if adminMode == EditQuestion {
			markup := editQuestion(b.getQuestions(question.Direction))
			editedMsg = tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите номер вопроса:")
			editedMsg.ReplyMarkup = &markup
		} else if adminMode == DeleteQuestion {
			if b.getCount(question.Direction) == 0 || b.getCount(question.Direction) == 1 {
				finalMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Вопросов по данному направлению нету :(")
				b.bot.Send(finalMsg)

				adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
				beforeMenu = mainMenu()
				adminMsg.ReplyMarkup = adminMenu()
				b.bot.Send(adminMsg)
			}

			markup := deleteQuestion(b.getQuestions(question.Direction))
			editedMsg = tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите номер вопроса:")
			editedMsg.ReplyMarkup = &markup
		}

		b.bot.Send(editedMsg)
	case QaDirection:
		beforeMenu = adminMenu()
		question.Direction = QaDirection

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите тип вопроса:")
		adminMode := users[update.CallbackQuery.Message.Chat.ID].User.AdminMode
		if adminMode == AddQuestion {
			markup := addAnswerType()
			editedMsg.ReplyMarkup = &markup
		} else if adminMode == EditQuestion {
			markup := editQuestion(b.getQuestions(question.Direction))
			editedMsg = tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите номер вопроса:")
			editedMsg.ReplyMarkup = &markup
		} else if adminMode == DeleteQuestion {
			if b.getCount(question.Direction) == 0 || b.getCount(question.Direction) == 1 {
				finalMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Вопросов по данному направлению нету :(")
				b.bot.Send(finalMsg)

				adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
				beforeMenu = mainMenu()
				adminMsg.ReplyMarkup = adminMenu()
				b.bot.Send(adminMsg)
			}

			markup := deleteQuestion(b.getQuestions(question.Direction))
			editedMsg = tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите номер вопроса:")
			editedMsg.ReplyMarkup = &markup
		}

		b.bot.Send(editedMsg)
	case AnalystDirection:
		beforeMenu = adminMenu()
		question.Direction = AnalystDirection

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите тип вопроса:")
		adminMode := users[update.CallbackQuery.Message.Chat.ID].User.AdminMode
		if adminMode == AddQuestion {
			markup := addAnswerType()
			editedMsg.ReplyMarkup = &markup
		} else if adminMode == EditQuestion {
			markup := editQuestion(b.getQuestions(question.Direction))
			editedMsg = tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите номер вопроса:")
			editedMsg.ReplyMarkup = &markup
		} else if adminMode == DeleteQuestion {
			if b.getCount(question.Direction) == 0 || b.getCount(question.Direction) == 1 {
				finalMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Вопросов по данному направлению нету :(")
				b.bot.Send(finalMsg)

				adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
				beforeMenu = mainMenu()
				adminMsg.ReplyMarkup = adminMenu()
				b.bot.Send(adminMsg)
			}

			markup := deleteQuestion(b.getQuestions(question.Direction))
			editedMsg = tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите номер вопроса:")
			editedMsg.ReplyMarkup = &markup
		}

		b.bot.Send(editedMsg)

	case withAnswers:
		beforeMenu = addAnswerType()
		question.WithAnswers = true

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите количество вопросов:")

		markup := addCountAnswers()
		editedMsg.ReplyMarkup = &markup

		b.bot.Send(editedMsg)
	case withAnswersCount2:
		beforeMenu = addAnswerType()
		question.CountAnswers = 2

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Пожалуйста, заполните данные\n"+messagesToAddQuestion[toAddIndex])

		b.bot.Send(editedMsg)

		// b.fillQuestion(callbackQuery.Message)
		users[update.CallbackQuery.Message.Chat.ID].Condition = adminCondition
		// condition = adminCondition
	case withAnswersCount3:
		beforeMenu = addAnswerType()
		question.CountAnswers = 3

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Пожалуйста, заполните данные\n"+messagesToAddQuestion[toAddIndex])

		b.bot.Send(editedMsg)

		// b.fillQuestion(callbackQuery.Message)
		users[update.CallbackQuery.Message.Chat.ID].Condition = adminCondition
		// condition = adminCondition
	case withAnswersCount4:
		beforeMenu = addAnswerType()
		question.CountAnswers = 4

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Пожалуйста, заполните данные\n"+messagesToAddQuestion[toAddIndex])

		b.bot.Send(editedMsg)

		// b.fillQuestion(callbackQuery.Message)
		users[update.CallbackQuery.Message.Chat.ID].Condition = adminCondition
		// condition = adminCondition

	case firstAnswer:
		question.CorrectAnswer = "1"
		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Вопрос сохраняется...\nПравильный ответ "+question.CorrectAnswer+" - "+question.FirstAnswer)
		b.bot.Send(editedMsg)
		toAddIndex = 0
		users[update.CallbackQuery.Message.Chat.ID].Condition = standardConditon
		// condition = standardConditon

		// err := b.AddQuestionToSheets(question)
		b.pg.CountQuestions(question)
		err := b.pg.InsertQuestion(question)
		if err != nil {
			log.Println("Не удалось добавить вопрос:", err)
			errorMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Не удалось добавить вопрос")
			b.bot.Send(errorMsg)
		} else {
			finalMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Вопрос успешно сохранён!")
			b.bot.Send(finalMsg)

			adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
			beforeMenu = mainMenu()
			adminMsg.ReplyMarkup = adminMenu()
			b.bot.Send(adminMsg)
		}
	case secondAnswer:
		question.CorrectAnswer = "2"
		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Вопрос сохраняется...\nПравильный ответ "+question.CorrectAnswer+" - "+question.SecondAnswer)
		b.bot.Send(editedMsg)
		toAddIndex = 0
		users[update.CallbackQuery.Message.Chat.ID].Condition = standardConditon

		// err := b.AddQuestionToSheets(question)
		b.pg.CountQuestions(question)
		err := b.pg.InsertQuestion(question)
		if err != nil {
			log.Println("Не удалось добавить вопрос:", err)
			errorMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Не удалось добавить вопрос")
			b.bot.Send(errorMsg)
		} else {
			finalMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Вопрос успешно сохранён!")
			b.bot.Send(finalMsg)

			adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
			beforeMenu = mainMenu()
			adminMsg.ReplyMarkup = adminMenu()
			b.bot.Send(adminMsg)
		}
	case thirdAnswer:
		question.CorrectAnswer = "3"
		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Вопрос сохраняется...\nПравильный ответ "+question.CorrectAnswer+" - "+question.ThirdAnswer)
		b.bot.Send(editedMsg)
		toAddIndex = 0
		users[update.CallbackQuery.Message.Chat.ID].Condition = standardConditon

		// err := b.AddQuestionToSheets(question)
		b.pg.CountQuestions(question)
		err := b.pg.InsertQuestion(question)
		if err != nil {
			log.Println("Не удалось добавить вопрос:", err)
			errorMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Не удалось добавить вопрос")
			b.bot.Send(errorMsg)
		} else {
			finalMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Вопрос успешно сохранён!")
			b.bot.Send(finalMsg)

			adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
			beforeMenu = mainMenu()
			adminMsg.ReplyMarkup = adminMenu()
			b.bot.Send(adminMsg)
		}
	case fourthAnswer:
		question.CorrectAnswer = "4"
		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Вопрос сохраняется...\nПравильный ответ "+question.CorrectAnswer+" - "+question.FourthAnswer)
		b.bot.Send(editedMsg)
		toAddIndex = 0
		users[update.CallbackQuery.Message.Chat.ID].Condition = standardConditon

		// err := b.AddQuestionToSheets(question)
		b.pg.CountQuestions(question)
		err := b.pg.InsertQuestion(question)
		if err != nil {
			log.Println("Не удалось добавить вопрос:", err)
			errorMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Не удалось добавить вопрос")
			b.bot.Send(errorMsg)
		} else {
			finalMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Вопрос успешно сохранён!")
			b.bot.Send(finalMsg)

			adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
			beforeMenu = mainMenu()
			adminMsg.ReplyMarkup = adminMenu()
			b.bot.Send(adminMsg)
		}

	case withoutAnswers:
		beforeMenu = addAnswerType()
		question.WithAnswers = false
		question.CountAnswers = 1

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Пожалуйста, заполните данные\n"+messagesToAddQuestion[toAddIndex])

		b.bot.Send(editedMsg)
		// b.fillQuestion(callbackQuery.Message)
		// toAddIndex++
		// condition = adminCondition
		users[update.CallbackQuery.Message.Chat.ID].Condition = adminCondition

	case toEditQuestion1:
	case toEditQuestion2:
	case toEditQuestion3:
	case toEditQuestion4:
	case toEditQuestion5:
	case toEditQuestion6:
	case toEditQuestion7:
	case toEditQuestion8:
	case toEditQuestion9:
	case toEditQuestion10:

	case toDeleteQuestion1, toDeleteQuestion2, toDeleteQuestion3, toDeleteQuestion4, toDeleteQuestion5,
		toDeleteQuestion6, toDeleteQuestion7, toDeleteQuestion8, toDeleteQuestion9, toDeleteQuestion10:
		b.toDeleteQuestionHandler(update)
	}

	// switch {
	// case strings.Contains(callbackQuery.Data, rightAnswerIs):
	// 	beforeMenu = addAnswerType()

	// 	lastChar := callbackQuery.Data[len(callbackQuery.Data)-1]

	// 	// Преобразуем последний символ в строке в число
	// 	question.CorrectAnswer = string(lastChar)

	// 	editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Вопрос сохраняется...\nПравильный ответ "+question.CorrectAnswer+" - "+question.FirstAnswer)
	// 	b.bot.Send(editedMsg)
	// 	toAddIndex = 0
	// 	condition = 0

	// 	err := b.AddQuestionToSheets(question)
	// 	if err != nil {
	// 		log.Println("Не удалось добавить вопрос:", err)
	// 		errorMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Не удалось добавить вопрос")
	// 		b.bot.Send(errorMsg)
	// 	} else {
	// 		finalMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Вопрос успешно сохранён!")
	// 		b.bot.Send(finalMsg)

	// 		adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
	// 		beforeMenu = mainMenu()
	// 		adminMsg.ReplyMarkup = adminMenu()
	// 		b.bot.Send(adminMsg)
	// 	}
}

func (b *Bot) toDeleteQuestionHandler(update tgbotapi.Update) {
	callbackQuery := update.CallbackQuery
	switch callbackQuery.Data {
	case toDeleteQuestion1:
		err := b.DeleteQuestionFromSheet(1)
		if err != nil {
			log.Println("Не удалось добавить вопрос:", err)
			errorMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Не удалось удалить вопрос")
			b.bot.Send(errorMsg)
		} else {
			finalMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Вопрос успешно удалён!")
			b.bot.Send(finalMsg)

			adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
			beforeMenu = mainMenu()
			adminMsg.ReplyMarkup = adminMenu()
			b.bot.Send(adminMsg)
		}
	case toDeleteQuestion2:
		err := b.DeleteQuestionFromSheet(2)
		if err != nil {
			log.Println("Не удалось добавить вопрос:", err)
			errorMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Не удалось удалить вопрос")
			b.bot.Send(errorMsg)
		} else {
			finalMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Вопрос успешно удалён!")
			b.bot.Send(finalMsg)

			adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
			beforeMenu = mainMenu()
			adminMsg.ReplyMarkup = adminMenu()
			b.bot.Send(adminMsg)
		}
	case toDeleteQuestion3:
		err := b.DeleteQuestionFromSheet(3)
		if err != nil {
			log.Println("Не удалось добавить вопрос:", err)
			errorMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Не удалось удалить вопрос")
			b.bot.Send(errorMsg)
		} else {
			finalMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Вопрос успешно удалён!")
			b.bot.Send(finalMsg)

			adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
			beforeMenu = mainMenu()
			adminMsg.ReplyMarkup = adminMenu()
			b.bot.Send(adminMsg)
		}
	case toDeleteQuestion4:
		err := b.DeleteQuestionFromSheet(4)
		if err != nil {
			log.Println("Не удалось добавить вопрос:", err)
			errorMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Не удалось удалить вопрос")
			b.bot.Send(errorMsg)
		} else {
			finalMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Вопрос успешно удалён!")
			b.bot.Send(finalMsg)

			adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
			beforeMenu = mainMenu()
			adminMsg.ReplyMarkup = adminMenu()
			b.bot.Send(adminMsg)
		}
	case toDeleteQuestion5:
		err := b.DeleteQuestionFromSheet(5)
		if err != nil {
			log.Println("Не удалось добавить вопрос:", err)
			errorMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Не удалось удалить вопрос")
			b.bot.Send(errorMsg)
		} else {
			finalMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Вопрос успешно удалён!")
			b.bot.Send(finalMsg)

			adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
			beforeMenu = mainMenu()
			adminMsg.ReplyMarkup = adminMenu()
			b.bot.Send(adminMsg)
		}
	case toDeleteQuestion6:
		err := b.DeleteQuestionFromSheet(6)
		if err != nil {
			log.Println("Не удалось добавить вопрос:", err)
			errorMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Не удалось удалить вопрос")
			b.bot.Send(errorMsg)
		} else {
			finalMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Вопрос успешно удалён!")
			b.bot.Send(finalMsg)

			adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
			beforeMenu = mainMenu()
			adminMsg.ReplyMarkup = adminMenu()
			b.bot.Send(adminMsg)
		}
	case toDeleteQuestion7:
		err := b.DeleteQuestionFromSheet(7)
		if err != nil {
			log.Println("Не удалось добавить вопрос:", err)
			errorMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Не удалось удалить вопрос")
			b.bot.Send(errorMsg)
		} else {
			finalMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Вопрос успешно удалён!")
			b.bot.Send(finalMsg)

			adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
			beforeMenu = mainMenu()
			adminMsg.ReplyMarkup = adminMenu()
			b.bot.Send(adminMsg)
		}
	case toDeleteQuestion8:
		err := b.DeleteQuestionFromSheet(8)
		if err != nil {
			log.Println("Не удалось добавить вопрос:", err)
			errorMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Не удалось удалить вопрос")
			b.bot.Send(errorMsg)
		} else {
			finalMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Вопрос успешно удалён!")
			b.bot.Send(finalMsg)

			adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
			beforeMenu = mainMenu()
			adminMsg.ReplyMarkup = adminMenu()
			b.bot.Send(adminMsg)
		}
	case toDeleteQuestion9:
		err := b.DeleteQuestionFromSheet(9)
		if err != nil {
			log.Println("Не удалось добавить вопрос:", err)
			errorMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Не удалось удалить вопрос")
			b.bot.Send(errorMsg)
		} else {
			finalMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Вопрос успешно удалён!")
			b.bot.Send(finalMsg)

			adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
			beforeMenu = mainMenu()
			adminMsg.ReplyMarkup = adminMenu()
			b.bot.Send(adminMsg)
		}
	case toDeleteQuestion10:
		err := b.DeleteQuestionFromSheet(10)
		if err != nil {
			log.Println("Не удалось добавить вопрос:", err)
			errorMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Не удалось удалить вопрос")
			b.bot.Send(errorMsg)
		} else {
			finalMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Вопрос успешно удалён!")
			b.bot.Send(finalMsg)

			adminMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, helloToAdmin)
			beforeMenu = mainMenu()
			adminMsg.ReplyMarkup = adminMenu()
			b.bot.Send(adminMsg)
		}
	}
}

func (b *Bot) fillQuestion(message *tgbotapi.Message) {
	switch {
	case toAddIndex == 0:
		question.Text = message.Text
		// question.Answers = make([]string, question.CountAnswers)
		// fmt.Println("\n\nqst.txt", question.Text, "\n\n")
		// fmt.Println("countanswers", question.CountAnswers)
		toAddIndex++

	case toAddIndex == 1:
		fmt.Println(toAddIndex, message.Text)
		question.FirstAnswer = message.Text
		if !question.WithAnswers {
			question.CorrectAnswer = message.Text
			toAddIndex = 0
			// condition = standardConditon

			// err := b.AddQuestionToSheets(question)
			b.pg.CountQuestions(question)
			err := b.pg.InsertQuestion(question)
			if err != nil {
				log.Println("Не удалось добавить вопрос:", err)
				errorMsg := tgbotapi.NewMessage(message.Chat.ID, "Не удалось добавить вопрос")
				b.bot.Send(errorMsg)
			} else {
				finalMsg := tgbotapi.NewMessage(message.Chat.ID, "Вопрос успешно сохранён!")
				b.bot.Send(finalMsg)

				adminMsg := tgbotapi.NewMessage(message.Chat.ID, helloToAdmin)
				beforeMenu = mainMenu()
				adminMsg.ReplyMarkup = adminMenu()
				b.bot.Send(adminMsg)
			}

			// fmt.Println("\n\n", question, "\n\n")
			return
		}
		toAddIndex++
	case toAddIndex == 2:
		question.SecondAnswer = message.Text
		if question.CountAnswers == 2 {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Какой ответ правильный?")

			markup := whichAnswerRight()
			msg.ReplyMarkup = &markup

			b.bot.Send(msg)
			return
		}
		toAddIndex++
	case toAddIndex == 3:
		question.ThirdAnswer = message.Text
		if question.CountAnswers == 3 {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Какой ответ правильный?")

			markup := whichAnswerRight()
			msg.ReplyMarkup = &markup

			b.bot.Send(msg)
			return
		}
		toAddIndex++
	case toAddIndex == 4:
		question.FourthAnswer = message.Text
		if question.CountAnswers == 4 {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Какой ответ правильный?")

			markup := whichAnswerRight()
			msg.ReplyMarkup = &markup

			b.bot.Send(msg)
			return
		}
		toAddIndex++
	}

	msgTosend := tgbotapi.NewMessage(message.Chat.ID, messagesToAddQuestion[toAddIndex])
	b.bot.Send(msgTosend)
}
