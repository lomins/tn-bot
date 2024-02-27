package telegram

const commandStart = "/start"

const (
	commandStartPanel = "Начать"
)

const (
	AddQuestion      string = "addQuestion"
	MainMenu                = "startPanel"
	EditQuestion            = "editQuestion"
	DeleteQuestion          = "deleteQuestion"
	Back                    = "back"
	commandStartQuiz        = "startQuiz"
	AdminPanel              = "admin"
	AddAnswerType           = "addAnswerType"
	FillQuestion            = "fillQuestion"
	WhichAnswerRight        = "WhichAnswerRight"
)

const (
	firstAnswer  = "firstAnswer"
	secondAnswer = "secondAnswer"
	thirdAnswer  = "thirdAnswer"
	fourthAnswer = "fourthAnswer"
)

var messagesToAdd = []string{
	"Введите текст вопроса",
	"1-ый ответ",
	"2-ый ответ",
	"3-ый ответ",
	"4-ый ответ",
	"Какой вопрос правильный?",
}
var toAddIndex = 0
