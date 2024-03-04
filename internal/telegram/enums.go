package telegram

const commandStart = "/start"

const (
	commandStartPanel = "Начать"
)

const (
	standardConditon = iota
	adminCondition
	userCondition
	userInQuizCondition
)

const (
	CommandUserStart = "commandUserStart"
	EditUser         = "editUser"
	EditFIO          = "editFio"
	EditVUZ          = "editVUZ"
	EditGroup        = "editGroup"
)

const (
	BSU   = "BSU"
	BGTU  = "BGTU"
	INDUS = "INDUS"
)

const (
	helloToAdmin = "Привет, админ!"
)

const (
	DefaultAdmin   = "defaultAdmin"
	AddQuestion    = "addQuestion"
	EditQuestion   = "editQuestion"
	DeleteQuestion = "deleteQuestion"
)

const (
	MainMenu         = "startPanel"
	Back             = "back"
	BackUserInput    = "backUserInput"
	CommandStartQuiz = "startQuiz"
	AdminPanel       = "admin"
	AddAnswerType    = "addAnswerType"
	FillQuestion     = "fillQuestion"
	WhichAnswerRight = "WhichAnswerRight"
)

const (
	// rightAnswerIs = "rightAnswerIs"
	firstAnswer  = "firstAnswer"
	secondAnswer = "secondAnswer"
	thirdAnswer  = "thirdAnswer"
	fourthAnswer = "fourthAnswer"
)

const (
	userAnswer  = "userAnswer"
	userAnswer1 = "userAnswer1"
	userAnswer2 = "userAnswer2"
	userAnswer3 = "userAnswer3"
	userAnswer4 = "userAnswer4"
)

const (
	toEditQuestion   = "toEditQuestion"
	toEditQuestion1  = "toEditQuestion1"
	toEditQuestion2  = "toEditQuestion2"
	toEditQuestion3  = "toEditQuestion3"
	toEditQuestion4  = "toEditQuestion4"
	toEditQuestion5  = "toEditQuestion5"
	toEditQuestion6  = "toEditQuestion6"
	toEditQuestion7  = "toEditQuestion7"
	toEditQuestion8  = "toEditQuestion8"
	toEditQuestion9  = "toEditQuestion9"
	toEditQuestion10 = "toEditQuestion10"
)

const (
	toDeleteQuestion   = "toDeleteQuestion"
	toDeleteQuestion1  = "toDeleteQuestion1"
	toDeleteQuestion2  = "toDeleteQuestion2"
	toDeleteQuestion3  = "toDeleteQuestion3"
	toDeleteQuestion4  = "toDeleteQuestion4"
	toDeleteQuestion5  = "toDeleteQuestion5"
	toDeleteQuestion6  = "toDeleteQuestion6"
	toDeleteQuestion7  = "toDeleteQuestion7"
	toDeleteQuestion8  = "toDeleteQuestion8"
	toDeleteQuestion9  = "toDeleteQuestion9"
	toDeleteQuestion10 = "toDeleteQuestion10"
)

var adminPanel = map[string]bool{
	AdminPanel:        true,
	AddQuestion:       true,
	EditQuestion:      true,
	DeleteQuestion:    true,
	JavaDirection:     true,
	FrontDirection:    true,
	QaDirection:       true,
	AnalystDirection:  true,
	withAnswers:       true,
	withAnswersCount2: true,
	withAnswersCount3: true,
	withAnswersCount4: true,
	withoutAnswers:    true,
	firstAnswer:       true,
	secondAnswer:      true,
	thirdAnswer:       true,
	fourthAnswer:      true,

	toEditQuestion1:  true,
	toEditQuestion2:  true,
	toEditQuestion3:  true,
	toEditQuestion4:  true,
	toEditQuestion5:  true,
	toEditQuestion6:  true,
	toEditQuestion7:  true,
	toEditQuestion8:  true,
	toEditQuestion9:  true,
	toEditQuestion10: true,

	// toDeleteQuestion:   true,
	toDeleteQuestion1:  true,
	toDeleteQuestion2:  true,
	toDeleteQuestion3:  true,
	toDeleteQuestion4:  true,
	toDeleteQuestion5:  true,
	toDeleteQuestion6:  true,
	toDeleteQuestion7:  true,
	toDeleteQuestion8:  true,
	toDeleteQuestion9:  true,
	toDeleteQuestion10: true,
}

var userPanel = map[string]bool{
	BackUserInput:        true,
	CommandStartQuiz:     true,
	CommandUserStart:     true,
	BSU:                  true,
	BGTU:                 true,
	INDUS:                true,
	EditUser:             true,
	EditFIO:              true,
	EditVUZ:              true,
	EditGroup:            true,
	JavaDirectionUser:    true,
	FrontDirectionUser:   true,
	QaDirectionUser:      true,
	AnalystDirectionUser: true,
	"answer1":            true,
	"answer2":            true,
	"answer3":            true,
	"answer4":            true,
}

var messagesToAddQuestion = []string{
	"Введите текст вопроса",
	"1-ый ответ",
	"2-ый ответ",
	"3-ый ответ",
	"4-ый ответ",
	"Какой ответ правильный?",
}

var messagesToAddUser = []string{
	"Введите ФИО:\nПример: Иванов Иван Иванович",
	"Выберите ВУЗ:",
	"Введите группу",
}

var toAddIndex = 0
