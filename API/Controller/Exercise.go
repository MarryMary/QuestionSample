package Controller

import (
	"github.com/gin-gonic/gin"
	"triela/Service"
)

func Push(c *gin.Context) {
	Service.PushExercise(c)
}
