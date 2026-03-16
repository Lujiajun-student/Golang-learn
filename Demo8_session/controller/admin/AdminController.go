package admin

import "github.com/gin-gonic/gin"

func AdminIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "admin index",
	})
}

func AdminNews(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "admin news",
	})
}
