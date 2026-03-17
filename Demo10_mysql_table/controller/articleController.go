package controller

import (
	"Golang-learn/Demo10_mysql_table/models"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	BaseController
}

func (ArticleController) List(c *gin.Context) {

	var articleList []models.Article
	// Preload 用于关联查询，在执行Find查询前，先预查询ArticleCate的记录，然后根据外键ArticleCateId来关联查询ArticleCate的记录
	models.DB.Preload("ArticleCate").Find(&articleList)
	c.JSON(200, gin.H{
		"articleList": articleList,
	})
}

func (ArticleController) ArticleCateList(c *gin.Context) {
	var articleCateList []models.ArticleCate
	// 先通过Preload预查询Article记录，然后再根据外键ArticleCateId来关联查询ArticleCate的记录
	models.DB.Preload("Article").Find(&articleCateList)
	c.JSON(200, gin.H{
		"articleCateList": articleCateList,
	})
}
