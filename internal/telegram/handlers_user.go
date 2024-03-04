package telegram

import (
	"fmt"
	"log"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleUserCallbackQuery(update tgbotapi.Update) {
	fmt.Println(users)
	callbackQuery := update.CallbackQuery

	// fmt.Printf("\n\nusers: %v, %v\n\n", users, update.CallbackQuery.Message.Chat.ID)

	switch callbackQuery.Data {
	case CommandUserStart:
		users[update.CallbackQuery.Message.Chat.ID].Condition = userCondition

		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Пожалуйста, заполните данные\n"+messagesToAddUser[users[update.CallbackQuery.Message.Chat.ID].Counter])

		b.bot.Send(editedMsg)

	case BackUserInput:
		users[update.CallbackQuery.Message.Chat.ID].Counter--
		// users[update.Message.Chat.ID].Counter--

	case BSU:
		users[update.CallbackQuery.Message.Chat.ID].User.VUZ = BSU
		users[update.CallbackQuery.Message.Chat.ID].Counter++

		if users[update.CallbackQuery.Message.Chat.ID].User.EditFlag {
			user := users[update.CallbackQuery.Message.Chat.ID].User
			str := fmt.Sprintf("Все данные правильные?\nФамилия: %s\nИмя: %s\nОтчество: %s\nВУЗ: %s\nГруппа: %s",
				user.Surname, user.Name, user.FathersName, user.VUZ, user.Group)
			msgTosend := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, str)
			markup := userStartMenu()
			msgTosend.ReplyMarkup = &markup

			b.bot.Send(msgTosend)
			user.EditFlag = false
			return
		}
		msgTosend := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, messagesToAddUser[users[update.CallbackQuery.Message.Chat.ID].Counter])

		b.bot.Send(msgTosend)
	case BGTU:
		users[update.CallbackQuery.Message.Chat.ID].User.VUZ = BGTU
		users[update.CallbackQuery.Message.Chat.ID].Counter++
		if users[update.CallbackQuery.Message.Chat.ID].User.EditFlag {
			user := users[update.CallbackQuery.Message.Chat.ID].User
			str := fmt.Sprintf("Все данные правильные?\nФамилия: %s\nИмя: %s\nОтчество: %s\nВУЗ: %s\nГруппа: %s",
				user.Surname, user.Name, user.FathersName, user.VUZ, user.Group)
			msgTosend := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, str)
			markup := userStartMenu()
			msgTosend.ReplyMarkup = &markup

			b.bot.Send(msgTosend)
			user.EditFlag = false
			return
		}
		msgTosend := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, messagesToAddUser[users[update.CallbackQuery.Message.Chat.ID].Counter])

		b.bot.Send(msgTosend)
	case INDUS:
		users[update.CallbackQuery.Message.Chat.ID].User.VUZ = INDUS
		users[update.CallbackQuery.Message.Chat.ID].Counter++
		if users[update.CallbackQuery.Message.Chat.ID].User.EditFlag {
			user := users[update.CallbackQuery.Message.Chat.ID].User
			str := fmt.Sprintf("Все данные правильные?\nФамилия: %s\nИмя: %s\nОтчество: %s\nВУЗ: %s\nГруппа: %s",
				user.Surname, user.Name, user.FathersName, user.VUZ, user.Group)
			msgTosend := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, str)
			markup := userStartMenu()
			msgTosend.ReplyMarkup = &markup

			b.bot.Send(msgTosend)
			user.EditFlag = false
			return
		}
		msgTosend := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, messagesToAddUser[users[update.CallbackQuery.Message.Chat.ID].Counter])

		b.bot.Send(msgTosend)

	case JavaDirectionUser:
		dir := parseDirection(JavaDirectionUser)
		questions, err := b.pg.SelectQuestions(dir)
		if err != nil {
			log.Println("\nerror select from pg, err = %s\n", err)
			// Обработка ошибки
		}

		users[update.CallbackQuery.Message.Chat.ID].Condition = userInQuizCondition
		users[update.CallbackQuery.Message.Chat.ID].User.Questions = questions
		// user := users[update.CallbackQuery.Message.Chat.ID]
		str := fmt.Sprintf("Тогда начнём?")

		msgToSend := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, str)

		b.bot.Send(msgToSend)

	case FrontDirectionUser:
		dir := parseDirection(FrontDirectionUser)
		questions, err := b.pg.SelectQuestions(dir)
		if err != nil {
			log.Println("\nerror select from pg, err = %s\n", err)
			// Обработка ошибки
		}

		users[update.CallbackQuery.Message.Chat.ID].Condition = userInQuizCondition
		users[update.CallbackQuery.Message.Chat.ID].User.Questions = questions
		str := fmt.Sprintf("Тогда начнём?")

		msgToSend := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, str)

		b.bot.Send(msgToSend)

	case QaDirectionUser:
		dir := parseDirection(QaDirectionUser)
		questions, err := b.pg.SelectQuestions(dir)
		if err != nil {
			log.Println("\nerror select from pg, err = %s\n", err)
			// Обработка ошибки
		}

		users[update.CallbackQuery.Message.Chat.ID].Condition = userInQuizCondition
		users[update.CallbackQuery.Message.Chat.ID].User.Questions = questions
		str := fmt.Sprintf("Тогда начнём?")

		msgToSend := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, str)

		b.bot.Send(msgToSend)
	case AnalystDirectionUser:
		dir := parseDirection(AnalystDirection)
		questions, err := b.pg.SelectQuestions(dir)
		if err != nil {
			log.Println("\nerror select from pg, err = %s\n", err)
			// Обработка ошибки
		}

		users[update.CallbackQuery.Message.Chat.ID].Condition = userInQuizCondition
		users[update.CallbackQuery.Message.Chat.ID].User.Questions = questions

	case EditUser:
		users[update.CallbackQuery.Message.Chat.ID].User.EditFlag = true
		editedMsg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Что хотите исправить?")

		markup := userEditMenu()
		editedMsg.ReplyMarkup = &markup

		b.bot.Send(editedMsg)

	case EditFIO:
		users[update.CallbackQuery.Message.Chat.ID].Condition = userCondition
		users[update.CallbackQuery.Message.Chat.ID].Counter = 0

		msgToSend := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Пожалуйста, заполните данные\n"+messagesToAddUser[users[update.CallbackQuery.Message.Chat.ID].Counter])
		b.bot.Send(msgToSend)
	case EditVUZ:
		users[update.CallbackQuery.Message.Chat.ID].Condition = userCondition
		users[update.CallbackQuery.Message.Chat.ID].Counter = 1

		msgToSend := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Пожалуйста, заполните данные\n"+messagesToAddUser[users[update.CallbackQuery.Message.Chat.ID].Counter])
		markup := chooseVUZ()
		msgToSend.ReplyMarkup = &markup
		b.bot.Send(msgToSend)
	case EditGroup:
		users[update.CallbackQuery.Message.Chat.ID].Condition = userCondition
		users[update.CallbackQuery.Message.Chat.ID].Counter = 2

		msgToSend := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Пожалуйста, заполните данные\n"+messagesToAddUser[users[update.CallbackQuery.Message.Chat.ID].Counter])
		b.bot.Send(msgToSend)

	case CommandStartQuiz:
		beforeMenu = mainMenu()
		editedMsg := tgbotapi.NewEditMessageText(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, "Выберите направление:")

		markup := chooseQuizDirection()
		editedMsg.ReplyMarkup = &markup

		b.bot.Send(editedMsg)
	}
}

func (b *Bot) fillUser(message *tgbotapi.Message) {
	// counter := users[message.Chat.ID].Counter
	switch {
	case users[message.Chat.ID].Counter == 0:
		fmt.Printf("\n\nChatID = %d\n\n", message.Chat.ID)
		// users[message.Chat.ID] = models.NewNilUserWithCounter()
		surname, name, fathersName := parseFullName(message.Text)
		users[message.Chat.ID].User.Name = name
		users[message.Chat.ID].User.Surname = surname
		users[message.Chat.ID].User.FathersName = fathersName

		users[message.Chat.ID].Counter++

		if users[message.Chat.ID].User.EditFlag {
			user := users[message.Chat.ID].User
			str := fmt.Sprintf("Все данные правильные?\nФамилия: %s\nИмя: %s\nОтчество: %s\nВУЗ: %s\nГруппа: %s",
				user.Surname, user.Name, user.FathersName, user.VUZ, user.Group)
			msgTosend := tgbotapi.NewMessage(message.Chat.ID, str)
			markup := userStartMenu()
			msgTosend.ReplyMarkup = &markup

			b.bot.Send(msgTosend)
			user.EditFlag = false
			return
		}

		msgTosend := tgbotapi.NewMessage(message.Chat.ID, messagesToAddUser[users[message.Chat.ID].Counter])
		markup := chooseVUZ()
		msgTosend.ReplyMarkup = &markup

		b.bot.Send(msgTosend)
		fmt.Printf("\n\nUserWithC = %v\nUser = \n", users[message.Chat.ID], users[message.Chat.ID].User)
		return

		// markup := BackUserInput
		// msgTosend.ReplyMarkup = markup

	case users[message.Chat.ID].Counter == 1:
		fmt.Printf("\n\nChatID = %d\n\n", message.Chat.ID)

		// msgTosend := tgbotapi.NewMessage(message.Chat.ID, messagesToAddUser[users[message.Chat.ID].Counter])

		users[message.Chat.ID].Counter++
		msgTosend := tgbotapi.NewMessage(message.Chat.ID, messagesToAddUser[users[message.Chat.ID].Counter])
		markup := chooseVUZ()
		msgTosend.ReplyMarkup = &markup

		b.bot.Send(msgTosend)
		//To EDIT go to BSU or BGTU
		return

	case users[message.Chat.ID].Counter == 2:
		users[message.Chat.ID].User.Group = message.Text

		user := users[message.Chat.ID].User

		users[message.Chat.ID].Counter++
		str := fmt.Sprintf("Все данные правильные?\nФамилия: %s\nИмя: %s\nОтчество: %s\nВУЗ: %s\nГруппа: %s",
			user.Surname, user.Name, user.FathersName, user.VUZ, user.Group)
		msgTosend := tgbotapi.NewMessage(message.Chat.ID, str)
		markup := userStartMenu()
		msgTosend.ReplyMarkup = &markup

		b.bot.Send(msgTosend)
		user.EditFlag = false
		return
	}

	// msgTosend := tgbotapi.NewMessage(message.Chat.ID, messagesToAddUser[users[message.Chat.ID].Counter])
	msgTosend := tgbotapi.NewMessage(message.Chat.ID, messagesToAddUser[users[message.Chat.ID].Counter])
	b.bot.Send(msgTosend)
}

func parseFullName(fullName string) (string, string, string) {
	// Задаем регулярное выражение для поиска Фамилии, Имени и Отчества
	re := regexp.MustCompile(`^([A-Za-zА-Яа-яЁё]+) ([A-Za-zА-Яа-яЁё]+) ([A-Za-zА-Яа-яЁё]+)$`)

	// Используем FindStringSubmatch для поиска совпадений
	matches := re.FindStringSubmatch(fullName)

	if len(matches) < 4 {
		return "", "", ""
	}

	// Первый элемент - полное совпадение, остальные элементы - группы
	surname := matches[1]
	name := matches[2]
	fathersName := matches[3]

	return surname, name, fathersName
}

func parseDirection(direction string) string {
	switch direction {
	case JavaDirectionUser:
		return "java"
	case FrontDirectionUser:
		return "frontend"
	case QaDirectionUser:
		return "qa"
	case AnalystDirectionUser:
		return "analyst"
	}
	return "java"
}
