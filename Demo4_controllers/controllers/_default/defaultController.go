package _default

import "github.com/gin-gonic/gin"

type DefaultController struct{}

// ShowDefaultHomePage 加载首页，对应defaultRouter.GET的"/"路径
func (defaultController *DefaultController) ShowDefaultHomePage(c *gin.Context) {
	c.String(200, "首页")
}

// ShowDefaultNewsPage 加载新闻，对应defaultRouter.GET的"/news"路径
func (defaultController *DefaultController) ShowDefaultNewsPage(c *gin.Context) {
	c.String(200, "新闻")
}
