package Service

import (
	"github.com/gin-gonic/gin"
	"triela/Session"
)

func IsAuthMe(c *gin.Context) {
	php_sess_id, err := c.Cookie("PHPSESSID")

	if err != nil {
		panic(err)
	}

	SESSION := Session.Read(php_sess_id)
	if val, ok := SESSION["IsAuth"]; ok {
		if val == true {
			response := gin.H{
				"STATUS":             "SUCCESS",
				"AUTH_STATUS_NUMBER": "1",
				"MESSAGE":            "You are Authed",
			}
			c.JSON(200, response)
		} else {
			response := gin.H{
				"STATUS":             "SUCCESS",
				"AUTH_STATUS_NUMBER": "2",
				"MESSAGE":            "You are Logged In,But you not passed Two-Factor authorize.",
			}
			c.JSON(200, response)
		}
	} else {
		response := gin.H{
			"STATUS":             "SUCCESS",
			"AUTH_STATUS_NUMBER": "3",
			"MESSAGE":            "Not Authed",
		}
		c.JSON(200, response)
	}
}

func MyUserId(c *gin.Context) {
	php_sess_id, _ := c.Cookie("PHPSESSID")

	SESSION := Session.Read(php_sess_id)

	if val, ok := SESSION["UserId"]; ok {
		if val == true {
			response := gin.H{
				"STATUS":             "SUCCESS",
				"AUTH_STATUS_NUMBER": "1",
				"UID":                SESSION["UserId"],
			}
			c.JSON(200, response)
		} else {
			response := gin.H{
				"STATUS":             "SUCCESS",
				"AUTH_STATUS_NUMBER": "2",
				"MESSAGE":            "You are Logged In,But you not passed Two-Factor authorize.",
			}
			c.JSON(200, response)
		}
	} else {
		response := gin.H{
			"STATUS":             "SUCCESS",
			"AUTH_STATUS_NUMBER": "3",
			"MESSAGE":            "Not Authed",
		}
		c.JSON(200, response)
	}
}
