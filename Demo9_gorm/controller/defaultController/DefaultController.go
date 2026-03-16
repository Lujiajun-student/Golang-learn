package defaultController

import "github.com/gin-gonic/gin"

func DefaultNews(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "news",
	})
}

func DefaultArticle(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "article",
	})
}
