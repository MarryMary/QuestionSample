package Routing

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func CatchRouting() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost",
			"http://localhost:8080",
			"http://localhost:8081",
		},

		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},

		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},

		AllowCredentials: true,
		
		MaxAge: 24 * time.Hour,
	}))

	r = GET_Routing(r)
	r = POST_Routing(r)

	return r
}
