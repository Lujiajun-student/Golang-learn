package router

import (
	"Golang-learn/Demo10_mysql_table/controller"

	"github.com/gin-gonic/gin"
)

func InitArticleRouter(r *gin.Engine) {

	var articleController controller.ArticleController

	articleGroup := r.Group("/article")
	{
		articleGroup.GET("/list", articleController.List)
		articleGroup.GET("/articlecatelist", articleController.ArticleCateList)
	}
}
