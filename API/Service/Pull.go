package Service

import (
	"github.com/gin-gonic/gin"
	"triela/Model"
	"triela/Structs"
)

func PullExercise(c *gin.Context) {
	res := Model.Ex_GetAll()
	json := gin.H{
		"STATUS": "SUCCESS",
		"DATA":   res,
	}

	c.JSON(200, json)
}

func FindExercise(c *gin.Context) {
	var FindRequest Structs.ExerciseFindRequest

	if err := c.ShouldBindJSON(&FindRequest); err != nil {
		response := gin.H{
			"STATUS":     "FAILED",
			"ERRCODE":    "0001",
			"MESSAGE":    "UNEXPECTED SERVER ERR",
			"DEBUG_DATA": FindRequest,
		}

		c.JSON(500, response)
		return
	}

	res := Model.Ex_Find(&FindRequest.ExId)
	json := gin.H{
		"STATUS": "SUCCESS",
		"DATA":   res,
	}

	c.JSON(200, json)
}

func PullQuestion(c *gin.Context) {
	res := Model.Q_GetAll()
	json := gin.H{
		"STATUS": "SUCCESS",
		"DATA":   res,
	}

	c.JSON(200, json)
}

func FindQuestion(c *gin.Context) {
	var FindQuestion Structs.QuestionFindRequest

	if err := c.ShouldBindJSON(&FindQuestion); err != nil {
		response := gin.H{
			"STATUS":     "FAILED",
			"ERRCODE":    "0001",
			"MESSAGE":    "UNEXPECTED SERVER ERR",
			"DEBUG_DATA": FindQuestion,
		}

		c.JSON(500, response)
		return
	}

	res := Model.Q_Find(&FindQuestion.QId)
	json := gin.H{
		"STATUS": "SUCCESS",
		"DATA":   res,
	}

	c.JSON(200, json)
}
