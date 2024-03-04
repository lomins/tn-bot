package models

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Question struct {
	QuestionNumber int    `db:"questionNumber"`
	Direction      string `db:"direction"`
	Text           string `db:"text"`
	WithAnswers    bool   `db:"withAnswers"`
	CorrectAnswer  string `db:"correctAnswer"`
	CountAnswers   int    `db:"countAnswers"`

	// Answers []string `json:"answers"`

	FirstAnswer  string `db:"firstAnswer"`
	SecondAnswer string `db:"secondAnswer"`
	ThirdAnswer  string `db:"thirdAnswer"`
	FourthAnswer string `db:"fourthAnswer"`
}

func NewNilQuestion() *Question {
	return &Question{
		QuestionNumber: 0,
		Direction:      "",
		Text:           "",
		WithAnswers:    false,
		CorrectAnswer:  "",
		CountAnswers:   0,
		FirstAnswer:    "",
		SecondAnswer:   "",
		ThirdAnswer:    "",
		FourthAnswer:   "",
	}
}

type User struct {
	Name         string
	Surname      string
	FathersName  string
	IsAdmin      bool
	AdminMode    string
	VUZ          string
	Group        string
	TelegramName string
	EditFlag     bool
	Questions    []Question

	PassedScreening bool
}

func NewNilUser() *User {
	return &User{
		Name:            "",
		Surname:         "",
		FathersName:     "",
		IsAdmin:         false,
		AdminMode:       "",
		VUZ:             "",
		Group:           "",
		TelegramName:    "",
		PassedScreening: false,
		EditFlag:        false,
		Questions:       make([]Question, 10),
	}
}

func NewNilUserWithCounter() *UserWithCounter {
	return &UserWithCounter{
		User:      NewNilUser(),
		Counter:   0,
		Condition: 0,
	}
}

type UserWithCounter struct {
	User      *User
	Counter   int
	Condition int
}

func (u *UserWithCounter) AnswerMarkup(index int) tgbotapi.InlineKeyboardMarkup {
	question := u.User.Questions[index]

	if question.CountAnswers == 1 {
		return tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.FirstAnswer, "answer1"),
			),
		)
	} else if question.CountAnswers == 2 {
		return tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.FirstAnswer, "answer1"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.FirstAnswer, "answer2"),
			),
		)
	} else if question.CountAnswers == 3 {
		return tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.FirstAnswer, "answer1"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.FirstAnswer, "answer2"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.FirstAnswer, "answer3"),
			),
		)
	} else {
		return tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.FirstAnswer, "answer1"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.FirstAnswer, "answer2"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.FirstAnswer, "answer3"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(question.FirstAnswer, "answer4"),
			),
		)
	}
}
