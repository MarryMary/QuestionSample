package Model

import (
	"triela/Structs"
)

func Q_Push(ExName string, ExType string, MainTitle string, Tag string, MainText string, QType string, Answer string, Choice string, Score string) bool {
	db := ConnectDB()
	if QType == "select" {
		stmt, err := db.Prepare("INSERT INTO Question(ExName, ExType, QName, QTags, QMain, AnswerArea, QAnswer, QChoice, QScore) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec(ExName, ExType, MainTitle, Tag, MainText, QType, Answer, Choice, Score)
		if err != nil {
			panic(err)
		}
	} else {
		stmt, err := db.Prepare("INSERT INTO Question(ExName, ExType, QName, QTags, QMain, AnswerArea, QAnswer, QScore) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec(ExName, ExType, MainTitle, Tag, MainText, QType, Answer, Score)
		if err != nil {
			panic(err)
		}
	}

	return false
}

func Q_GetAll() []Structs.Question {
	db := ConnectDB()
	stmt, err := db.Query("SELECT * FROM Question")
	if err != nil {
		panic(err)
	}

	Questions := []Structs.Question{}

	for stmt.Next() {
		var QScore int
		var ExName, ExType, QName, QTags, QMain, AnswerArea, QAnswer, QChoice string

		if err := stmt.Scan(&ExName, &ExType, &QName, &QTags, &QMain, &AnswerArea, &QAnswer, &QChoice, &QScore); err != nil {
			panic(err)
		}

		Question := Structs.Question{
			ExName:     ExName,
			ExType:     ExType,
			QName:      QName,
			QTags:      QTags,
			QMain:      QMain,
			AnswerArea: AnswerArea,
			QAnswer:    QAnswer,
			QChoice:    QChoice,
			QScore:     QScore,
		}

		Questions = append(Questions, Question)
	}

	if err = stmt.Err(); err != nil {
		panic(err)
	}

	return Questions
}

func Q_Find(ExerciseName string, ExerciseType string, QuestionName string) Structs.Question {
	db := ConnectDB()
	stmt := db.QueryRow("SELECT * FROM Question WHERE ExName = ? AND ExType = ? And QName = ?", ExerciseName, ExerciseType, QuestionName)

	var QScore int
	var ExName, ExType, QName, QTags, QMain, AnswerArea, QAnswer, QChoice string

	if err := stmt.Scan(&ExName, &ExType, &QName, &QTags, &QMain, &AnswerArea, &QAnswer, &QChoice, &QScore); err != nil {
		panic(err)
	}

	Question := Structs.Question{
		ExName:     ExName,
		ExType:     ExType,
		QName:      QName,
		QTags:      QTags,
		QMain:      QMain,
		AnswerArea: AnswerArea,
		QAnswer:    QAnswer,
		QChoice:    QChoice,
		QScore:     QScore,
	}

	if err := stmt.Err(); err != nil {
		panic(err)
	}
	return Question
}

func Q_Update(ExName string, ExType string, MainTitle string, Tag string, MainText string, QType string, Answer string, Choice string, Score string) bool {
	db := ConnectDB()

	stmt, err := db.Prepare("UPDATE Exercise SET QName = ?, QTags = ?, QMain = ?, AnswerArea = ?, QAnswer = ?, QScore = ?,  WHERE (ExName = ? AND ExType = ? And QName = ?)")

	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(MainTitle, Tag, MainText, QType, Answer, Choice, Score, ExName, ExType, MainTitle)

	if err != nil {
		panic(err)
	}

	return false
}
