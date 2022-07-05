package Routing

import (
	"github.com/gin-gonic/gin"
	"triela/Service"
)

func POST_Routing(r *gin.Engine) *gin.Engine {
	r.POST("/PushExercise", Service.PushExercise)
	r.POST("/FindExercise", Service.FindExercise)
	r.POST("/FixExercise", Service.UpdateExercise)
	r.POST("/PushQuestion", Service.PushQuestion)
	r.POST("/FindQuestion", Service.FindQuestion)

	return r
}
