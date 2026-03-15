package routers

import (
	"Golang-learn/Demo4_controllers/controllers/_default"

	"github.com/gin-gonic/gin"
)

// InitDefaultRouters 初始化默认路由
func InitDefaultRouters(r *gin.Engine) {
	var defaultController _default.DefaultController

	defaultRouters := r.Group("/")
	{
		// 获取分组后的路由来拼接路径
		defaultRouters.GET("/", defaultController.ShowDefaultNewsPage)
		defaultRouters.GET("/news", defaultController.ShowDefaultNewsPage)
	}
}
