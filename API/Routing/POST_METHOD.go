package Routing

import (
	"github.com/gin-gonic/gin"
	"triela/Service"
)

func POST_Routing(r *gin.Engine) *gin.Engine {
	r.POST("/PushExercise", Service.PushExercise)

	return r
}
