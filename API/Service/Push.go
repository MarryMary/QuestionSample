package Service

import (
	"github.com/gin-gonic/gin"
	"triela/Model"
)

type ExercisePushRequest struct {
	MainTitle string `json:"maintitle"`
	Detail    string `json:"detail"`
	Tag       string `json:"tag"`
	Limit     string `json:"limit"`
}

type ExerciseUpdateRequest struct {
	Id        string `json:"id"`
	MainTitle string `json:"maintitle"`
	Detail    string `json:"detail"`
	Tag       string `json:"tag"`
	Limit     string `json:"limit"`
}

type QuestionPushRequest struct {
	MainTitle string `json:"maintitle"`
	Tag       string `json:"tag"`
	MainText  string `json:"maintext"`
	QType     string `json:"qtype"`
	Answer    string `json:"answer"`
	Choice    string `json:"choice"`
	Score     string `json:"score"`
	IncExId   string `json:"incexid"`
}

type QuestionUpdateRequest struct {
	Id        string `json:"id"`
	MainTitle string `json:"maintitle"`
	Tag       string `json:"tag"`
	MainText  string `json:"maintext"`
	QType     string `json:"qtype"`
	Answer    string `json:"answer"`
	Choice    string `json:"choice"`
	Score     string `json:"score"`
}

func PushExercise(c *gin.Context) {
	var PushRequest ExercisePushRequest
	if err := c.ShouldBindJSON(&PushRequest); err != nil {
		response := gin.H{
			"STATUS":     "FAILED",
			"ERRCODE":    "0001",
			"MESSAGE":    "UNEXPECTED SERVER ERR",
			"DEBUG_DATA": PushRequest,
		}

		c.JSON(500, response)
		return
	}

	Model.Ex_Push(PushRequest.MainTitle, PushRequest.Detail, PushRequest.Tag, PushRequest.Limit)

	response := gin.H{
		"STATUS":  "SUCCESS",
		"MESSAGE": "EXERCISE JSON DATA INSERT COMPLETED",
	}

	c.JSON(200, response)
}

func UpdateExercise(c *gin.Context) {
	var UpdateRequest ExerciseUpdateRequest

	if err := c.ShouldBindJSON(&UpdateRequest); err != nil {
		response := gin.H{
			"STATUS":     "FAILED",
			"ERRCODE":    "0001",
			"MESSAGE":    "UNEXPECTED SERVER ERR",
			"DEBUG_DATA": UpdateRequest,
		}

		c.JSON(500, response)
		return
	}

	Model.Ex_Update(UpdateRequest.Id, UpdateRequest.MainTitle, UpdateRequest.Detail, UpdateRequest.Tag, UpdateRequest.Limit)
}

func PushQuestion(c *gin.Context) {
	var PushRequest QuestionPushRequest
	if err := c.ShouldBindJSON(&PushRequest); err != nil {
		response := gin.H{
			"STATUS":     "FAILED",
			"ERRCODE":    "0001",
			"MESSAGE":    "UNEXPECTED SERVER ERR",
			"DEBUG_DATA": PushRequest,
		}

		c.JSON(500, response)
		return
	}
	Model.Q_Push(PushRequest.MainTitle, PushRequest.Tag, PushRequest.MainText, PushRequest.QType, PushRequest.Answer, PushRequest.Choice, PushRequest.Score, PushRequest.IncExId)
}

func UpdateQuestion(c *gin.Context) {
	var UpdateQuestion QuestionUpdateRequest

	if err := c.ShouldBindJSON(&UpdateQuestion); err != nil {
		response := gin.H{
			"STATUS":     "FAILED",
			"ERRCODE":    "0001",
			"MESSAGE":    "UNEXPECTED SERVER ERR",
			"DEBUG_DATA": UpdateQuestion,
		}

		c.JSON(500, response)
		return
	}

	Model.Q_Update(UpdateQuestion.Id, UpdateQuestion.MainTitle, UpdateQuestion.Tag, UpdateQuestion.MainText, UpdateQuestion.QType, UpdateQuestion.Answer, UpdateQuestion.Choice, UpdateQuestion.Score)
}
