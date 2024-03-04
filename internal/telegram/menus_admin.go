package telegram

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func chooseTypeQuestionMenu() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Java-разработчик", JavaDirection),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Frontend-разработчик", FrontDirection),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Тестировщик", QaDirection),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Аналитик", AnalystDirection),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", Back),
		),
	)
}

func addAnswerType() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Текстовый", withoutAnswers),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("С выбором", withAnswers),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", Back),
		),
	)
}

func addCountAnswers() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("2", withAnswersCount2),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("3", withAnswersCount3),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("4", withAnswersCount4),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", Back),
		),
	)
}

func whichAnswerRight() tgbotapi.InlineKeyboardMarkup {

	markup := tgbotapi.NewInlineKeyboardMarkup()

	if question.CountAnswers == 2 {
		return tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.FirstAnswer, firstAnswer),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.SecondAnswer, secondAnswer),
			),
		)
	} else if question.CountAnswers == 3 {
		return tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.FirstAnswer, firstAnswer),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.SecondAnswer, secondAnswer),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.ThirdAnswer, thirdAnswer),
			),
		)
	} else if question.CountAnswers == 4 {
		return tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.FirstAnswer, firstAnswer),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.SecondAnswer, secondAnswer),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.ThirdAnswer, thirdAnswer),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.FourthAnswer, fourthAnswer),
			),
		)
	}

	return markup

	// var rows [][]tgbotapi.InlineKeyboardButton

	// // Создаем отдельную строку клавиатуры для каждого ответа
	// for i, answer := range question.Answers {
	// 	if answer != "" {
	// 		// Создаем кнопку с данными
	// 		btn := tgbotapi.NewInlineKeyboardButtonData(answer, rightAnswerIs+strconv.Itoa(i+1))
	// 		fmt.Println("\n", withAnswersCount+strconv.Itoa(i), "\n")
	// 		// Создаем новую строку клавиатуры и добавляем в нее текущую кнопку
	// 		row := tgbotapi.NewInlineKeyboardRow(btn)
	// 		// Добавляем строку в общий набор строк
	// 		rows = append(rows, row)
	// 	}
	// }

	// countAnswers = len(question.Answers) - 1

	// // Создаем разметку клавиатуры с добавленными строками
	// markup := tgbotapi.NewInlineKeyboardMarkup(rows...)
	// return markup
}

func editQuestion(questions [][]interface{}) tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton

	// Добавляем кнопки по одной в каждую строку
	// row := []tgbotapi.InlineKeyboardButton{}

	for i := 0; i < len(questions); i++ {
		// for i, question := range questions {
		if question != nil {
			if i == 0 {
				continue
			}
			// Создаем кнопку с данными
			btn := tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(i), toEditQuestion+strconv.Itoa(i))
			// Добавляем кнопку в текущую строку
			row := tgbotapi.NewInlineKeyboardRow(btn)
			// 		// Добавляем строку в общий набор строк
			// 		rows = append(rows, row)
			rows = append(rows, row)
		}
	}

	// Добавляем текущую строку в общий набор строк
	// rows = append(rows, row)

	// Создаем разметку клавиатуры с добавленными кнопками
	markup := tgbotapi.NewInlineKeyboardMarkup(rows...)
	return markup
}

func deleteQuestion(questions [][]interface{}) tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton

	// Добавляем кнопки по одной в каждую строку
	// row := []tgbotapi.InlineKeyboardButton{}

	for i := 0; i < len(questions); i++ {
		// for i, question := range questions {
		if question != nil {
			if i == 0 {
				continue
			}
			// Создаем кнопку с данными
			btn := tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(i), toDeleteQuestion+strconv.Itoa(i))
			fmt.Println("\n\n", toDeleteQuestion+strconv.Itoa(i) == toDeleteQuestion1, "\n\n")
			// Добавляем кнопку в текущую строку
			row := tgbotapi.NewInlineKeyboardRow(btn)
			// 		// Добавляем строку в общий набор строк
			// 		rows = append(rows, row)
			rows = append(rows, row)
		}
	}

	backBtn := tgbotapi.NewInlineKeyboardButtonData("Назад", Back)
	row := tgbotapi.NewInlineKeyboardRow(backBtn)
	rows = append(rows, row)
	// Добавляем текущую строку в общий набор строк
	// rows = append(rows, row)

	// Создаем разметку клавиатуры с добавленными кнопками
	markup := tgbotapi.NewInlineKeyboardMarkup(rows...)
	return markup
}
