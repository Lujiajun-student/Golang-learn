package defaultRouter

import (
	"Golang-learn/Demo9_gorm/controller/defaultController"

	"github.com/gin-gonic/gin"
)

func InitDefaultRouter(r *gin.Engine) {
	defaultGroup := r.Group("/")
	{
		defaultGroup.GET("/news", defaultController.DefaultNews)
		defaultGroup.GET("/article", defaultController.DefaultArticle)
	}
}
