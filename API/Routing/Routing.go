package Routing

import "github.com/gin-gonic/gin"

func CatchRouting() *gin.Engine {
	r := gin.Default()
	r = GET_Routing(r)
	r = POST_Routing(r)

	return r
}