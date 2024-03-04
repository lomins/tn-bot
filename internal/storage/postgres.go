package storage

import (
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/tn-bot/internal/config"
	"github.com/tn-bot/internal/models"
)

type Pg struct {
	Sqlx *sqlx.DB
	Cfg  config.Config
}

func NewPostgres(cfg config.Config) *Pg {
	conn, err := sqlx.Open("postgres", cfg.PgConn)
	if err != nil {
		log.Fatal("Couldn't open database: ", cfg.PgConn, err)
	}

	pg := Pg{conn, cfg}

	if _, err := pg.Sqlx.Exec(createTablesStr); err != nil {
		log.Fatal("Can't create db: ", pg.Cfg.PgConn, conn, err)
	}
	return &pg
}

func (pg *Pg) Close() error {
	return pg.Sqlx.Close()
}

func (pg *Pg) InsertQuestion(question *models.Question) error {
	// q := "INSERT INTO orders (order_uid, order_data) VALUES(:order_uid, :order_data)"
	// q := `INSERT INTO questionsjava (questionnumber, questiontext, withanswers, correctanswer,
	// 	countanswers, firstanswer, secondanswer, thirdanswer, fourthanswer)
	// 	VALUES
	// 	(:questionNumber, :questionText', :withAnswers, :correctAnswer, :countAnswers,
	// 		:firstAnswer:, :secondAnswer, :thirdAnswer, :fourthAnswer);`
	dir := strings.ToLower(question.Direction)

	q := fmt.Sprintf(`INSERT INTO questions%s (questionnumber, questiontext, withanswers, correctanswer,
				countanswers, firstanswer, secondanswer, thirdanswer, fourthanswer)
				VALUES
				(:questionNumber, :text, :withAnswers, :correctAnswer, :countAnswers, 
					:firstAnswer, :secondAnswer, :thirdAnswer, :fourthAnswer);`, dir)

	_, err := pg.Sqlx.NamedExec(q, question)
	if err != nil {
		if isDuplicateKeyError(err) {
			log.Println("pg.InsertQuestion QuestionNumber already exists: ", question.QuestionNumber)
			return err
		}
		log.Println("pg.InsertOrder unexpected error: ", question.QuestionNumber, err)
		return err
	}
	return nil
}

func isDuplicateKeyError(err error) bool {
	pgErr, ok := err.(*pq.Error)
	return ok && pgErr.Code == "23505"
}

func (pg *Pg) SelectQuestions(direction string) ([]models.Question, error) {
	questions := make([]models.Question, 10)
	dir := strings.ToLower(direction)

	q := fmt.Sprintf("SELECT * FROM questions%s", strings.ToLower(dir))
	rows, err := pg.Sqlx.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var question models.Question
		if err := rows.Scan(&question.QuestionNumber, &question.Text, &question.WithAnswers, &question.CorrectAnswer, &question.CountAnswers, &question.FirstAnswer, &question.SecondAnswer, &question.ThirdAnswer, &question.FourthAnswer); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}

func (pg *Pg) CountQuestions(question *models.Question) {
	q := fmt.Sprintf(`SELECT COUNT(*)+1
	FROM questions%s`, question.Direction)

	row := pg.Sqlx.QueryRow(q)
	err := row.Scan(&question.QuestionNumber)
	if err != nil {
		fmt.Println("Error fetching question count:", err)
	}
}

const createTablesStr = `
CREATE TABLE IF NOT EXISTS questionsJava(
	questionNumber int NOT NULL,
	questionText text NOT NULL,
	withAnswers bool NOT NULL,
	correctAnswer VARCHAR(60) NOT NULL,
	countAnswers int NOT NULL,
	firstAnswer VARCHAR(60) NOT NULL,
	secondAnswer VARCHAR(60),
	thirdAnswer VARCHAR(60),
	fourthAnswer VARCHAR(60));

CREATE TABLE IF NOT EXISTS questionsFrontend(
	questionNumber int NOT NULL,
	questionText text NOT NULL,
	withAnswers bool NOT NULL,
	correctAnswer VARCHAR(60) NOT NULL,
	countAnswers int NOT NULL,
	firstAnswer VARCHAR(60) NOT NULL,
	secondAnswer VARCHAR(60),
	thirdAnswer VARCHAR(60),
	fourthAnswer VARCHAR(60));

CREATE TABLE IF NOT EXISTS questionsQA(
	questionNumber int NOT NULL,
	questionText text NOT NULL,
	withAnswers bool NOT NULL,
	correctAnswer VARCHAR(60) NOT NULL,
	countAnswers int NOT NULL,
	firstAnswer VARCHAR(60) NOT NULL,
	secondAnswer VARCHAR(60),
	thirdAnswer VARCHAR(60),
	fourthAnswer VARCHAR(60));

CREATE TABLE IF NOT EXISTS questionsAnalyst(
	questionNumber int NOT NULL,
	questionText text NOT NULL,
	withAnswers bool NOT NULL,
	correctAnswer VARCHAR(60) NOT NULL,
	countAnswers int NOT NULL,
	firstAnswer VARCHAR(60) NOT NULL,
	secondAnswer VARCHAR(60),
	thirdAnswer VARCHAR(60),
	fourthAnswer VARCHAR(60));`
