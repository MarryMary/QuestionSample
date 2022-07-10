package Cookie

import "github.com/gin-gonic/gin"

func Read(c *gin.Context) string {
	result, err := c.Cookie("PHPSESSID")
	if err != nil {
		panic(err)
	}

	return result
}
