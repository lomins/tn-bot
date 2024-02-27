package telegram

const (
	withAnswers       = "withAnswers"
	withoutAnswers    = "withoutAnswers"
	withAnswersCount2 = "countAnswers2"
	withAnswersCount3 = "countAnswers3"
	withAnswersCount4 = "countAnswers4"
)

const (
	javaDirection    = "Java"
	frontDirection   = "Frontend"
	qaDirection      = "QA"
	analystDirection = "Analyst"
)

type Question struct {
	direction     string
	text          string
	withAnswers   bool
	correctAnswer string
	countAnswers  int

	firstAnswer  string
	secondAnswer string
	thirdAnswer  string
	fourthAnswer string
}
