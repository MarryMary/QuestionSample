package Service

import (
	"github.com/gin-gonic/gin"
	"triela/Model"
	"triela/Structs"
)

func PushExercise(c *gin.Context) {
	var PushRequest Structs.ExercisePushRequest
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

	Model.Ex_Push(PushRequest.MainTitle, PushRequest.Year, PushRequest.Season, PushRequest.Genre, PushRequest.Tag, PushRequest.Limit)

	response := gin.H{
		"STATUS":  "SUCCESS",
		"MESSAGE": "EXERCISE JSON DATA INSERT COMPLETED",
	}

	c.JSON(200, response)
}

func UpdateExercise(c *gin.Context) {
	var UpdateRequest Structs.ExerciseUpdateRequest

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

	Model.Ex_Update(UpdateRequest.OldTitle, UpdateRequest.OldType, UpdateRequest.MainTitle, UpdateRequest.Year, UpdateRequest.Season, UpdateRequest.Type, UpdateRequest.Tag, UpdateRequest.Limit)
}

func PushQuestion(c *gin.Context) {
	var PushRequest Structs.QuestionPushRequest
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
	var UpdateQuestion Structs.QuestionUpdateRequest

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
