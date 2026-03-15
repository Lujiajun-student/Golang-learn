package routers

import (
	"Golang-learn/Demo4_controllers/controllers/api"

	"github.com/gin-gonic/gin"
)

// InitApiRouters 初始化API路由
func InitApiRouters(r *gin.Engine) {
	var apiController api.ApiController

	// 路由分组
	apiRouters := r.Group("/api")
	{
		// 临时创建结构体实例进行使用
		apiRouters.GET("/", api.ApiController{}.ShowHomePage)
		// 先构建实例，再调用方法
		apiRouters.GET("/news", apiController.ShowNewsPage)
		// 继承DefaultController的方法
		apiRouters.GET("/default", apiController.ShowDefaultHomePage)
	}
}
