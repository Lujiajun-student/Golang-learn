package routers

import "github.com/gin-gonic/gin"

// InitApiRouters 初始化API路由
func InitApiRouters(r *gin.Engine) {
	// 路由分组
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "api首页",
			})
		})
		apiRouters.GET("/news", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "api新闻",
			})
		})
	}
}
