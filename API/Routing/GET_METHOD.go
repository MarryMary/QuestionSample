package Routing

import (
	"github.com/gin-gonic/gin"
	"triela/Service"
)

func GET_Routing(r *gin.Engine) *gin.Engine {
	// /teapotでSTATUS:SUCCESS、/teapotangryでSTATUS:FAILEDのJSONを疑似的に発生させることが可能。ステータスコードはジョーク用に予約された418が返却されます。
	r.GET("/teapot", Service.Teapot)
	r.GET("/teapotangry", Service.TeapotAngry)
	r.GET("/GetAllExercise", Service.PullExercise)
	r.GET("/GetAllQuestion", Service.PullQuestion)
	r.GET("/IsAuthMe", Service.IsAuthMe)

	return r
}
