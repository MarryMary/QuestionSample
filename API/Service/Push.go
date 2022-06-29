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
