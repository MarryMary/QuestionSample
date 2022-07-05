package Model

import (
	"triela/Structs"
)

func Q_Push(MainTitle string, Tag string, MainText string, QType string, Answer string, Choice string, Score string, IncExId string) bool {
	db := ConnectDB()
	if QType == "select" {
		stmt, err := db.Prepare("INSERT INTO Question(QName, QTags, QMain, AnswerArea, QAnswer, QChoice, QScore, IncExId) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec(MainTitle, Tag, MainText, QType, Answer, Choice, Score, IncExId)
		if err != nil {
			panic(err)
		}
	} else {
		stmt, err := db.Prepare("INSERT INTO Question(QName, QTags, QMain, AnswerArea, QAnswer, QScore, IncExId) VALUES(?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec(MainTitle, Tag, MainText, QType, Answer, Score, IncExId)
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
		var QId, QScore, IncExId int
		var QName, QTags, QMain, AnswerArea, QAnswer, QChoice string

		if err := stmt.Scan(&QId, &QName, &QTags, &QMain, &AnswerArea, &QAnswer, &QChoice, &QScore, &IncExId); err != nil {
			panic(err)
		}

		Question := Structs.Question{
			QId:        QId,
			QName:      QName,
			QTags:      QTags,
			QMain:      QMain,
			AnswerArea: AnswerArea,
			QAnswer:    QAnswer,
			QChoice:    QChoice,
			QScore:     QScore,
			IncExId:    IncExId,
		}

		Questions = append(Questions, Question)
	}

	if err = stmt.Err(); err != nil {
		panic(err)
	}

	return Questions
}

func Q_Find(QueId *string) Structs.Question {
	db := ConnectDB()
	stmt := db.QueryRow("SELECT * FROM Question WHERE QId = ?", QueId)

	var QId, QScore, IncExId int
	var QName, QTags, QMain, AnswerArea, QAnswer, QChoice string

	if err := stmt.Scan(&QId, &QName, &QTags, &QMain, &AnswerArea, &QAnswer, &QChoice, &QScore, &IncExId); err != nil {
		panic(err)
	}

	Question := Structs.Question{
		QId:        QId,
		QName:      QName,
		QTags:      QTags,
		QMain:      QMain,
		AnswerArea: AnswerArea,
		QAnswer:    QAnswer,
		QChoice:    QChoice,
		QScore:     QScore,
		IncExId:    IncExId,
	}

	if err := stmt.Err(); err != nil {
		panic(err)
	}
	return Question
}

func Q_Update(QId string, MainTitle string, Tag string, MainText string, QType string, Answer string, Choice string, Score string) bool {
	db := ConnectDB()

	stmt, err := db.Prepare("UPDATE Exercise SET QName = ?, QTags = ?, QMain = ?, AnswerArea = ?, QAnswer = ?, QScore = ?,  WHERE (QId = ?)")

	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(MainTitle, Tag, MainText, QType, Answer, Choice, Score, QId)

	if err != nil {
		panic(err)
	}

	return false
}
