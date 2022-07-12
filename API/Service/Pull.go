package Service

import (
	"github.com/gin-gonic/gin"
	"triela/Cookie"
	"triela/Model"
	"triela/Session"
	"triela/Structs"
)

func PullExercise(c *gin.Context) {
	if Session.IsIn(Cookie.Read(c), "IsAuth") && Session.Get(Cookie.Read(c), "IsAuth") == true {
		res := Model.Ex_GetAll(Session.Get(Cookie.Read(c), "UserId"), "0")
		json := gin.H{
			"STATUS": "SUCCESS",
			"DATA":   res,
		}

		c.JSON(200, json)
	} else {
		json := gin.H{
			"STATUS":  "FAILED",
			"MESSAGE": "You are not authorized yet.",
		}

		c.JSON(403, json)
	}
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

	res := Model.Ex_Find(FindRequest.ExName, FindRequest.ExType)
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

	res := Model.Q_Find(FindQuestion.ExName, FindQuestion.ExType, FindQuestion.QName)
	json := gin.H{
		"STATUS": "SUCCESS",
		"DATA":   res,
	}

	c.JSON(200, json)
}
