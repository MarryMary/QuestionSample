package Service

import "github.com/gin-gonic/gin"

func Teapot(c *gin.Context) {
	json := gin.H{
		"STATUS":  "SUCCESS",
		"MESSAGE": "この急須は生意気だ。",
	}

	c.JSON(418, json)
}

func TeapotAngry(c *gin.Context) {
	json := gin.H{
		"STATUS":  "FAILED",
		"ERRCODE": "0002",
		"MESSAGE": "ティーポット様がお茶以外のものを入れるなとお怒りでいらっしゃる、紅茶を入れて差し上げろ",
	}

	c.JSON(418, json)
}
