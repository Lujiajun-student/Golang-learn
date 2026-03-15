package api

import (
	"Golang-learn/Demo4_controllers/controllers/_default"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
	_default.DefaultController
}

// ShowHomePage 加载首页，对应apiRouter.GET的"/"路径
func (ApiController *ApiController) ShowHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "api 首页",
	})
}

// ShowNewsPage 对应apiRouter.GET的"/news"路径
func (ApiController *ApiController) ShowNewsPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "api 新闻",
	})
}
