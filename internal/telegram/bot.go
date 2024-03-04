package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tn-bot/internal/models"
	"github.com/tn-bot/internal/storage"
	"google.golang.org/api/sheets/v4"
)

type Bot struct {
	bot         *tgbotapi.BotAPI
	pg          *storage.Pg
	sheetsAdmin *storage.Sheets
	sheetsUsers *storage.Sheets
}

var users = make(map[int64]*models.UserWithCounter, 16)

func NewBot(bot *tgbotapi.BotAPI, pg *storage.Pg, sheetsAdmin *storage.Sheets, sheetsUsers *storage.Sheets) *Bot {
	return &Bot{
		bot:         bot,
		pg:          pg,
		sheetsAdmin: sheetsAdmin,
		sheetsUsers: sheetsUsers,
	}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates := b.initUpdateChannel()

	// fmt.Println("\n\n", users, "\n\n")
	// users[123] = models.NewNilUserWithCounter()
	// fmt.Println("\n\n", users, "\n\n")
	users[341097587] = models.NewNilUserWithCounter()
	users[341097587].User.IsAdmin = true

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil { // If we got a message
			if users[update.Message.Chat.ID] == nil {
				users[update.Message.Chat.ID] = models.NewNilUserWithCounter()
				fmt.Println("\n\n", users, "\n\n")
			}
			if users[update.Message.Chat.ID].Condition == userInQuizCondition {
				b.AnswerTheQuestion(users[update.Message.Chat.ID], update)
				continue
			}
			b.handleMessages(update.Message)
		} else if update.CallbackQuery != nil { // If we got a callback
			if users[update.CallbackQuery.Message.Chat.ID] == nil {
				users[update.CallbackQuery.Message.Chat.ID] = models.NewNilUserWithCounter()
				fmt.Println("\n\n", users, "\n\n")
			}
			if users[update.CallbackQuery.Message.Chat.ID].Condition == userInQuizCondition {
				b.AnswerTheQuestion(users[update.CallbackQuery.Message.Chat.ID], update)
				continue
			}
			b.handleCallbackQuery(update)
		}
	}
}

func (b *Bot) AnswerTheQuestion(user *models.UserWithCounter, update tgbotapi.Update) {
	question := user.User.Questions[user.Counter]

	if update.CallbackQuery != nil {
		switch {
		case user.Counter >= 0 && user.Counter < len(user.User.Questions):
			fmt.Println("\nstep\n")
			// question := u.User.Questions[u.Counter]
			if question.WithAnswers {
				markup := user.AnswerMarkup(user.Counter)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, question.Text)
				msg.ReplyMarkup = &markup
				b.bot.Send(msg)
			} else {

			}
			user.Counter++

		case update.CallbackQuery.Data == "answer1":
			if question.CorrectAnswer == "1" {

			}
		case update.CallbackQuery.Data == "answer2":
			if question.CorrectAnswer == "2" {

			}
		case update.CallbackQuery.Data == "answer3":
			if question.CorrectAnswer == "3" {

			}
		case update.CallbackQuery.Data == "answer4":
			if question.CorrectAnswer == "4" {

			}
		}
		if update.Message != nil {
			switch {
			case user.Counter >= 0 && user.Counter < len(user.User.Questions):
				fmt.Println("\nstep\n")
				// question := u.User.Questions[u.Counter]
				if question.WithAnswers {
					markup := user.AnswerMarkup(user.Counter)

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, question.Text)
					msg.ReplyMarkup = &markup
					b.bot.Send(msg)
				} else {

				}
				user.Counter++
			}
		}
	}
}

func (b *Bot) initUpdateChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) AddQuestionToSheets(question *models.Question) error {
	switch question.Direction {
	case JavaDirection:
		return b.addQuestionToSheetsDirection(JavaDirection)
	case FrontDirection:
		return b.addQuestionToSheetsDirection(FrontDirection)
	case QaDirection:
		return b.addQuestionToSheetsDirection(QaDirection)
	case AnalystDirection:
		return b.addQuestionToSheetsDirection(AnalystDirection)
	}

	return nil
}

func (b *Bot) addQuestionToSheetsDirection(direction string) error {
	var dir string
	switch direction {
	case JavaDirection:
		dir = "Java"
	case FrontDirection:
		dir = "Frontend"
	case QaDirection:
		dir = "QA"
	case AnalystDirection:
		dir = "Analyst"
	default:
		dir = "Unknown"
	}

	// qst := "dsdfsdfs"
	// qsttype := "text"
	// values := [][]interface{}{{1, qst, qsttype}}
	values := [][]interface{}{{b.getCount(dir), question.Text, question.WithAnswers, question.CorrectAnswer,
		question.CountAnswers, question.FirstAnswer, question.SecondAnswer, question.ThirdAnswer, question.FourthAnswer}}

	lengthRange := len(values) + 65
	writeRange := fmt.Sprintf("%s!A:%c", dir, lengthRange)

	valueRange := &sheets.ValueRange{
		Values: values,
	}

	_, err := b.sheetsAdmin.Srv.Spreadsheets.Values.Append(b.sheetsAdmin.SpreadsheetId, writeRange, valueRange).
		ValueInputOption("USER_ENTERED").
		Context(b.sheetsAdmin.Ctx).
		Do()
	if err != nil {
		log.Println("unable to append data to sheet: ", err)
		return err
	}

	question = models.NewNilQuestion()
	return nil
}

func (b *Bot) DeleteQuestionFromSheet(rowIndex int) error {
	switch question.Direction {
	case JavaDirection:
		return b.deleteQuestionFromSheetsDirection(JavaDirection, rowIndex)
	case FrontDirection:
		return b.deleteQuestionFromSheetsDirection(FrontDirection, rowIndex)
	case QaDirection:
		return b.deleteQuestionFromSheetsDirection(QaDirection, rowIndex)
	case AnalystDirection:
		return b.deleteQuestionFromSheetsDirection(AnalystDirection, rowIndex)
	}

	return nil
}

func (b *Bot) deleteQuestionFromSheetsDirection(direction string, rowIndex int) error {
	var dir string
	switch direction {
	case JavaDirection:
		dir = "Java"
	case FrontDirection:
		dir = "Frontend"
	case QaDirection:
		dir = "QA"
	case AnalystDirection:
		dir = "Analyst"
	default:
		dir = "Unknown"
	}

	sheetId := b.getSheetIDByDirection(dir)
	if sheetId == -1 {
		fmt.Printf("\n\nCan't find SheetId: %v\n\n", sheetId)
		return fmt.Errorf("Can't find SheetId: %v", sheetId)
	}

	request := sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				DeleteDimension: &sheets.DeleteDimensionRequest{
					Range: &sheets.DimensionRange{
						SheetId:    sheetId,
						Dimension:  "ROWS",
						StartIndex: int64(rowIndex), // Индекс начинается с 0.
						EndIndex:   int64(rowIndex) + 1,
					},
				},
			},
		},
	}

	// _, err := b.bot.Srv.Spreadsheets.BatchUpdate(sheetId, &request).Context(b.sheets.Ctx).Do()

	_, err := b.sheetsAdmin.Srv.Spreadsheets.BatchUpdate(b.sheetsAdmin.SpreadsheetId, &request).Context(b.sheetsAdmin.Ctx).Do()
	if err != nil {
		log.Println("Unable to delete row: %v", err)
		return err
	}
	return nil
}

func (b *Bot) getSheetIDByDirection(direction string) (sheetID int64) {
	// Здесь вы должны реализовать логику для определения ID листа по направлению.
	// Например, вы можете использовать маппинг направления на ID листа.
	// Ниже приведен пример маппинга:
	directionToSheetID := map[string]int64{
		"Java":     0,
		"Frontend": 1946486958,
		"QA":       1195887942,
		"Analyst":  424018855,
		"Unknown":  -1, // Для неизвестных направлений можете вернуть пустую строку или обработать по-другому.
	}

	return directionToSheetID[direction]
}

func (b *Bot) getCount(direction string) int {
	readRange := direction + "!A:A"

	resp, err := b.sheetsAdmin.Srv.Spreadsheets.Values.Get(b.sheetsAdmin.SpreadsheetId, readRange).Do()
	if err != nil {
		log.Println("unable to retrieve data from sheet: ", err)
	}

	return len(resp.Values)
}

func (b *Bot) getQuestions(direction string) [][]interface{} {
	var dir string
	switch direction {
	case JavaDirection:
		dir = "Java"
	case FrontDirection:
		dir = "Frontend"
	case QaDirection:
		dir = "QA"
	case AnalystDirection:
		dir = "Analyst"
	default:
		dir = "Unknown"
	}

	// questionsCount := b.getCount(dir)

	readRange := dir + "!A:I"

	resp, err := b.sheetsAdmin.Srv.Spreadsheets.Values.Get(b.sheetsAdmin.SpreadsheetId, readRange).Do()
	if err != nil {
		log.Println("unable to retrieve data from sheet: ", err)
	}

	values := resp.Values
	fmt.Println(values)
	return values
}
