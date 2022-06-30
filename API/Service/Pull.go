package Service

import (
	"github.com/gin-gonic/gin"
	"triela/Model"
)

func PullExercise(c *gin.Context) {
	res := Model.Ex_GetAll()
	json := gin.H{
		"STATUS": "SUCCESS",
		"DATA":   res,
	}

	c.JSON(200, json)

}
